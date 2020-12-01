package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func load_program(fname string) (input []int) {
	// taken from bufio/example_test.go for splitting at comma
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}

	// read and parse the file by comma splitting
	file, _ := os.Open(fname)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(onComma)

	for scanner.Scan() {
		c := scanner.Text()
		c = strings.TrimSuffix(c, "\n")
		inp, _ := strconv.Atoi(c)
		input = append(input, inp)
	}

	return
}

// convert the number and mode to a pointer to the value
// so in essence convert the argument such that it is
func get_val_p(num *int, mode byte, mem *[]int) *int {
	switch mode {
	case '0':
		return &(*mem)[*num]
	case '1':
		return num
	}

	return nil
}

func main() {
	//// part 1
	//load the program code
	mem := load_program("day5_input.txt")

	// iterate through the instructions and execute
	// format c = a $(op) b
	for pc := 0; pc < len(mem); {
		//convert instruction to string and prepend leading zeros, for easy to use later
		instr_str := strconv.Itoa(mem[pc])
		var sb strings.Builder
		for i := 0; i < 5-len(instr_str); i++ {
			sb.WriteString("0")
		}
		sb.WriteString(instr_str)
		instr_str = sb.String()
		fmt.Println("[>] instr_str: ", instr_str)
		//slice the last two, this is the opcode
		opcode, _ := strconv.Atoi(instr_str[3:])
		fmt.Println("[>] opcode: ", opcode)
		//calculate all pointers for parameters $a $b $c, nil if non applicable
		var a *int = nil
		var b *int = nil
		var c *int = nil
		switch opcode {
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 7:
			fallthrough
		case 8:
			// all three parameters
			a = get_val_p(&mem[pc+1], instr_str[2], &mem)
			b = get_val_p(&mem[pc+2], instr_str[1], &mem)
			c = get_val_p(&mem[pc+3], instr_str[0], &mem)
		case 3:
			fallthrough
		case 4:
			// only one parameters
			a = get_val_p(&mem[pc+1], instr_str[2], &mem)
		case 5:
			fallthrough
		case 6:
			// two parameters
			a = get_val_p(&mem[pc+1], instr_str[2], &mem)
			b = get_val_p(&mem[pc+2], instr_str[1], &mem)
		}

		//execute the instruction with regard to opcode
		switch opcode {
		case 1:
			// addition
			fmt.Println("[>] full instruction: ", mem[pc], *a, *b, mem[pc+3])
			*c = *a + *b
			pc += 4
		case 2:
			// multiply
			fmt.Println("[>] full instruction: ", mem[pc], *a, *b, mem[pc+3])
			*c = *a * *b
			pc += 4
		case 3:
			//input
			fmt.Println("[>] full instruction: ", mem[pc], mem[pc+1])
			fmt.Println("Please provide input: ")
			var input string
			_, err := fmt.Scanln(&input)
			if err != nil {
				fmt.Println(os.Stderr, err)
				return
			}
			mem[mem[pc+1]], _ = strconv.Atoi(input)
			pc += 2
		case 4:
			//output
			fmt.Println("[>] full instruction: ", mem[pc], *a)
			fmt.Println("[*] print: ", *a)
			pc += 2
		case 5:
			// jump-if-true
			fmt.Println("[>] full instruction: ", mem[pc], *a, *b)
			if *a != 0 {
				pc = *b
			} else {
				pc += 3
			}
		case 6:
			// jump-if-false
			fmt.Println("[>] full instruction: ", mem[pc], *a, *b)
			if *a == 0 {
				pc = *b
			} else {
				pc += 3
			}
		case 7:
			// less than
			fmt.Println("[>] full instruction: ", mem[pc], *a, *b, mem[pc+3])
			if *a < *b {
				*c = 1
			} else {
				*c = 0
			}
			pc += 4
		case 8:
			// equals
			fmt.Println("[>] full instruction: ", mem[pc], *a, *b, mem[pc+3])
			if *a == *b {
				*c = 1
			} else {
				*c = 0
			}
			pc += 4
		case 99:
			// halt
			//fmt.Println(mem[0])
			return
		}
	}
}
