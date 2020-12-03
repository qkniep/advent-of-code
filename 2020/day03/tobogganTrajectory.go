package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var treeCounts = [5]int{0, 0, 0, 0, 0}
	var slopes = [5][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	// read the map input and check each slope for each line
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for n, s := range slopes {
			treeCounts[n] += checkSlope(line, i, s[0], s[1])
		}
	}

	var product = 1
	for _, t := range treeCounts {
		product *= t
	}

	fmt.Printf("Trees encountered on slope 3 right, 1 down: %d\n", treeCounts[1])
	fmt.Printf("Product of encounter nums among all slopes: %d\n", product)
}

// returns 1 if we encounter a tree in this iteration, 0 otherwise
func checkSlope(line string, i int, right int, down int) int {
	if i % down != 0 {
		return 0
	} else if line[right*i/down % len(line)] == '#' {
		return 1
	}
	return 0
}
