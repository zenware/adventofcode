package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
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

type Day2Report []int

func Day2ReportBuilder(input io.Reader) []Day2Report {
	scanner := bufio.NewScanner(input)
	var reports []Day2Report

	for scanner.Scan() {
		var r Day2Report
		str_levels := strings.Fields(scanner.Text())

		for _, s := range str_levels {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			r = append(r, num)
		}
		reports = append(reports, r)
	}
	return reports
}

// Verifies the following constraints:
// - All levels are either increasing or decreasing
// - Any two adjacent levels differ by at LEAST one and at MOST three
// Technically it is more efficient, and possible, to check all of these properties
// in one pass, but I don't wanna.
func (report Day2Report) IsSafe() bool {
	isIncreasing := report[1] > report[0]

	for i := 1; i < len(report); i++ {
		abs_diff := max(report[i-1], report[i]) - min(report[i-1], report[i])
		// Differences must be AT LEAST one and AT MOST three.
		if abs_diff < 1 || abs_diff > 3 {
			return false
		}

		if isIncreasing {
			if report[i-1] >= report[i] {
				return false
			}
		} else {
			if report[i-1] <= report[i] {
				return false
			}
		}

	}
	return true
}

// How many reports are safe?
func Day2Puzzle1(reports []Day2Report) int {
	safe_report_count := 0

	for i := 0; i < len(reports); i++ {
		if reports[i].IsSafe() {
			safe_report_count++
		}
	}

	return safe_report_count
}

// Brute force solution to Day 2, Puzzle 2.
// Allows the reactor/record to safely tolerate a single bad level.
func (report Day2Report) RemoveFromIndex(index int) Day2Report {
	// https://stackoverflow.com/a/57213476
	ret := make(Day2Report, 0, len(report)-1)
	ret = append(ret, report[:index]...)
	return append(ret, report[index+1:]...)
}

func Day2Puzzle2(reports []Day2Report) int {
	safe_report_count := 0

	for i := 0; i < len(reports); i++ {
		// If a report is safe to begin with, then add it to the count and move on.
		if reports[i].IsSafe() {
			safe_report_count++
			continue
		}
		// Problem Dampener
		// Generate reports with one level removed and check if they are safe.
		// If the are, then add them to the count.
		for j := 0; j < len(reports[i]); j++ {
			dampened_report := make(Day2Report, len(reports[i])-1)
			copy(dampened_report, reports[i].RemoveFromIndex(j))
			if dampened_report.IsSafe() {
				safe_report_count++
				break // Prevent a single report from being counted twice.
			}
		}
	}

	return safe_report_count
}

func Day2Puzzles() {
	file, err := os.Open("2024/day02/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(2)
	}
	defer file.Close()

	reports := Day2ReportBuilder(file)
	fmt.Println("Day 2, Puzzle 1:", Day2Puzzle1(reports))
	fmt.Println("Day 2, Puzzle 2:", Day2Puzzle2(reports))
}

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

type Day4WordSearch [][]rune

// Count the occurrences of a word in the word search.
//
// Words can be found horizontal, vertical, diagonal, backwards, and overlapping.
// The following is an example of 2 occurrences of "XMAS"
// ..X.
// SAMX
// ..A.
// ..S.
func (wordsearch Day4WordSearch) CountWordOccurrences(word string) int {
	count := 0
	lines := wordsearch.ExtractLines(len(word))

	xmax_match := regexp.MustCompile(`XMAS`)

	for _, line := range lines {
		forward_xmas_matches := xmax_match.FindAllString(line, -1)
		// NOTE: Hopefully casting types is enough to not be a reference.
		revline := []rune(line)
		slices.Reverse(revline)
		reverse_xmas_matches := xmax_match.FindAllString(string(revline), -1)
		count += len(forward_xmas_matches) + len(reverse_xmas_matches)
	}

	return count
}

