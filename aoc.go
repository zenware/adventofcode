package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open 2024/day01/input.txt
	file, err := os.Open("2024/day01/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(2)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(2)
		return
	}

	fmt.Println(lines)
}
