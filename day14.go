package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func permutate(selected uint64, bits_floating []int, all *[]uint64) {
	if len(bits_floating) == 0 {
		*all = append(*all, selected)
		return
	}
	bit := bits_floating[0]
	bits_floating = bits_floating[1:]
	permutate(selected|(1<<bit), bits_floating, all)
	permutate(selected & ^(1<<bit), bits_floating, all)
}

func main() {
	file, err := os.Open("day14_input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	part, _ := strconv.Atoi(os.Args[1])

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var mask_on, mask_off uint64
	var mask_floating []int
	mem := map[uint64]uint64{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "mask") {
			mask_on = 0
			mask_off = 0
			mask_floating = []int{}
			var mask string
			_, _ = fmt.Sscanf(line, "mask = %s", &mask)
			for i, c := range mask {
				switch c {
				case '1':
					mask_on |= 1 << (len(mask) - i - 1)
					break
				case '0':
					mask_off |= 1 << (len(mask) - i - 1)
					break
				case 'X':
					mask_floating = append(mask_floating, len(mask)-i-1)
				}
			}
		} else {
			var addr, value uint64
			_, _ = fmt.Sscanf(line, "mem[%d] = %d", &addr, &value)
			if part == 1 {
				value &= ^mask_off
				value |= mask_on
				mem[addr] = value
			} else if part == 2 {
				addr |= mask_on
				l := []uint64{}
				permutate(addr, mask_floating, &l)
				for _, a := range l {
					mem[a] = value
					//fmt.Printf("%d = %d\n", a, value)
				}
			}
			//fmt.Printf("%d = %d\n", addr, value)
		}
	}
	var sum uint64
	for _, v := range mem {
		sum += v
	}

	fmt.Println(sum)
}
