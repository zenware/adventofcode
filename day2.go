package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Day2Report []int

func Day2ReportBuilder(input io.Reader) []Day2Report {
	scanner := bufio.NewScanner(input)
	var reports []Day2Report

	for scanner.Scan() {
		var report Day2Report
		str_levels := strings.Fields(scanner.Text())

		for _, s := range str_levels {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			report = append(report, num)
		}
		reports = append(reports, report)
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
