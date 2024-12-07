package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

type Day3Instruction struct {
	Operation string
	Left      int
	Right     int
}

// In this case we're worried about corrupted memory.
// The only valid instructions read like: "mul(3,4)"
// If there are spaces, or different punctuation, then it's not valid.
func Day3InstructionBuilder(input io.Reader) []Day3Instruction {
	var instructions []Day3Instruction

	scanner := bufio.NewScanner(input)
	instruction_match := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	for scanner.Scan() {
		instruction_matches := instruction_match.FindAllString(scanner.Text(), -1)
		for _, match := range instruction_matches {
			//fmt.Println("Instruction Match:", match)
			switch match {
			case "do()":
				new_instruction := Day3Instruction{Operation: "do", Left: 0, Right: 0}
				instructions = append(instructions, new_instruction)
			case "don't()":
				new_instruction := Day3Instruction{Operation: "don't", Left: 0, Right: 0}
				instructions = append(instructions, new_instruction)
			default:
				new_instruction := Day3Instruction{Operation: "mul", Left: 0, Right: 0}
				fmt.Sscanf(match, "mul(%d,%d)", &new_instruction.Left, &new_instruction.Right)
				instructions = append(instructions, new_instruction)
			}
		}
	}

	return instructions
}

// Because why not?
type TinyVM struct {
	MulEnabled bool
}

// This is because I'm not certain that
func NewTinyVM() *TinyVM {
	return &TinyVM{MulEnabled: true}
}

func (vm *TinyVM) Execute(instructions []Day3Instruction) ([]int, error) {
	results := []int{}

	for _, instr := range instructions {
		switch instr.Operation {
		case "do":
			vm.MulEnabled = true
		case "don't":
			vm.MulEnabled = false
		case "mul":
			// Skip operation if not enabled.
			if vm.MulEnabled {
				results = append(results, instr.Left*instr.Right)
			}
		default:
			fmt.Printf("Unknown Operation: %s\n", instr.Operation)
		}
	}

	return results, nil
}

// I thought there would be multiple kinds of instruction to execute, and...
// there are! However, they also require state, and this isn't suitable.
// Therefore, this exists just for running mul instructions...
// and will remain here to keep the Day3Puzzle1 function intact for now.
func (instruction Day3Instruction) Execute() (int, error) {
	// Right now we only support multiplication.
	switch instruction.Operation {
	case "mul":
		return instruction.Left * instruction.Right, nil
	}
	// TODO: Handle returning an error.
	// Really return 0 means "Invalid" right now.
	// But that may not actually be the case.
	return 0, fmt.Errorf("invalid instruction: %s", instruction.Operation)
}

func Day3Puzzle1(instructions []Day3Instruction) int {
	total := 0
	for _, instruction := range instructions {
		// Despite not being able to execute "do" or "don't" instructions,
		// this will still return the correct result.
		result, err := instruction.Execute()
		if err != nil {
			//fmt.Println("Error executing instruction:", err)
			continue
		}
		total += result
	}
	return total
}

func Day3Puzzle2(instructions []Day3Instruction) int {
	total := 0

	// Using a smol VM to track whether mul is enabled.
	// Run all of the instructions and stash the results in a []int
	vm := NewTinyVM()
	results, err := vm.Execute(instructions)
	if err != nil {
		fmt.Println("Error executing instructions:", err)
		return 0
	}

	for _, num := range results {
		total += num
	}
	return total
}

func Day3Puzzles() {
	file, err := os.Open("2024/day03/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(2)
	}
	defer file.Close()

	instructions := Day3InstructionBuilder(file)
	fmt.Println("Day 3, Puzzle 1:", Day3Puzzle1(instructions))
	fmt.Println("Day 3, Puzzle 2:", Day3Puzzle2(instructions))
}
