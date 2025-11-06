package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var startingNums = make([]int, 0)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	for _, numStr := range strings.Split(strings.TrimSpace(line), ",") {
		num, _ := strconv.Atoi(numStr)
		startingNums = append(startingNums, num)
	}

	fmt.Printf("2020th spoken number: %v\n", findNthSpoken(startingNums, 2020))
	fmt.Printf("30000000th spoken number: %v\n", findNthSpoken(startingNums, 30000000))
}

func findNthSpoken(startingNums []int, n int) int {
	var turnSpoken = make(map[int]int, 0)
	var lastSpoken, turnsApart int

	for turn := 1; turn <= n; turn++ {
		if turn <= len(startingNums) {
			lastSpoken = startingNums[turn-1]
		} else {
			lastSpoken = turnsApart
		}
		if turnSpoken[lastSpoken] == 0 {
			turnsApart = 0
		} else {
			turnsApart = turn - turnSpoken[lastSpoken]
		}
		turnSpoken[lastSpoken] = turn
	}

	return lastSpoken
}
