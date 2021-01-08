package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nineCups, millionCups []int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nineCups = make([]int, len(scanner.Bytes())+1)
	previous, first := int(scanner.Bytes()[0]-'0'), int(scanner.Bytes()[0]-'0')
	for _, cup := range scanner.Bytes()[1:] {
		nineCups[previous] = int(cup - '0')
		previous = int(cup - '0')
	}
	nineCups[previous] = first

	millionCups = make([]int, 1_000_000+1)
	copy(millionCups, nineCups)
	millionCups[previous] = 10
	for i := 10; i <= 1_000_000; i++ {
		millionCups[previous] = i
		previous = i
	}
	millionCups[previous] = first

	// perform the game rounds
	runCrabGame(nineCups, 100, first)
	runCrabGame(millionCups, 10_000_000, first)

	// determine final order (part 1)
	for i := nineCups[1]; i != 1; i = nineCups[i] {
		fmt.Printf("%v", i)
	}
	fmt.Printf("\n")

	// find the two cups immediately after 1 and multiply their IDs (part 2)
	for i := millionCups[1]; ; i = millionCups[i] {
		if i == 1 {
			c1 := millionCups[i]
			c2 := millionCups[c1]
			fmt.Println(c1 * c2)
			break
		}
	}
}

func runCrabGame(cups []int, rounds int, startCup int) {
	var current = startCup
	var pickedUp = make([]int, 3)

	for round := 0; round < rounds; round++ {
		// pick up next 3 cups
		pickedUp[0] = cups[current]
		pickedUp[1] = cups[pickedUp[0]]
		pickedUp[2] = cups[pickedUp[1]]

		// determine destination cup label
		destination := current - 1
		if destination == 0 {
			destination = len(cups) - 1
		}
		for destination == pickedUp[0] || destination == pickedUp[1] || destination == pickedUp[2] {
			destination--
			if destination == 0 {
				destination = len(cups) - 1
			}
		}

		// build the new cup arrangement
		tmp := cups[pickedUp[2]]
		cups[pickedUp[2]] = cups[destination]
		cups[destination] = cups[current]
		cups[current] = tmp
		current = cups[current]
	}
}
