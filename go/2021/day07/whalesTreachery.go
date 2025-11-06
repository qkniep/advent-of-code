package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var positions []int

	// read horizontal positions
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for _, p := range strings.Split(scanner.Text(), ",") {
		pos, _ := strconv.Atoi(p)
		positions = append(positions, pos)
	}

	// find bounds for positions
	var min = 999_999_999
	var max = 0
	for _, pos := range positions {
		if pos < min {
			min = pos
		} else if pos > max {
			max = pos
		}
	}

	// find minimum fuel cost needed to align the submarines
	var minCost1, minCost2 = 999_999_999, 999_999_999
	for alignTo := min; alignTo < max; alignTo++ {
		var sum1, sum2 int
		for _, pos := range positions {
			if pos > alignTo {
				sum1 += pos - alignTo
				sum2 += (pos - alignTo) * (pos - alignTo + 1) / 2
			} else {
				sum1 += alignTo - pos
				sum2 += (alignTo - pos) * (alignTo - pos + 1) / 2
			}
		}
		if sum1 < minCost1 {
			minCost1 = sum1
		} else if sum2 < minCost2 {
			minCost2 = sum2
		}
	}

	fmt.Println("Fuel to align with basic cost function:", minCost1)
	fmt.Println("Fuel to align with triangular cost function:", minCost2)
}
