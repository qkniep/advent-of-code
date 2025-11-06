package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var severity int
	var depths, ranges []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var d, r int
		fmt.Sscanf(scanner.Text(), "%d: %d", &d, &r)
		depths = append(depths, d)
		ranges = append(ranges, r)
		if caught(d, r) {
			severity += d * r
		}
	}

	fmt.Printf("Total trip severity: %v\n", severity)
	fmt.Printf("Min delay to not get caught: %v\n", findMinDelay(depths, ranges))
}

// Find the minimum delay for which we are not caught on any layer.
func findMinDelay(depths []int, ranges []int) int {
	for delay := 1; ; delay++ {
		gotCaught := false
		for i, d := range depths {
			if caught(d+delay, ranges[i]) {
				gotCaught = true
				break
			}
		}
		if !gotCaught {
			return delay
		}
	}
}

// Checks whether we are caught if we reach a layer of range `rng` at some time.
func caught(time int, rng int) bool {
	m := rng*2 - 2
	dm := time % m
	return dm == 0
}
