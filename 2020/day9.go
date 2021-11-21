package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func get_next_number(scanner *bufio.Scanner) (int, bool) {
	if !scanner.Scan() {
		return 0, false
	}

	line := scanner.Text()
	number, _ := strconv.Atoi(line)

	return number, true
}

func is_valid(numbers *[]int, target int, start int) bool {
	for i := start; i < len(*numbers); i++ {
		for j := i; j < len(*numbers); j++ {
			if (*numbers)[i]+(*numbers)[j] == target {
				return true
			}
		}
	}

	return false
}

func main() {
	preamble := 25

	file, err := os.Open("day9_input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers []int

	for { //i := 0; i < preamble; i++ {
		number, flag := get_next_number(scanner)
		if !flag {
			break
		}
		numbers = append(numbers, number)
	}
	target := 0
	for i := preamble; i < len(numbers); i++ {
		if !is_valid(&numbers, numbers[i], i-preamble) {
			fmt.Println("First invalid number:", numbers[i])
			target = numbers[i]
			break
		}
	}
	var window []int
	for i := 0; i < len(numbers); i++ {
		sum := 0
		window = []int{}
		for z := i; z < len(numbers); z++ {
			sum += numbers[z]
			window = append(window, numbers[z])
			if sum >= target {
				break
			}
		}
		if sum == target {
			break
		}
	}
	sort.Ints(window)
	fmt.Println("encryption weakness:", window[0]+window[len(window)-1])
}
