package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"
)

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
