package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("day10_input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	input := []int{0}

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		input = append(input, i)
	}
	sort.Ints(input)
	counts := map[int]int{}

	for i := 1; i < len(input); i++ {
		counts[Abs(input[i]-input[i-1])]++
	}

	counts[3] += 1

	fmt.Println(counts)
	fmt.Println(counts[1] * counts[3])
}
