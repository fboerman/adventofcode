package main

import (
	"bufio"
	"fmt"
	"os"
)

func check_slope(chart [][]byte, slope [2]int) int {
	trees := 0
	x := slope[0]
	for y := slope[1]; y < len(chart); y += slope[1] {
		if chart[y][x] == '#' {
			trees++
		}
		x = (x + slope[0]) % len(chart[0])
	}

	return trees
}

func main() {
	file, _ := os.Open("day3_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var chart [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		var row []byte
		for _, cell := range line {
			row = append(row, byte(cell))
		}
		chart = append(chart, row)
	}

	var slopes = [...][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	result := 0
	for _, slope := range slopes {
		trees := check_slope(chart, slope)
		fmt.Println(trees)
		if result == 0 {
			result = trees
		} else {
			result *= trees
		}
	}
	fmt.Println(result)
}
