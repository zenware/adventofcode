package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func Day1ListBuilder(input io.Reader) ([]int, []int) {
	scanner := bufio.NewScanner(input)
	var leftlist []int
	var rightlist []int

	for scanner.Scan() {
		var leftitem int
		var rightitem int
		fmt.Sscanln(scanner.Text(), &leftitem, &rightitem)
		leftlist = append(leftlist, leftitem)
		rightlist = append(rightlist, rightitem)
	}
	return leftlist, rightlist
}

func Day1Puzzle1(leftlist []int, rightlist []int) int {
	slices.Sort(leftlist)
	slices.Sort(rightlist)

	total_distance := 0
	for i := 0; i < len(leftlist); i++ {
		distance := max(leftlist[i], rightlist[i]) - min(leftlist[i], rightlist[i])
		total_distance += distance
	}

	return total_distance
}

// Add up numbers in the left list, after multiplying them by occurrences in
// theright list.
func Day1Puzzle2(leftlist []int, rightlist []int) int {
	slices.Sort(leftlist)
	slices.Sort(rightlist)

	// Count how many times each number appears in the right list.
	occurrence_counter := make(map[int]int)
	for i := 0; i < len(rightlist); i++ {
		occurrence_counter[rightlist[i]]++
	}

	// Multiply each number in the left list by the times it appears in the
	// right list. And sum the results.
	total_similarity := 0
	for i := 0; i < len(leftlist); i++ {
		similarity := leftlist[i] * occurrence_counter[leftlist[i]]
		total_similarity += similarity
	}
	return total_similarity
}

func Day1Puzzles() {
	file, err := os.Open("2024/day01/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(2)
	}
	defer file.Close()

	leftlist, rightlist := Day1ListBuilder(file)

	fmt.Println("Day 1, Puzzle 1:", Day1Puzzle1(leftlist, rightlist))
	fmt.Println("Day 1, Puzzle 2:", Day1Puzzle2(leftlist, rightlist))
}
