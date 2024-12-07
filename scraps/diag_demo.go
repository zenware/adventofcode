package main

import (
	"fmt"
	"sort"
)

func main() {
	grid := [][]rune{
		{'A', 'B', 'C', 'D', 'E'},
		{'F', 'G', 'H', 'I', 'J'},
		{'K', 'L', 'M', 'N', 'O'},
		{'P', 'Q', 'R', 'S', 'T'},
	}

	diagonals := extractTopLeftToBottomRightDiagonals(grid, 200)
	for i, diag := range diagonals {
		fmt.Printf("Diagonal %d: %s\n", i+1, diag)
	}
}

func extractTopLeftToBottomRightDiagonals(grid [][]rune, min_len int) []string {
	rows := len(grid)
	if rows == 0 {
		return nil
	}
	cols := len(grid[0])
	diagonalsMap := make(map[int][]rune)

	if min_len > cols {
		min_len = cols
	}

	// Iterate over the grid and group characters by diagonal key (i - j)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols-(cols-min_len); j++ {
			key := i + j
			diagonalsMap[key] = append(diagonalsMap[key], grid[i][j])
		}
	}

	// Collect keys and sort them to maintain order
	keys := make([]int, 0, len(diagonalsMap))
	for k := range diagonalsMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Extract diagonals in order
	var diagonals []string
	for _, k := range keys {
		diagonal := string(diagonalsMap[k])
		diagonals = append(diagonals, diagonal)
	}

	return diagonals
}
