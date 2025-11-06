package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var ids = make([]int, 0)

	// read input strings and convert to seat IDs
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		id := seatID(line)
		ids = append(ids, id)
	}

	// sort ID slice and iterate over sorted to find hole
	sort.Ints(ids)
	myID := -1
	for i := 0; i < len(ids); i++ {
		if ids[i] > ids[i-1]+1 {
			myID = ids[i-1] + 1
		}
	}

	fmt.Printf("Max seat ID: %d\n", ids[len(ids)-1])
	fmt.Printf("Own seat ID: %d\n", myID)
}

// Returns the seat ID (row * 8 + column) corresponding to a given partitioning string.
func seatID(partitioning string) int {
	rowPart := partitioning[0:7]
	colPart := partitioning[7:10]
	row := binaryPartitioning(rowPart, 0, 127, 'F', 'B')
	col := binaryPartitioning(colPart, 0, 7, 'L', 'R')
	return row*8 + col
}

// Performs binary partitioning of the interval between `min` and `max` (both included).
// The partitioning is based on the partitioning string `s`, read from left to right,
// on `lowRune` continues with the lower half of the interval on `highRune` with the upper half.
// Returns the unique final number if the interval contains `2^n` numbers, where `n = len(s)`.
func binaryPartitioning(s string, min int, max int, lowRune rune, highRune rune) int {
	for _, r := range s {
		if r == lowRune {
			max = (min + max) / 2
		} else if r == highRune {
			min = (min + max + 1) / 2
		}
	}
	return min
}
