//naive version

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Group struct {
	num        int
	questions  map[string]int
	num_anyone int
}

func parse_groupline(line string, g *Group) {
	g.num++
	for _, c := range line {
		letter := string(c)
		if g.questions[letter] == 0 {
			g.num_anyone++
		}
		g.questions[letter]++
	}
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
	var groups []Group
	g := Group{questions: map[string]int{}}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			groups = append(groups, g)
			g = Group{questions: map[string]int{}}
		} else {
			parse_groupline(line, &g)
		}
	}
	groups = append(groups, g)
	sum_part1 := 0
	sum_part2 := 0
	for _, g := range groups {
		sum_part1 += g.num_anyone
		for _, q := range g.questions {
			if q == g.num {
				sum_part2++
			}
		}
	}
	fmt.Println("Sum of someone answered yes:", sum_part1)
	fmt.Println("Sum of everyone answered yes:", sum_part2)

}
