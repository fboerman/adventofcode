package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
	steps int
}

func isInArray(p Point, arr []Point) *Point {
	for i:=0;i<len(arr);i++ {
		if p.X == arr[i].X && p.Y == arr[i].Y {
			return &arr[i]
		}
	}

	return nil
}

func main() {
	file, _ := os.Open("day3_input.txt")
	defer file.Close()
	var lines []string

	scanner := bufio.NewScanner(file)

	// first read all lines, so that part 2 does not need new fileaccess
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	wire1 := strings.Split(lines[0], ",")
	wire2 := strings.Split(lines[1], ",")

	//parse the first wire, create a list of its points
	var W1 []Point
	x :=0
	y :=0
	steps := 0
	for i := 0; i < len(wire1); i++ {
		instr := wire1[i]
		numsteps, _ := strconv.Atoi(instr[1:])
		for i:=0; i<numsteps; i++ {
			steps++
			switch direction := instr[0:1]; direction {
			case "R":
				x++
			case "D":
				y--
			case "L":
				x--
			case "U":
				y++
			}
			W1 = append(W1, Point{X:x, Y:y, steps:steps})
		}
	}

	// parse the second wire, for each point calculated check if it is an intersection,
	// for part 1 if yes calculate the manhattan distance and check if this is the smallest one
	// for part 2 if yes calculate the sum of steps and check if this is the smallest one
	manhattan_distance := math.MaxFloat64
	lowest_steps := math.MaxInt32
	x = 0
	y = 0
	steps = 0
	for i := 0; i < len(wire2); i++ {
		instr := wire2[i]
		numsteps, _ := strconv.Atoi(instr[1:])
		for i:=0; i<numsteps; i++ {
			steps++
			switch direction := instr[0:1]; direction {
			case "R":
				x++
			case "D":
				y--
			case "L":
				x--
			case "U":
				y++
			}
			intersection := isInArray(Point{X:x, Y:y}, W1)
			if intersection != nil {
				// for part 1
				dist := math.Abs(float64(x)) + math.Abs(float64(y))
				if dist < manhattan_distance{
					manhattan_distance = dist
				}
				// for part 2
				if min_steps_now := steps + intersection.steps; min_steps_now < lowest_steps {
					lowest_steps = min_steps_now
				}
			}
		}
	}

	fmt.Println("part 1: ", manhattan_distance)
	fmt.Println("part 2: ", lowest_steps)

}
