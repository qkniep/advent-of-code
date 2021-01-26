package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const inFormat = "%s would %s %d happiness units by sitting next to %s"

// This solution is for the most part copied directly from day 9.
func main() {
	var people []string
	var happiness = make(map[string]map[string]int, 0)
	var scanner = bufio.NewScanner(os.Stdin)

	// read input and create happiness graph
	for scanner.Scan() {
		var from, to, verb string
		var amount int

		input := strings.TrimRight(scanner.Text(), ".")
		fmt.Sscanf(input, inFormat, &from, &verb, &amount, &to)
		if addToSet(&people, from) {
			happiness[from] = make(map[string]int, 0)
		}
		if addToSet(&people, to) {
			happiness[to] = make(map[string]int, 0)
		}
		if verb == "gain" {
			happiness[from][to] += amount
			happiness[to][from] += amount
		} else if verb == "lose" {
			happiness[from][to] -= amount
			happiness[to][from] -= amount
		}
	}

	maxHappiness1, maxHappiness2 := findArrangement(people, happiness)
	fmt.Println("Max. difference in happiness w/o seating self:", maxHappiness1)
	fmt.Println("Max. difference in happiness including self:", maxHappiness2)
}

// Tries all permutations, thus takes O(n!) time.
// TODO use O(n^2*2^n) algorithm for finding hammilton cycles
func findArrangement(people []string, happiness map[string]map[string]int) (int, int) {
	var maxPath, maxCycle int
	perm(people, func(cs []string) {
		var sum, minEdge int
		for i := 1; i < len(cs); i++ {
			edgeScore := happiness[cs[i-1]][cs[i]]
			sum += edgeScore
			if edgeScore < minEdge {
				minEdge = edgeScore
			}
		}

		if sum > maxPath {
			maxPath = sum
			maxCycle = sum + happiness[cs[0]][cs[len(cs)-1]]
		}
	}, 0)
	return maxCycle, maxPath
}

// Runs a function for every permutation of the given array's elements.
// source: https://yourbasic.org/golang/generate-permutation-slice-string/
func perm(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// Adds a new string to the set of strings (union).
// Returns true if the set was changed, false if the element already was in the set.
func addToSet(ss *[]string, s string) bool {
	var inSet bool
	for _, entry := range *ss {
		if entry == s {
			inSet = true
		}
	}
	if !inSet {
		*ss = append(*ss, s)
	}
	return !inSet
}
