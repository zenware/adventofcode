package main

import (
	"strings"
	"testing"
)

// Distances
func TestDay1Puzzle1(t *testing.T) {
	testCases := []struct {
		inputString string
		expected    int
	}{
		{`3   4
4   3
2   5
1   3
3   9
3   3`, 11},
		{`3   3`, 0},
		{`3   4`, 1},
	}

	for _, tc := range testCases {
		leftlist, rightlist := Day1ListBuilder(strings.NewReader(tc.inputString))
		result := Day1Puzzle1(leftlist, rightlist)
		if result != tc.expected {
			t.Errorf("Day1Puzzle1(%s) returned %d, expected %d", tc.inputString, result, tc.expected)
		}
	}
}

// Similarity Scores
func TestDay1Puzzle2(t *testing.T) {
	testCases := []struct {
		inputString string
		expected    int
	}{
		{`3   4
4   3
2   5
1   3
3   9
3   3`, 31},
		{`3   3`, 3},
		{`3   4`, 0},
	}

	for _, tc := range testCases {
		leftlist, rightlist := Day1ListBuilder(strings.NewReader(tc.inputString))
		result := Day1Puzzle2(leftlist, rightlist)
		if result != tc.expected {
			t.Errorf("Day1Puzzle1(%s) returned %d, expected %d", tc.inputString, result, tc.expected)
		}
	}
}

// Report Safety Counts
func TestDay2Puzzle1(t *testing.T) {
	testCases := []struct {
		inputString string
		expected    int
	}{
		{`7 6 4 2 1 
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`, 2},
		{`1 9 1 9 1`, 0},
		{`1 1 1 1 1`, 0},
		{`1 2 3 4 5`, 1},
		{`5 4 3 2 1`, 1},
		{`2 3 2 4 5`, 0},
		{`9 2 3 4 5
1 9 3 4 5
1 2 9 4 5
1 2 3 9 5
1 2 3 4 9`, 0},
	}

	for _, tc := range testCases {
		reports := Day2ReportBuilder(strings.NewReader(tc.inputString))
		result := Day2Puzzle1(reports)
		if result != tc.expected {
			t.Errorf("Day2Puzzle2(%s) returned %d, expected %d", tc.inputString, result, tc.expected)
		}
	}
}
