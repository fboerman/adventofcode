package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type BagHolder struct {
	num int
	bag *Bag
}

type Bag struct {
	name    string
	holders []BagHolder
}

func find_color(current *Bag, origin *Bag, searchterm string) bool {
	if current.name == searchterm {
		return true
	}

	if origin.name == current.name {
		return false
	}

	for _, hold := range current.holders {
		if find_color(hold.bag, origin, searchterm) {
			return true
		}
	}

	return false
}

func count_bags(current *Bag, origin *Bag) int {
	counter := 0
	for _, hold := range current.holders {
		if hold.bag.name == origin.name {
			continue
		}
		c := count_bags(hold.bag, origin)
		counter += hold.num * c
	}
	counter += 1 //myself
	return counter
}

func main() {
	file, err := os.Open("day7_input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	bags := map[string]*Bag{}

	for scanner.Scan() {
		// make sure the words are uniform
		line := strings.Trim(
			strings.Replace(scanner.Text(), "bags", "bag", -1),
			".")
		parts := strings.Split(line, " contain ")
		var names [2]string
		_, _ = fmt.Sscanf(parts[0], "%s %s bags", &names[0], &names[1])
		name := fmt.Sprintf("%s %s", names[0], names[1])
		b := bags[name]
		if b == nil {
			b = &Bag{name: name}
			bags[name] = b
		}
		if strings.Contains(parts[1], "no other") {
			continue
		}
		contents := strings.Split(parts[1], ", ")
		for _, content := range contents {
			hold := BagHolder{}
			_, _ = fmt.Sscanf(content, "%d %s %s bag", &hold.num, &names[0], &names[1])
			name_c := fmt.Sprintf("%s %s", names[0], names[1])
			b_c := bags[name_c]
			if b_c == nil {
				b_c = &Bag{name: name_c}
				bags[name_c] = b_c
			}
			hold.bag = b_c
			b.holders = append(b.holders, hold)
		}

	}
	sum := 0
	for _, bag := range bags {
		if bag.name == "shiny gold" {
			continue
		}
		for _, hold := range bag.holders {
			if find_color(hold.bag, bag, "shiny gold") {
				sum += 1
				break
			}
		}

	}
	counter := count_bags(bags["shiny gold"], bags["shiny gold"]) - 1 // -1 because shiny bag itself doesnt count
	fmt.Println("part 1:", sum)
	fmt.Println("part 2:", counter)
}
