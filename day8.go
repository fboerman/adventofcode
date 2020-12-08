package main

import (
	"bufio"
	"fmt"
	"os"
)

type Instruction struct {
	action string
	arg    int
}

func run(Imem []Instruction, registers map[string]int, pc int, adjustment int) (int, map[string]int) {
	// return flags:
	// 1: terminated due to repeating instruction
	// 2: terminated normally
	visited_flags := map[int]bool{}
	adjusted_flag := false
	for !visited_flags[pc] {
		if pc >= len(Imem) {
			// terminated
			return 2, registers
		}
		instr := Imem[pc]
		visited_flags[pc] = true
		if pc == adjustment && !adjusted_flag {
			switch instr.action {
			case "nop":
				instr.action = "jmp"
				break
			case "jmp":
				instr.action = "nop"
				break
			}
			adjusted_flag = true
		}
		switch instr.action {
		case "nop":
			// do nothing
			break
		case "acc":
			registers["acc"] += instr.arg
			break
		case "jmp":
			pc += instr.arg
			break
		}
		if instr.action != "jmp" {
			pc++
		}
	}
	return 1, registers
}

func main() {
	file, err := os.Open("day8_input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var Imem []Instruction

	for scanner.Scan() {
		line := scanner.Text()
		inst := Instruction{}
		_, _ = fmt.Sscanf(line, "%s %d", &inst.action, &inst.arg)
		Imem = append(Imem, inst)
	}

	fmt.Println("Size of Imem:", len(Imem))
	flag, registers := run(Imem, map[string]int{}, 0, -1)
	fmt.Printf("original run, acc: %d, flag: %d\n", registers["acc"], flag)

	// create a list of all jmp and nop operations and then run each one mutated
	var adjustments []int
	for i := 0; i < len(Imem); i++ {
		switch Imem[i].action {
		case "jmp":
			adjustments = append(adjustments, i)
			break
		case "nop":
			if Imem[i].arg != 0 {
				adjustments = append(adjustments, i)
			}
			break
		}
	}
	for _, adj := range adjustments {
		flag, registers := run(Imem, map[string]int{}, 0, adj)
		if flag == 2 {
			fmt.Printf("adjusted run, acc: %d, flag: %d, changed instruction: %d\n", registers["acc"], flag, adj)
		}
	}

}
