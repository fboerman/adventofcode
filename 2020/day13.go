package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("day13_input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	lines := strings.Split(string(content), "\n")
	if len(lines) != 2 {
		fmt.Println("Invalid input file!")
		os.Exit(-1)
	}

	T, _ := strconv.Atoi(lines[0])
	busses := strings.Split(lines[1], ",")
	nextbus := math.MaxInt32
	nextbus_id := -1
	for _, bus := range busses {
		if bus == "x" {
			continue
		}
		bus_id, _ := strconv.Atoi(bus)
		departure := (T/bus_id + 1) * bus_id
		if departure < nextbus {
			nextbus = departure
			nextbus_id = bus_id
		}
	}

	fmt.Println("First next bus leaves at:", nextbus)
	fmt.Println("Puzzle hash:", (nextbus-T)*nextbus_id)
}
