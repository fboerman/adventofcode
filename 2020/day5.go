package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("day5_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	high := 0
	var seat_ids []int

	for scanner.Scan() {
		line := scanner.Text()

		// initiate the bounds of rows and columns (between a and b)
		var r_a, r_b, c_a, c_b int
		r_b = 127
		c_b = 7

		for _, v := range line {
			c := string(v)
			switch c {
			case "B":
				r_a = r_a + ((r_b+1)-r_a)>>1
				break
			case "F":
				r_b = r_b - ((r_b+1)-r_a)>>1
				break
			case "R":
				c_a = c_a + ((c_b+1)-c_a)>>1
				break
			case "L":
				c_b = c_b - ((c_b+1)-c_a)>>1
				break
			default:
				fmt.Println("[!] Invalid character: ", c)
				os.Exit(-1)
			}
		}
		if r_a != r_b || c_a != c_b {
			fmt.Println("[!] non conversion for: ", line)
			os.Exit(-1)
		}
		seat_id := r_a*8 + c_a
		if seat_id > high {
			high = seat_id
		}
		//fmt.Printf("row: %d, column: %d, seat id: %d\n", r_a, c_a, seat_id)
		seat_ids = append(seat_ids, seat_id)
	}

	fmt.Println("Highest seat id: ", high)
	table := make([]int, high+1)

	for _, id := range seat_ids {
		table[id] = 1
	}

	for i, flag := range table {
		if i != 0 && i != high {
			if table[i-1] == 1 && flag == 0 && table[i+1] == 1 {
				fmt.Println("Your seat id: ", i)
			}
		}

	}
}
