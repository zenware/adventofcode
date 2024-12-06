package main

import (
	"fmt"
	"os"
)

func Day6Puzzle1() int {
	return 0
}

func Day6Puzzle2() int {
	return 0
}

func Day6Puzzles() {
	file, err := os.Open("2024/day06/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(2)
	}
	defer file.Close()

	rules, updates := Day5BuildPageOrderingRulesAndUpdates(file)

	fmt.Println("Day 6, Puzzle 1:", Day5Puzzle1(rules, updates))
	fmt.Println("Day 6, Puzzle 2:", Day5Puzzle2(rules, updates))
}
