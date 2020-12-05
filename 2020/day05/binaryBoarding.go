package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	max := 0
	ids := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		id := seatID(line)
		ids = append(ids, id)
		if id > max {
			max = id
		}
	}
	sort.Ints(ids)
	previous := ids[0]
	for _, id := range ids[1:] {
		if id > previous + 1 {
			break
		}
		previous = id
	}
	fmt.Printf("Max ID: %d\n", max)
	fmt.Printf("My ID: %d\n", previous+1)
}

func seatID(partitioning string) int {
	rowPart := partitioning[0:7]
	colPart := partitioning[7:10]
	rowMin, rowMax := 0, 127
	for _, r := range rowPart {
		if r == 'F' {
			rowMax = (rowMin + rowMax) / 2
		} else if r == 'B' {
			rowMin = (rowMin + rowMax + 1) / 2
		}
	}
	colMin, colMax := 0, 7
	for _, r := range colPart {
		if r == 'L' {
			colMax = (colMin + colMax) / 2
		} else if r == 'R' {
			colMin = (colMin + colMax + 1) / 2
		}
	}
	return rowMin * 8 + colMin
}
