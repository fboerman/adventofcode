package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(numbers []int) int{
	// iterate through all numbers
	for i:=0; i<len(numbers); i++ {
		// check only for the numbers AFTER this one, because we already visited all before this one
		for j:=i; j<len(numbers); j++ {
			if numbers[i] + numbers[j] == 2020 {
				return numbers[i] * numbers[j]
			}
		}
	}
	return 0
}

func part2(numbers []int) int{
	// iterate through all numbers
	for i:=0; i<len(numbers); i++ {
		// check only for the numbers AFTER this one, because we already visited all before this one
		for j:=i; j<len(numbers); j++ {
			for q:=j; q<len(numbers); q++ {
				if numbers[i] + numbers[j] + numbers[q] == 2020 {
					return numbers[i] * numbers[j] * numbers[q]
				}
			}
		}
	}
	return 0
}

func main() {
	// open the file
	file, _ := os.Open("day1_input.txt")
	defer file.Close()

	// scan the file and convert to integer
	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		inp, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, int(inp))
	}

	fmt.Printf("Result part 1: %d\n", part1(numbers))
	fmt.Printf("Result part 2: %d\n", part2(numbers))

}