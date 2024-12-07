package main

import (
	"fmt"
	"os"
)

func Day7Puzzle1(aocmap Day6Map) int {
	return 0
}

func Day7Puzzle2() int {
	return 0
}

func Day7Puzzles() {
	file, err := os.Open("2024/day07/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(2)
	}
	defer file.Close()

	fmt.Println("Day 7, Puzzle 1:", Day7Puzzle1())
	fmt.Println("Day 7, Puzzle 2:", Day7Puzzle2())
}
