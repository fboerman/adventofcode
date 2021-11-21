package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// open the file
	file, _ := os.Open("day2_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	valid_part1 := 0
	valid_part2 := 0
	for scanner.Scan() {
		var a, b int
		var letter byte
		var pwd string

		line := scanner.Text()
		fmt.Sscanf(line, "%d-%d %c: %s", &a, &b, &letter, &pwd)
		count := 0
		for _, char := range pwd {
			if byte(char) == letter {
				count += 1
			}
		}

		if count >= a && count <= b {
			valid_part1 += 1
		}

		if (pwd[a-1] == letter && pwd[b-1] != letter) || (pwd[b-1] == letter && pwd[a-1] != letter) {
			valid_part2 += 1
		}
	}

	fmt.Println("Part 1: ", valid_part1)
	fmt.Println("Part 2: ", valid_part2)

}
