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
	var turnSpoken = make(map[int]int, 0)
	var lastSpoken, turnsApart int

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	for _, numStr := range strings.Split(strings.TrimSpace(line), ",") {
		num, _ := strconv.Atoi(numStr)
		startingNums = append(startingNums, num)
	}

	for turn := 1; turn <= 30000000; turn++ {
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

	fmt.Printf("2020th spoken number: %v\n", lastSpoken)
	//fmt.Printf("Version 2 decoder chip: %v\n", sumOfValues(memory2))
}
