package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
	// Open 2024/day01/input.txt
	file, err := os.Open("2024/day01/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(2)
		return
	}

	return file, nil
}
*/

func Day1Puzzle1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	var leftlist []int
	var rightlist []int

	input_lines := 0
	for scanner.Scan() {
		var leftitem int
		var rightitem int
		fmt.Sscanln(scanner.Text(), &leftitem, &rightitem)
		leftlist = append(leftlist, leftitem)
		rightlist = append(rightlist, rightitem)
		input_lines++
	}

	slices.Sort(leftlist)
	slices.Sort(rightlist)

	if slices.Min(leftlist) != leftlist[0] {
		fmt.Println("Something has gone wrong")
	} else {
		fmt.Println("Something has gone right")
	}

	total_distance := 0
	for i := 0; i < len(leftlist); i++ {
		distance := max(leftlist[i], rightlist[i]) - min(leftlist[i], rightlist[i])
		total_distance += distance
	}

	fmt.Println("Total Distance:", total_distance)

	return total_distance
}

func main() {
	file, err := os.Open("2024/day01/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(2)
	}
	defer file.Close()

	fmt.Println("Day 1, Puzzle 1:", Day1Puzzle1(file))
}
