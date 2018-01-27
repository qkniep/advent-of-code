package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	jumps := readInput()
	jumpsCopy := append([]int(nil), jumps...)
	fmt.Printf("Part 1: %v\n", escapeMaze(jumps, false))
	fmt.Printf("Part 2: %v\n", escapeMaze(jumpsCopy, true))
}

func readInput() (j []int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jump, err := strconv.Atoi(scanner.Text())
		if err == nil {
			j = append(j, jump)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func escapeMaze(j []int, part2 bool) (steps int) {
	currentPos := 0
	for currentPos >= 0 && currentPos < len(j) {
		newPos := currentPos + j[currentPos]
		if part2 && j[currentPos] >= 3 {
			j[currentPos] -= 1
		} else {
			j[currentPos] += 1
		}
		currentPos = newPos
		steps += 1
	}
	return
}
