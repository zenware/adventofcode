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
