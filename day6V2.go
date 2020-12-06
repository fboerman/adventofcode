// little more efficient version

package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func smash_group(group []uint) (int, int) {
	reg1 := group[0]
	reg2 := group[0]
	for _, g := range group[1:] {
		reg1 |= g
		reg2 &= g
	}

	return bits.OnesCount(reg1), bits.OnesCount(reg2)
}

func main() {
	file, err := os.Open("day6_input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum_part1 := 0
	sum_part2 := 0

	var group []uint
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			part1, part2 := smash_group(group)
			sum_part1 += part1
			sum_part2 += part2
			group = []uint{}
		} else {
			var register uint
			for _, c := range line {
				register |= uint(2 << (uint(c) - 97))
			}
			group = append(group, register)
		}
	}
	part1, part2 := smash_group(group)
	sum_part1 += part1
	sum_part2 += part2

	fmt.Println("Sum of someone answered yes:", sum_part1)
	fmt.Println("Sum of everyone answered yes:", sum_part2)

}
