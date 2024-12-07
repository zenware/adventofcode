package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

type Day6Map [][]rune

func Day6MapBuilder(input io.Reader) Day6Map {
	var aocmap Day6Map

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := scanner.Text()
		aocmap = append(aocmap, []rune(row))
	}
	return aocmap
}

// Find the guards position in the map.
// Return -1, -1 if not found.
func (aocmap Day6Map) FindGuardPosition() (int, int) {
	guard_symbols := []rune{'^', 'v', '<', '>'}
	for row := 0; row < len(aocmap); row++ {
		for col := 0; col < len(aocmap[row]); col++ {
			if slices.Contains(guard_symbols, aocmap[row][col]) {
				return row, col
			}
		}
	}
	return -1, -1
}

// Take a step in the direction the guard is facing, returning the new position
//
// If the gaurd would step into an obstruction, instead rotate 90 degrees.
// Since the guard is part of the map, rotating them requires modifying the map
// If the guard would walk off the map, return -1, -1
func (aocmap *Day6Map) GuardTakeStep(row int, col int) (int, int) {
	obstruction := '#'

	new_row, new_col := -1, -1

	switch (*aocmap)[row][col] {
	case '^': // North
		new_row, new_col := row-1, col
		// Stay in the same place, but rotate 90 degrees.
		if (*aocmap)[new_row][new_col] == obstruction {
			(*aocmap)[row][col] = '>'
			return row, col
		}
	case '>': // East
		new_row, new_col := row, col+1
		if (*aocmap)[new_row][new_col] == obstruction {
			(*aocmap)[row][col] = 'v'
			return row, col
		}
	case 'v': // South
		new_row, new_col := row+1, col
		if (*aocmap)[new_row][new_col] == obstruction {
			(*aocmap)[row][col] = '<'
			return row, col
		}
	case '<': // West
		new_row, new_col := row, col-1
		if (*aocmap)[new_row][new_col] == obstruction {
			(*aocmap)[row][col] = '^'
			return row, col
		}
	default:
		panic("invalid guard symbol")
	}

	rows, cols := len((*aocmap)), len((*aocmap)[0])
	if new_row < 0 || new_row > rows || new_col < 0 || new_col > cols {
		return -1, -1
	}

	return new_row, new_col
}

type MapPosition struct {
	row int
	col int
}

// Take steps until the guard walks off the map.
// Returning the number of unique positions the guard visited.
func (aocmap Day6Map) GuardWalk() int {
	// TODO: Make sure the for loop doesn't run forever.
	rows := len(aocmap)
	cols := len(aocmap[0])

	guard_row, guard_col := aocmap.FindGuardPosition()
	visited_positions := make(map[MapPosition]bool)
	visited_positions[MapPosition{guard_row, guard_col}] = true
	step_counter := 0
	total_area := rows * cols
	for {
		// if we have visited all the positions, we must necessarily be done.
		if step_counter >= total_area {
			if len(visited_positions) != total_area {
				fmt.Println("We looped enough times to visit all positions, but we didn't visit all of them.")
				fmt.Println("This means there is a bug in the code.")
				os.Exit(2)
			}
			break
		}
		guard_row, guard_col := aocmap.GuardTakeStep(guard_row, guard_col)
		step_counter++
		if guard_row < 0 || guard_row > rows || guard_col < 0 || guard_col > cols {
			break
		}
		visited_positions[MapPosition{guard_row, guard_col}] = true
	}
	return len(visited_positions)
}

// How many distinct positions will the guard visit before leaving the map?
// The map uses the symbols `.` for floor and `#` for obstructions.
// The guard uses the following symbols to indicate the direction they face:
//
//	`^` up, `v` down, `<` left, `>` right.
//
// The guard moves forward until they hit an obstruction, and turns 90 degrees.
// Eventually they will leave the map.
// How many distinct positions do they visit before leaving the map?
func Day6Puzzle1(aocmap Day6Map) int {
	//guard_row, guard_col := aocmap.FindGuardPosition()
	distinct_positions := aocmap.GuardWalk()
	return distinct_positions
}

func Day6Puzzle2() int {
	return 0
}

func Day6Puzzles() {
	file, err := os.Open("2024/day06/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(2)
	}
	defer file.Close()

	aocmap := Day6MapBuilder(file)

	fmt.Println("Day 6, Puzzle 1:", Day6Puzzle1(aocmap))
	fmt.Println("Day 6, Puzzle 2:", Day6Puzzle2())
}
