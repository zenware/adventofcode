package main

import (
	"strings"
	"testing"
)

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
		result := Day1Puzzle1(strings.NewReader(tc.inputString))
		if result != tc.expected {
			t.Errorf("Day1Puzzle1(%s) returned %d, expected %d", tc.inputString, result, tc.expected)
		}
	}
}