// Converts the word search grid into a []string of linearly searchable lines.
//
// Horizontal, Vertical, Left->Right Diagonal, Right->Left Diagonal.
// min_length is the minimum length of a line to be extracted.
// - This is used to prevent storing diagonal lines of length 1
// TODO: Actually use the length to cut corners.
func (wordsearch Day4WordSearch) ExtractLines(min_length int) []string {
	var lines []string
	rows := len(wordsearch)
	cols := len(wordsearch[0])

	// Extract Horizontal Lines (rows)
	for row := 0; row < rows; row++ {
		lines = append(lines, string(wordsearch[row]))
	}

	// Extract Vertical Lines (columns)
	for col := 0; col < cols; col++ {
		var sb strings.Builder // I kind of hate this, but ¯\_(ツ)_/¯
		for row := 0; row < rows; row++ {
			sb.WriteRune(wordsearch[row][col])
		}
		lines = append(lines, sb.String())
	}

	// Because maps magically make solutions O(1)
	diagMap := make(map[int][]rune)
	// TODO: This is still just here for no reason, with no corner cutting.
	if min_length > cols {
		// Ensure we don't go out of bounds when cutting corners.
		min_length = cols
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Get both diagonals at once!
			toplbotrDiagKey := row - col
			toprbotlDiagKey := row + col
			diagMap[toplbotrDiagKey] = append(diagMap[toplbotrDiagKey], wordsearch[row][col])
			diagMap[toprbotlDiagKey] = append(diagMap[toprbotlDiagKey], wordsearch[row][col])
		}
	}
	// Kludge to get the keys in order.
	// https://go.dev/blog/maps#iteration-order
	keys := make([]int, 0, len(diagMap))
	for k := range diagMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		diag := string(diagMap[key])
		lines = append(lines, diag)
	}
	return lines
}

func (wordsearch Day4WordSearch) Day4FindMASCrossings() int {
	count := 0
	rows := len(wordsearch)
	cols := len(wordsearch[0])

	// We aren't actually getting an empty grid in practice, but just in case.
	if rows == 0 && cols == 0 {
		return 0
	}

	// Find all the "A"'s, and check their diags
	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			// Skip non-A's
			if wordsearch[row][col] != 'A' {
				continue
			}
			// Test the diagonals.
			//var tlbrDiag, trblDiag []rune
			tlbrDiag := string([]rune{wordsearch[row-1][col-1], wordsearch[row][col], wordsearch[row+1][col+1]})
			trblDiag := string([]rune{wordsearch[row-1][col+1], wordsearch[row][col], wordsearch[row+1][col-1]})

			validTlbr := tlbrDiag == "MAS" || tlbrDiag == "SAM"
			validTrbl := trblDiag == "MAS" || trblDiag == "SAM"

			if validTlbr && validTrbl {
				count++
			}
		}
	}
	return count
}

func Day4WordSearchBuilder(input io.Reader) Day4WordSearch {
	var wordsearch Day4WordSearch
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		row := scanner.Text()
		wordsearch = append(wordsearch, []rune(row))
	}
	return wordsearch
}

// How many times does XMAS appear in the word search?
func Day4Puzzle1(wordsearch Day4WordSearch) int {
	return wordsearch.CountWordOccurrences("XMAS")
}

// How many times does "MAS" cross itself in the word search?
// :/ this has come as a surprise.
// My best guess for a shortcut is to find all the "A"'s, and check their diags
func Day4Puzzle2(wordsearch Day4WordSearch) int {
	return wordsearch.Day4FindMASCrossings()
}

func Day4Puzzles() {
	file, err := os.Open("2024/day04/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(2)
	}
	defer file.Close()

	wordsearch := Day4WordSearchBuilder(file)
	fmt.Println("Day 4, Puzzle 1:", Day4Puzzle1(wordsearch))
	fmt.Println("Day 4, Puzzle 2:", Day4Puzzle2(wordsearch))
}

func main() {
	//Day1Puzzles()
	//Day2Puzzles()
	//Day3Puzzles()
	Day4Puzzles()
}
