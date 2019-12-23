package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("day1_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		inp, _ := strconv.Atoi(scanner.Text())
		num := int(inp)
		sum += int(math.Floor(float64(num)/3)) - 2
	}
	fmt.Print("part 1: ")
	fmt.Println(sum)
}
