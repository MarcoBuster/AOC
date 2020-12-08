package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Instruction struct {
	op  string
	arg int
}

func alreadyInArray(x *[]int, element int) bool {
	for _, c := range *x {
		if element == c {
			return true
		}
	}
	return false
}

func parseInstruction(line string) Instruction {
	op := line[0:3]
	arg, _ := strconv.Atoi(line[4:])
	return Instruction{op, arg}
}

func executePart1(code []Instruction) int {
	seen := make([]int, 0, len(code))
	programCounter := 0
	accumulator := 0
	var instruction Instruction
	for !alreadyInArray(&seen, programCounter) {
		seen = append(seen, programCounter)
		instruction = code[programCounter]
		switch instruction.op {
		case "nop":
			programCounter++
			break
		case "acc":
			programCounter++
			accumulator += instruction.arg
			break
		case "jmp":
			programCounter += instruction.arg
			break
		}
	}
	return accumulator
}

func executePart2(code []Instruction) int {
	seen := make([]int, 0, len(code))
	programCounter := 0
	accumulator := 0
	i := 0
	var instruction Instruction
	// You can start to feel that is very stupid by the next line
	for i < 10000 { // infinite loop check
		seen = append(seen, programCounter)
		if programCounter > len(code)-1 {
			break
		}
		instruction = code[programCounter]
		switch instruction.op {
		case "nop":
			programCounter++
			break
		case "acc":
			programCounter++
			accumulator += instruction.arg
			break
		case "jmp":
			programCounter += instruction.arg
			break
		}
		i++
	}
	if i == 10000 {
		return -1
	}
	return accumulator
}

// This is the DUMBEST method possible to solve the challenge.
// If you're looking for an elegant solution don't look mine.
func dumbBruteforce(code *[]Instruction, lastChanged *int) {
	var instruction Instruction
	var opcode string
	for i := *lastChanged; i < len(*code); i++ {
		instruction = (*code)[i]
		opcode = instruction.op
		if opcode == "jmp" {
			(*code)[i].op = "nop"
		} else if opcode == "nop" {
			(*code)[i].op = "jmp"
		}
		if (opcode == "jmp" || opcode == "nop") && i != *lastChanged {
			*lastChanged = i
			break
		}
	}
}

func main() {
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)
	var instructions []Instruction
	for b.Scan() {
		line := b.Text()
		instructions = append(instructions, parseInstruction(line))
	}
	part1 := executePart1(instructions)
	fmt.Println("Answer (part 1):", part1)

	// Change first instruction because dumbBruteforce changes it later.
	// So fucking dumb.
	if instructions != nil {
		firstOp := instructions[0].op
		if firstOp == "jmp" {
			instructions[0].op = "nop"
		} else if firstOp == "nop" {
			instructions[0].op = "jmp"
		}
	}

	// This is the worst AoC code I've ever written (so far).
	// It works, but please don't stare at it for too long
	// or you are going to have serious mental injuries.
	part2 := -1
	lastChanged := 0
	for part2 == -1 {
		dumbBruteforce(&instructions, &lastChanged)
		part2 = executePart2(instructions)
	}
	fmt.Println("Answer (part 2):", part2)
}
