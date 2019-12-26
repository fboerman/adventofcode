package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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

// run the Intcodemachine with persistent memory, so take in memory by reference (so with pointer)
// out: list of outputs
// state: true if halted, run false if ran out of inputs
// pc: last executed pc
func IntcodeMachine(mem_p *[]int, in []int, pc int) (out []int, state bool, pc_out int){
	mem := *mem_p
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
			//fmt.Println("[>] full instruction: ", mem[pc], *a, *b, mem[pc+3])
			*c = *a + *b
			pc += 4
		case 2:
			// multiply
			//fmt.Println("[>] full instruction: ", mem[pc], *a, *b, mem[pc+3])
			*c = *a * *b
			pc += 4
		case 3:
			//input
			//fmt.Println("[>] full instruction: ", mem[pc], mem[pc+1])
			//fmt.Println("Please provide input: ")
			//var input string
			//_, err := fmt.Scanln(&input)
			//if err != nil {
			//	fmt.Println(os.Stderr, err)
			//	return
			//}
			//mem[mem[pc+1]], _ = strconv.Atoi(input)
			if len(in) == 0 {
				state = false
				pc_out = pc
				return
			}
			inp := in[0]
			in = in[1:]
			mem[mem[pc+1]] = inp
			pc += 2
		case 4:
			//output
			//fmt.Println("[>] full instruction: ", mem[pc], *a)
			//fmt.Println("[*] print: ", *a)

			out = append(out, *a)
			pc += 2
		case 5:
			// jump-if-true
			//fmt.Println("[>] full instruction: ", mem[pc], *a, *b)
			if *a != 0 {
				pc = *b
			} else {
				pc += 3
			}
		case 6:
			// jump-if-false
			//fmt.Println("[>] full instruction: ", mem[pc], *a, *b)
			if *a == 0 {
				pc = *b
			} else {
				pc += 3
			}
		case 7:
			// less than
			//fmt.Println("[>] full instruction: ", mem[pc], *a, *b, mem[pc+3])
			if *a < *b {
				*c = 1
			} else {
				*c = 0
			}
			pc += 4
		case 8:
			// equals
			//fmt.Println("[>] full instruction: ", mem[pc], *a, *b, mem[pc+3])
			if *a == *b {
				*c = 1
			} else {
				*c = 0
			}
			pc += 4
		case 99:
			// halt
			//fmt.Println(mem[0])
			state = true
			pc_out = pc
			return
		}
	}

	return
}

type Result struct {
	settings []int
	output_signal int
}

// no feedback loop version
// run the program with the settings
func RunAmplifiers(settings []int, program string, c chan Result){
	signal := 0
	for i := 0; i<5;i++ {
		mem := load_program(program)
		out, _, _ := IntcodeMachine(&mem, []int{settings[i], signal}, 0)
		signal = out [0]
	}

	r := Result{settings, signal}
	c <- r
}

// from: https://www.golangprograms.com/golang-program-to-generate-slice-permutations-of-number-entered-by-user.html
func rangeSlice(start, stop int) []int {
	if start > stop {
		panic("Slice ends before it started")
	}
	xs := make([]int, stop-start)
	for i := 0; i < len(xs); i++ {
		xs[i] = i + start
	}
	return xs
}
func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}
// end copy

var wg sync.WaitGroup

func main() {
	//create all possible inputs and start a goroutine for each one
	c := make(chan Result)
	all_possible_settings := permutation(rangeSlice(0,5))
	for _, setting := range all_possible_settings {
		go RunAmplifiers(setting, "day7_input.txt", c)
	}

	// wait for all goroutines to finish
	wg.Wait()

	// retrieve all results and check which one has higest output
	highest_candidate := Result{[]int{}, 0}
	for i:=0; i<len(all_possible_settings); i++ {
		candidate := <- c
		if candidate.output_signal > highest_candidate.output_signal {
			highest_candidate = candidate
		}
	}
	fmt.Println("part 1:")
	fmt.Println("highest signal is ", highest_candidate.output_signal, " with setting ", highest_candidate.settings)

}