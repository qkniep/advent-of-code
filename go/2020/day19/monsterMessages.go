package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var rules = make(map[int]string, 128)
	var matches = 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		fields := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(fields[0][:len(fields[0])-1])
		rules[n] = scanner.Text()[len(fields[0]):]
	}

	var inputs []string
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	for _, s := range inputs {
		if ruleMatchesFullStr(s, rules, 0) {
			matches++
		}
	}
	fmt.Printf("Matches: %v\n", matches)

	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"

	matches = 0
	for _, s := range inputs {
		if ruleMatchesFullStr(s, rules, 0) {
			matches++
		}
	}
	fmt.Printf("Matches with recursive rules: %v\n", matches)
}

// Returns true iff the full string `s` exactly matches the given rule.
func ruleMatchesFullStr(s string, rules map[int]string, ruleToMatch int) bool {
	var finalIndices = matchRule(s, rules, ruleToMatch, 0)
	for _, index := range finalIndices {
		if index == len(s) {
			return true
		}
	}
	return false
}

// Checks whether a given rule could match the string `s` starting at index `from`.
// Returns a list of all possible indices where we could end up after matching the rule,
// this is list is empty iff it's impossible to match the rule at that index.
func matchRule(s string, rules map[int]string, rule int, from int) []int {
	if from >= len(s) {
		return []int{}
	}
	// check base rules (match a single character)
	if rules[rule] == " \"a\"" || rules[rule] == " \"b\"" {
		if s[from] == rules[rule][2] {
			return []int{from + 1}
		}
		return []int{}
	}
	// check all other rules
	var finalIndices []int
	for _, subRule := range strings.Split(rules[rule], "|") {
		oldIndices := []int{from}
		for _, ruleStr := range strings.Fields(subRule) {
			ruleInt, _ := strconv.Atoi(ruleStr)
			newIndices := []int{}
			for _, index := range oldIndices {
				newIndices = append(newIndices, matchRule(s, rules, ruleInt, index)...)
			}
			oldIndices = newIndices
		}
		finalIndices = append(finalIndices, oldIndices...)
	}
	return finalIndices
}
