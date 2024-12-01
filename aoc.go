package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
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
	var leftlist []string
	var rightlist []string

	for scanner.Scan() {
		split_lists := strings.Split(scanner.Text(), " ")
		leftlist = append(leftlist, split_lists[0])
		rightlist = append(rightlist, split_lists[1])
	}
	fmt.Println(rightlist)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(2)
		return
	}

	sort.Slice(leftlist, func(i, j int) bool {
		return leftlist[i] < leftlist[j]
	})

	//slices.Sort(rightlist)
	//sort.Sort(rightlist)
	slices.SortFunc(rightlist, func(a, b string) int {
		return strings.Compare(a, b)
	})

	//slices.Reduce(leftlist, func(a, b string) string {})

	slices.Compare(leftlist, rightlist)
	fmt.Println(rightlist)
}
