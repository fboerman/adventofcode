package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sync"
)

type Astroid struct {
	X int
	Y int
}

func parse_astroid_map(map_raw []byte) (result []Astroid) {
	y := 0
	x := 0

	for _, c := range map_raw {
		//if it is an astroid create struct
		//if it is emtpy space then we have shifted one to right
		//if newline then we have shifted one line down
		switch c {
		case '#':
			result = append(result, Astroid{x, y})
			x++
		case '\n':
			x = 0
			y++
		default:
			x++
		}

	}

	return
}

type AstroidInfo struct {
	src               Astroid
	num_line_of_sight int
}

func distance(source Astroid, target Astroid) float64 {
	return math.Sqrt(math.Pow(float64(source.X-target.X), 2) + math.Pow(float64(source.Y-target.Y), 2))
}

// calculate the number of line-of-sight astroids from given astroid on astroid map
func calc_astroids(astroid_map []Astroid, source Astroid, c chan AstroidInfo) {
	// for each astroid that is not the source
	//   calculate its line-of-sight function from source to selected astroid
	//   check for all other astroids if they are on this line, if so decrease number line of sight by one
	result := AstroidInfo{source, len(astroid_map) - 1}

	for _, selected := range astroid_map {
		if source.X == selected.X && source.Y == selected.Y {
			continue
		}
		// line of sight is straight line so form y=a*x+b
		a := float32(selected.Y-source.Y) / float32(selected.X-source.X)
		b := float32(selected.Y) - a*float32(selected.X)

		//iterate again and check if there is an astroid in the way, if so then decrease number reachable
		selected_distance := distance(source, selected)
		for _, selected2 := range astroid_map {
			if (source.X == selected2.X && source.Y == selected2.Y) ||
				(selected.X == selected2.X && selected.Y == selected2.Y) {
				continue
			}
			// the second clause in the OR is for when the line is straight up since a=inf then

			if (float32(selected2.Y) == float32(selected2.X)*a+b) || (selected.X-source.X == 0 && selected.X == selected2.X) {
				//the selected2 is on the line of sight, now check if it is between the two points
				//smaller then in the if statement to get rid of the floating point error
				distance_total := distance(source, selected2) + distance(selected2, selected)

				if math.Abs(distance_total-selected_distance) < 1e-9 {
					result.num_line_of_sight--
					break
				}
			}
		}
	}

	c <- result
}

var wg sync.WaitGroup

func main() {
	map_raw, _ := ioutil.ReadFile("day10_input.txt")

	astroids := parse_astroid_map(map_raw)

	fmt.Println("number of astroids on map: ", len(astroids))
	//fmt.Println(calc_astroids(astroids, astroids[6]).num_line_of_sight)
	c := make(chan AstroidInfo)

	for _, astr := range astroids {
		//fmt.Printf("astr X: %d Y:%d has line of sight to %d\n", astr.X, astr.Y, calc_astroids(astroids, astr).num_line_of_sight)
		go calc_astroids(astroids, astr, c)
	}

	wg.Wait()
	winner := <-c
	for i := 1; i < len(astroids); i++ {
		info := <-c
		if info.num_line_of_sight > winner.num_line_of_sight {
			winner = info
		}
		fmt.Println(info)
	}

	fmt.Println("part 1:")
	fmt.Println(winner)
}
