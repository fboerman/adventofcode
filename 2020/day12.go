package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	file, err := os.Open("day12_input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// curent position
	// grid is topleft oriented so down is positive ship_y and right is positive ship_x
	ship_x := 0.0
	ship_y := 0.0
	ship_dir := 0 // degrees, 0 is eastwards, clockwise is positive

	waypoint_x := 10.0
	waypoint_y := -1.0

	ship2_x := 0.0
	ship2_y := 0.0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var instr byte
		var arg_i int
		_, _ = fmt.Sscanf(scanner.Text(), "%c%03d", &instr, &arg_i)
		arg := float64(arg_i)
		switch instr {
		case 'N':
			ship_y -= arg
			waypoint_y -= arg
			break
		case 'S':
			ship_y += arg
			waypoint_y += arg
			break
		case 'E':
			ship_x += arg
			waypoint_x += arg
			break
		case 'W':
			ship_x -= arg
			waypoint_x -= arg
			break
		case 'L':
			ship_dir -= arg_i
			temp_x := waypoint_x
			temp_y := waypoint_y
			arg *= -1
			arg *= math.Pi / 180
			waypoint_x = temp_x*math.Cos(arg) - temp_y*math.Sin(arg)
			waypoint_y = temp_x*math.Sin(arg) + temp_y*math.Cos(arg)
			break
		case 'R':
			ship_dir += arg_i
			temp_x := waypoint_x
			temp_y := waypoint_y
			arg *= math.Pi / 180
			waypoint_x = temp_x*math.Cos(arg) - temp_y*math.Sin(arg)
			waypoint_y = temp_x*math.Sin(arg) + temp_y*math.Cos(arg)
			break
		case 'F':
			ship_x += arg * math.Cos(float64(ship_dir)*math.Pi/180)
			ship_y += arg * math.Sin(float64(ship_dir)*math.Pi/180)

			ship2_x += arg * waypoint_x
			ship2_y += arg * waypoint_y

			break
		}
		fmt.Println(scanner.Text(), ship2_x, ship2_y, waypoint_x, waypoint_y)
	}

	fmt.Println("Manhattan distance part1:", math.Abs(ship_x)+math.Abs(ship_y))
	fmt.Println("Manhattan distance part2:", math.Abs(ship2_x)+math.Abs(ship2_y))
}
