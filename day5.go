package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

// In text specified like `12|7` means that page 12 must be before page 7.
type Day5PageOderingRule struct {
	EarlierPageNumber int
	LaterPageNumber   int
}

// Since there will be many of them.
type Day5PageOrderingRules []Day5PageOderingRule

// List of pages indicating an update.
// We will need to use the Day5PageOrderingRules to determine if these are
// in the correct order.
// Ex: `7,9,12,45,3`
type Day5PageUpdate []int
type Day5PageUpdates []Day5PageUpdate

// Determine if a page update follows the ordering rules.
// If a specific page number does not appear in the update, then related rules
// can be ignored.
// That does hint that perhaps another map would be useful.
func (update Day5PageUpdate) FollowsRules(rules Day5PageOrderingRules) bool {
	// For each int in the update, check if it matches any of the rules.
	//
	// If it does match, then check that any subsequent ints, don't violate the
	// rules it matches.
	//
	// I do actually want to build a map, keyed by earlier page number, and
	// append the later page number(s) as an []int. That way I can easily check
	// if a page number is even in the map, and if not skip it. But if so, then
	// I can check the following pages to see if they are in its list, if any
	// of them are, then it's not valid.
	rules_map := make(map[int][]int)
	for _, rule := range rules {
		rules_map[rule.EarlierPageNumber] = append(rules_map[rule.EarlierPageNumber], rule.LaterPageNumber)
	}

	// Reverse the update, and iterate from the end.
	// I think checking the update backwards makes it work easier with our map.
	//
	// Each time we find a page number that matches a rule,
	// - The rule indicates a list of pages that follow it.
	// - So if any of those pages appear in a previous page, we're not valid.
	for i := len(update) - 1; i >= 0; i-- {
		// Ignore page numbers without rules.
		rules := rules_map[update[i]]
		if rules == nil {
			continue
		}

		// Check all previous pages to see if they are contained in the rules
		// For real data:
		// If we have []int{1, 2, 3}
		// I starts at 2, then goes to 1-1, then to 0-1
		// So I can't use -1 here, but for some reason I feel like I have to.
		// I didn't have to do that, so I'm not sure why I was thinking it.
		for _, prev_page := range update[0:i] {
			// PrevPage in this case is the contents of the update slice.
			// I was trying to use it as an index like update[prev_page]
			if slices.Contains(rules, prev_page) {
				return false
			}
		}
	}

	return true
}

// TODO: Not currently certain of Go division behavior.
// https://go.dev/ref/spec#Arithmetic_operators
// Seems like the types themselves might specify the division behavior.
// In this case, len() returns an int, so the division is int / int.
// NOTE: Now that this is working I'm more certain of the division behavior.
func (update Day5PageUpdate) GetMiddlePageNumber() int {
	return update[len(update)/2]
}

// This reads in a text file like:
// ```12|7\n\n7, 12\n```
// A blank line separates the rules from the updates.
func Day5BuildPageOrderingRulesAndUpdates(input io.Reader) (Day5PageOrderingRules, Day5PageUpdates) {
	var pageOrderingRules Day5PageOrderingRules
	var pageUpdates Day5PageUpdates

	var lines []string
	scanner := bufio.NewScanner(input)
	// Scan rules until the first blank line and then scan updates.
	// Or read them all into memory first as []string and then split might be easier.
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	rules_parsed := false
	for _, line := range lines {
		if line == "" {
			rules_parsed = true
			continue
		}

		if !rules_parsed { // Parse rules
			var pageOderingRule Day5PageOderingRule
			fmt.Sscanf(line, "%d|%d", &pageOderingRule.EarlierPageNumber, &pageOderingRule.LaterPageNumber)
			pageOrderingRules = append(pageOrderingRules, pageOderingRule)
			continue
		}

		var update Day5PageUpdate
		for _, s := range strings.Split(line, ",") {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			update = append(update, num)
		}
		pageUpdates = append(pageUpdates, update)

	}

	return pageOrderingRules, pageUpdates
}

// Find the middle number of each correctly ordered page update.
// Do we care about cycles in the rules? I don't think I see anything specified.
// - Although, if there is a cycle, it could cause some trickery.
func Day5Puzzle1(rules Day5PageOrderingRules, updates Day5PageUpdates) int {

	var correctly_ordered_updates Day5PageUpdates

	// I would like to have some convenient way of compiling the rules and
	// running a simple check. But I can't think of a way to do that off the
	// top of my head.
	// So I'm pretty sure that means, for each update, we need to check all the
	// rules to see if it follows them.
	// I'm sure there's a better way to do this.
	for _, update := range updates {
		if update.FollowsRules(rules) {
			correctly_ordered_updates = append(correctly_ordered_updates, update)
		}
	}

	// This part is mostly contrived, just get the middle element.
	sum := 0
	for _, update := range correctly_ordered_updates {
		sum += update.GetMiddlePageNumber()
	}

	return sum
}

func Day5Puzzle2(rules Day5PageOrderingRules, updates Day5PageUpdates) int {
	return 0
}

func Day5Puzzles() {
	file, err := os.Open("2024/day05/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(2)
	}
	defer file.Close()

	rules, updates := Day5BuildPageOrderingRulesAndUpdates(file)

	fmt.Println("Day 5, Puzzle 1:", Day5Puzzle1(rules, updates))
	fmt.Println("Day 5, Puzzle 2:", Day5Puzzle2(rules, updates))
}
