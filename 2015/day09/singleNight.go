package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var cities []string
	var distances = make(map[string]map[string]int, 0)
	var shortestRoute, longestRoute int
	var scanner = bufio.NewScanner(os.Stdin)

	// read distances between cities
	for scanner.Scan() {
		parsed := strings.Fields(scanner.Text())
		origin, destination, distStr := parsed[0], parsed[2], parsed[4]
		if addToSet(&cities, origin) {
			distances[origin] = make(map[string]int, 0)
		}
		if addToSet(&cities, destination) {
			distances[destination] = make(map[string]int, 0)
		}
		distance, _ := strconv.Atoi(distStr)
		distances[origin][destination] = distance
		distances[destination][origin] = distance
	}

	shortestRoute, longestRoute = optimalRoutes(cities, distances)

	fmt.Println("Length of shortest route:", shortestRoute)
	fmt.Println("Length of longest route:", longestRoute)
}

// Tries all permutations, thus takes O(n!) time.
// TODO use O(n^2*2^n) algorithm for finding hammilton cycles
func optimalRoutes(cities []string, distances map[string]map[string]int) (int, int) {
	var min, max = 999999, 0
	perm(cities, func(cs []string) {
		var sum int
		for i := 1; i < len(cs); i++ {
			sum += distances[cs[i-1]][cs[i]]
		}

		if sum < min {
			min = sum
		}
		if sum > max {
			max = sum
		}
	}, 0)
	return min, max
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
