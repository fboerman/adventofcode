package main

import (
	"bufio"
	"fmt"
	"os"
)

type World struct {
	Data    map[[3]int]bool
	X, Y, Z [2]int
}

func print_world(W *World) {
	for z := W.Z[0]; z < W.Z[1]; z++ {
		fmt.Printf("z=%d\n", z)
		for y := W.Y[0]; y < W.Y[1]; y++ {
			for x := W.X[0]; x < W.X[1]; x++ {
				if W.Data[[3]int{x, y, z}] {
					fmt.Printf(" #")
				} else {
					fmt.Printf(" .")
				}

			}
			fmt.Printf("\n")
		}
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func alive_neighours(W *World, x, y, z int) int {
	num := 0

	// loop through all 27 cubes in the 3 by 3 subslice of the world
	for _, z_ := range [3]int{-1, -0, 1} {
		for _, y_ := range [3]int{-1, -0, 1} {
			for _, x_ := range [3]int{-1, -0, 1} {
				// select only neighbors which differ at most one in any of their coordinates
				if Abs(x_)+Abs(y_)+Abs(z_) == 1 {
					if W.Data[[3]int{x + x_, y + y_, z + z_}] {
						num++
					}
				}
			}
		}
	}
	fmt.Printf("x:%d\ty:%d\tz:%d\tnum=%d\n", x, y, z, num)
	return num
}

func next(W_read *World, W_write *World) {
	for z := W_read.Z[0] - 1; z < W_read.Z[1]+1; z++ {
		for y := W_read.Y[0] - 1; y < W_read.Y[1]+1; y++ {
			for x := W_read.X[0] - 1; x < W_read.X[1]+1; x++ {
				// for each point check all neighbors
				num := alive_neighours(W_read, x, y, z)
				state := W_read.Data[[3]int{x, y, z}]
				if state {
					W_write.Data[[3]int{x, y, z}] = num == 2 || num == 3

				} else {
					if num == 3 {
						W_write.Data[[3]int{x, y, z}] = true
						// update boundaries if a cube is energized
						if x < W_write.X[0] {
							W_write.X[0] = x
						}
						if x > W_write.X[1] {
							W_write.X[1] = x
						}
						if y < W_write.Y[0] {
							W_write.Y[0] = y
						}
						if y > W_write.Y[1] {
							W_write.Y[1] = y
						}
						if z < W_write.Z[0] {
							W_write.Z[0] = z
						}
						if z > W_write.Z[1] {
							W_write.Z[1] = z
						}
					}
				}
			}
		}
	}
}

func main() {
	file, err := os.Open("day17_input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	world1 := World{X: [2]int{0, -1}, Data: map[[3]int]bool{}, Z: [2]int{0, 1}}
	world2 := World{X: [2]int{0, -1}, Data: map[[3]int]bool{}, Z: [2]int{0, 1}}
	y := 0
	x := -1
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			world1.Data[[3]int{i, y, 0}] = c == '#'
			world2.Data[[3]int{i, y, 0}] = c == '#'
		}
		y++
		if x == -1 {
			x = len(line)
		}
	}

	world1.Y[1] = y
	world2.Y[1] = y
	world1.X[1] = x
	world2.X[1] = x

	read_world := &world1
	write_world := &world2

	print_world(read_world)

	next(read_world, write_world)
	read_world.X = write_world.X
	read_world.Y = write_world.Y
	read_world.Z = write_world.Z
	fmt.Println("iteration 1")
	print_world(write_world)

}
