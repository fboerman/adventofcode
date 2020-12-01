package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

// from https://stackoverflow.com/questions/39868029/how-to-generate-a-sequence-of-numbers-in-golang
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}


func load_program() (input []int) {
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
	file, _ := os.Open("day2_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(onComma)

	for scanner.Scan() {
		inp, _ := strconv.Atoi(scanner.Text())
		input = append(input, inp)
	}

	return
}

func execute(noun int, verb int, C int, c chan int)  {
	// load the program memory to make sure this goroutine has a seperate copy
	mem := load_program()
	// place the noun and verb
	mem[1] = noun
	mem[2] = verb
	// iterate through the instructions and execute
	// format c = a $(op) b
	for pc:=0; pc < len(mem)-4; pc+=4 {
		if mem[pc+1] >= len(mem) || mem[pc+2] >= len(mem) || mem[pc+3] >= len(mem){
			return
		}
		a := mem[mem[pc+1]]
		b := mem[mem[pc+2]]
		c := &mem[mem[pc+3]]

		switch opcode := mem[pc]; opcode {
		case 1:
			// addition
			*c = a + b
		case 2:
			// multiply
			*c = a * b
		case 99:
			// halt
			break
		}
	}
	//fmt.Println(noun, verb, mem[0])
	if mem[0] == C {
		c <- noun
		c <- verb
	}
}

var wg sync.WaitGroup

func main() {
	//// part 1
	//load the program code
	mem := load_program()

	// adjust the program code as per instruction of the puzzle
	mem[1] = 12
	mem[2] = 2

	// iterate through the instructions and execute
	// format c = a $(op) b
	for pc:=0; pc < len(mem)-4; pc+=4 {
		a := mem[mem[pc+1]]
		b := mem[mem[pc+2]]
		c := &mem[mem[pc+3]]

		switch opcode := mem[pc]; opcode {
		case 1:
			// addition
			*c = a + b
		case 2:
			// multiply
			*c = a * b
		case 99:
			// halt
			break
		}
	}

	fmt.Println("part 1:")
	fmt.Println(mem[0])

	// part 2
	fmt.Println("part 2:")
	fmt.Print("Enter constant to search for: ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(os.Stderr, err)
		return
	}
	magicC, _ := strconv.Atoi(input)
	fmt.Println("Searching for ", magicC)
	c := make(chan int)
	for noun := range makeRange(0, 99) {
		for verb := range makeRange(0, 99) {
			go execute(noun, verb, magicC, c)
		}
	}
	// wait for all goroutines
	wg.Wait()
	// retrieve the result. the only data in this channel should be the answers
	noun, verb := <- c, <- c

	fmt.Println(noun, verb, 100*noun+verb)
}