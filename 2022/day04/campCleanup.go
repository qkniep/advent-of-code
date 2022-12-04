package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type assignment struct {
	min1 int
	max1 int
	min2 int
	max2 int
}

func main() {
	var assignments []assignment
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ranges := strings.Split(scanner.Text(), ",")
		range1, range2 := strings.Split(ranges[0], "-"), strings.Split(ranges[1], "-")
		min1, _ := strconv.Atoi(range1[0])
		max1, _ := strconv.Atoi(range1[1])
		min2, _ := strconv.Atoi(range2[0])
		max2, _ := strconv.Atoi(range2[1])
		assignments = append(assignments, assignment{min1, max1, min2, max2})
	}

	var countCompleteOverlap = 0
	var countAnyOverlap = 0
	for _, a := range assignments {
		if (a.min1 >= a.min2 && a.max1 <= a.max2) || (a.min2 >= a.min1 && a.max2 <= a.max1) {
			countCompleteOverlap++
			countAnyOverlap++
		} else if (a.min1 >= a.min2 && a.min1 <= a.max2) || (a.max1 >= a.min2 && a.max1 <= a.max2) {
			countAnyOverlap++
		}
	}

	fmt.Println("Number of assignments completely overlapping:", countCompleteOverlap)
	fmt.Println("Number of assignments with any overlap:", countAnyOverlap)
}
