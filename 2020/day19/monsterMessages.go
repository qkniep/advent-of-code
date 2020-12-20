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
	var possible = make(map[string]bool, 0)
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

	// generate possible values
	possible = possibleValues(rules, 0)

	for scanner.Scan() {
		//if (ruleMatches(scanner.Text(), rules, 0)) {
		if possible[scanner.Text()] {
			matches++
		}
	}

	fmt.Printf("Matches: %v\n", matches)
}

func possibleValues(rules map[int]string, rule int) map[string]bool {
	fmt.Printf("Parse rule #%v.\n", rule)
	var possible = make(map[string]bool, 0)
	if rules[rule] == " \"a\"" || rules[rule] == " \"b\"" {
		possible[rules[rule][2:3]] = true
	} else if strings.Contains(rules[rule], "|") {
		for _, subRule := range strings.Split(rules[rule], "|") {
			for s := range concatPossibleValues(subRule, rules, rule) {
				possible[s] = true
			}
		}
	} else {
		for s := range concatPossibleValues(rules[rule], rules, rule) {
			possible[s] = true
		}
	}
	return possible
}

func concatPossibleValues(s string, rules map[int]string, rule int) map[string]bool {
	var subRules = strings.Fields(s)
	var possible = make([]map[string]bool, len(subRules))
	for i := range subRules {
		possible[i] = make(map[string]bool, 0)
	}
	for i, r := range subRules {
		n, _ := strconv.Atoi(r)
		for s := range possibleValues(rules, n) {
			possible[i][s] = true
		}
	}
	if len(subRules) == 1 {
		return possible[0]
	}
	if len(subRules) != 2 {
		panic("GRRRR!")
	}
	var result = make(map[string]bool, 0)
	for s1 := range possible[0] {
		for s2 := range possible[1] {
			result[s1+s2] = true
		}
	}
	return result
}

func ruleMatches(s string, rules map[int]string, ruleToMatch int) bool {
	return true
}
