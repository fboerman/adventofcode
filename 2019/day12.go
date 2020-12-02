package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"math"
	"os"
)

type Moon struct {
	i   int
	vel [3]int
	pos [3]int
}

// not working first attempt
//func moon_axis_is_notunique(m Moon, moons []Moon, q int) bool {
//	for i:=0;i<len(moons);i++ {
//		if m.pos[q] == moons[i].pos[q] && m.i != moons[i].i {
//			return true
//		}
//	}
//
//	return false
//}
//
//func exec_step(moons []Moon) {
//	// sort for one axis
//	// apply formula on that axis velocity += l-1-2i
//	for q:=0;q<3;q++ {
//		sort.Slice(moons, func(i, j int) bool {return moons[i].pos[q] < moons[j].pos[q]})
//		// update velocity, and skip if axis is not unique
//		for i:=0;i<len(moons);i++ {
//			if moon_axis_is_notunique(moons[i], moons, q) {
//				continue
//			}
//			moons[i].vel[q] += 3-2*i
//		}
//		// update position
//		for i:=0;i<len(moons);i++ {
//			moons[i].pos[q] += moons[i].vel[q]
//		}
//	}
//}
// end of first attempt

func signbit(x int) int {
	if x > 0 {
		return 1
	}
	return -1
}

//second attempt
func exec_step(moons []Moon) {
	for q := 0; q < 3; q++ {
		for i := 0; i < len(moons); i++ {
			for j := i + 1; j < len(moons); j++ {
				if moons[i].pos[q] == moons[j].pos[q] {
					continue
				}
				sign := signbit(moons[i].pos[q] - moons[j].pos[q])
				moons[i].vel[q] -= sign
				moons[j].vel[q] += sign
			}
		}

		for i := 0; i < len(moons); i++ {
			moons[i].pos[q] += moons[i].vel[q]
		}
	}

}

func hash_in_list(hash [16]byte, l [][16]byte) bool {
	for _, y := range l {
		if bytes.Equal(hash[:], y[:]) {
			return true
		}
	}
	return false
}

func hash_tuples(list [][2]int) [16]byte {
	var buffer bytes.Buffer
	for _, v := range list {
		buffer.WriteString(fmt.Sprintf("%d|%d", v[0], v[1]))
		buffer.WriteString(",")
	}
	return md5.Sum(buffer.Bytes())
}

func get_state_of_axis(moons []Moon, q int) [][2]int {
	var state [][2]int
	for _, moon := range moons {
		state = append(state, [2]int{moon.pos[q], moon.vel[q]})
	}

	return state
}

func main() {
	// open the file
	file, _ := os.Open("day12_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var moons []Moon
	for scanner.Scan() {
		var x, y, z int
		line := scanner.Text()
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		moons = append(moons, Moon{len(moons), [3]int{}, [3]int{x, y, z}})
		//fmt.Println(moons[len(moons)-1])
	}

	if os.Args[1] == "1" {
		for step := 0; step < 1000; step++ {
			exec_step(moons)

			//fmt.Printf("step: %d\n", step+1)
			//for i := 0; i < len(moons); i++ {
			//fmt.Printf("%d: pos=<x= %d, y=%d, z= %d>, vel=<x= %d, y=%d, z=%d>\n", moons[i].i,
			//	moons[i].pos[0], moons[i].pos[1], moons[i].pos[2],
			//	moons[i].vel[0], moons[i].vel[1], moons[i].vel[2])
			//}
		}

		energy := 0.0
		for i := 0; i < len(moons); i++ {
			vel := 0.0
			pos := 0.0
			for q := 0; q < 3; q++ {
				vel += math.Abs(float64(moons[i].vel[q]))
				pos += math.Abs(float64(moons[i].pos[q]))
			}
			energy += vel * pos
		}

		fmt.Println("Part 1: total energy: ", energy)
	} else if os.Args[1] == "2" {
		// I will admit that I got a tip that lcm could be used, I did not get that at first
		// https://en.wikipedia.org/wiki/Least_common_multiple#Planetary_alignment
		var x_states [][16]byte
		var y_states [][16]byte
		var z_states [][16]byte
		x_period := -1
		y_period := -1
		z_period := -1
		step := 0

		for ; ; step++ {
			exec_step(moons)
			// calculate the hashes for state of each non solved axis
			// first build state then hash
			// check if this hash has occurred before, if so record number of steps, if not add to hash list
			if x_period == -1 {
				state_x := get_state_of_axis(moons, 0)
				hash := hash_tuples(state_x)
				if hash_in_list(hash, x_states) {
					x_period = step
					fmt.Println("period for x: ", x_period)
				} else {
					x_states = append(x_states, hash)
				}
			}

			if y_period == -1 {
				state_y := get_state_of_axis(moons, 1)
				hash := hash_tuples(state_y)
				if hash_in_list(hash, y_states) {
					y_period = step
					fmt.Println("period for y: ", y_period)
				} else {
					y_states = append(y_states, hash)
				}
			}

			if z_period == -1 {
				state_z := get_state_of_axis(moons, 2)
				hash := hash_tuples(state_z)

				if hash_in_list(hash, z_states) {
					z_period = step
					fmt.Println("period for z: ", z_period)
				} else {
					z_states = append(z_states, hash)
				}
			}

			if x_period != -1 && y_period != -1 && z_period != -1 {
				break
			}

		}

		fmt.Println("Part 2: calculate lcm from: ", x_period, y_period, z_period)

	}
}
