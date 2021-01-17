package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var currentFloor int
	var firstBasementPos = 0

	// read the instructions, keeping track of current floor
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for i, r := range scanner.Text() {
		if r == '(' {
			currentFloor++
		} else if r == ')' {
			currentFloor--
		}
		if currentFloor == -1 && firstBasementPos == 0 {
			firstBasementPos = i+1
		}
	}

	fmt.Printf("Final floor: %v\n", currentFloor)
	fmt.Printf("First basement instruction: %v\n", firstBasementPos)
}
