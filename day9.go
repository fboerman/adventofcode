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
func get_val_p(num *int, mode byte, mem *[]int, reg *[]int) *int {
	switch mode {
	case '0':
		return &(*mem)[*num]
	case '1':
		return num
	case '2':
		return &(*mem)[*num + (*reg)[0]]
	}

	return nil
}

// run the Intcodemachine with persistent memory, so take in memory by reference (so with pointer)
// out: list of outputs
// state: true if halted, run false if ran out of inputs
// pc: last executed pc
// registers are defined as:
//     0: relative base
func IntcodeMachine(mem_p *[]int, reg_p *[]int, in []int, pc int) (out []int, state bool, pc_out int){
	// resolve pointer to memory for easy of writing
	mem := *mem_p
	reg := *reg_p
	state = false

	// iterate through the instructions and execute
	// format c = a $(op) b
	for pc < len(mem) {
		//convert instruction to string and prepend leading zeros, for easy to use later
		instr_str := strconv.Itoa(mem[pc])
		var sb strings.Builder
		for i := 0; i < 5-len(instr_str); i++ {
			sb.WriteString("0")
		}
		sb.WriteString(instr_str)
		instr_str = sb.String()
		//fmt.Println("[>] instr_str: ", instr_str)
		//slice the last two, this is the opcode
		opcode, _ := strconv.Atoi(instr_str[3:])
		//fmt.Println("[>] opcode: ", opcode)
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
			a = get_val_p(&mem[pc+1], instr_str[2], &mem, &reg)
			b = get_val_p(&mem[pc+2], instr_str[1], &mem, &reg)
			c = get_val_p(&mem[pc+3], instr_str[0], &mem, &reg)
			//fmt.Println("[>] full instruction: ", mem[pc], *a, *b, mem[pc+3])
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 9:
			// only one parameters
			a = get_val_p(&mem[pc+1], instr_str[2], &mem, &reg)
			//fmt.Println("[>] full instruction: ", mem[pc], *a)
		case 5:
			fallthrough
		case 6:
			// two parameters
			a = get_val_p(&mem[pc+1], instr_str[2], &mem, &reg)
			b = get_val_p(&mem[pc+2], instr_str[1], &mem, &reg)
			//fmt.Println("[>] full instruction: ", mem[pc], *a, *b)
		}

		//execute the instruction with regard to opcode
		switch opcode {
		case 1:
			// addition
			*c = *a + *b
			pc += 4
		case 2:
			// multiply
			*c = *a * *b
			pc += 4
		case 3:
			//input

			if len(in) == 0 {
				state = false
				pc_out = pc
				return
			}
			inp := in[0]
			in = in[1:]
			*a = inp
			pc += 2
		case 4:
			//output
			out = append(out, *a)
			pc += 2
		case 5:
			// jump-if-true
			if *a != 0 {
				pc = *b
			} else {
				pc += 3
			}
		case 6:
			// jump-if-false
			if *a == 0 {
				pc = *b
			} else {
				pc += 3
			}
		case 7:
			// less than
			if *a < *b {
				*c = 1
			} else {
				*c = 0
			}
			pc += 4
		case 8:
			// equals
			if *a == *b {
				*c = 1
			} else {
				*c = 0
			}
			pc += 4
		case 9:
			// adjust relative base
			reg[0] += *a
			pc += 2
		case 99:
			// halt
			state = true
			pc_out = pc
			return
		}
	}

	return
}

func main() {
	//load the program
	program := load_program("day9_input.txt")
	//allocate large memory space and copy program in it
	mem := make([]int, 1e9)
	copy(mem, program)
	//allocate registers space
	reg := make([]int, 1)

	out, _, _ := IntcodeMachine(&mem, &reg, []int{1}, 0)

	fmt.Println(("Part 1:"))
	fmt.Println(out)

	//reset the memory and registers
	mem = make([]int, 1e9)
	copy(mem, program)
	reg = make([]int, 1)

	//run now with new input
	out, _, _ = IntcodeMachine(&mem, &reg, []int{2}, 0)
	fmt.Println("Part 2:")
	fmt.Println(out)
}
