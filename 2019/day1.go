package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calculateFuel(mass int) int{
	if mass <= 0 {
		return 0
	}
	fuel := int(math.Floor(float64(mass)/3)) - 2
	if fuel <= 0 {
		return 0
	}
	return calculateFuel(fuel) + fuel
}

func main() {
	file, _ := os.Open("day1_input.txt")
	defer file.Close()
	var lines []int

	scanner := bufio.NewScanner(file)


	// first read all lines, so that part 2 does not need new fileaccess
	for scanner.Scan() {
		inp, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, int(inp))
	}

	// calculate part 1
	var sum int
	for i:= 0; i < len(lines); i++ {
		sum += int(math.Floor(float64(lines[i])/3)) - 2
	}
	fmt.Printf("part 1: %d\n", sum)

	// calculate part 2
	sum = 0
	for i:= 0; i < len(lines); i++ {
		sum += calculateFuel(lines[i])
	}

	fmt.Printf("part 2: %d\n", sum)
}
