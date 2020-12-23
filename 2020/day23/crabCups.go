package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var cups []int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cups = make([]int, len(scanner.Text()))
	for i, cup := range scanner.Bytes() {
		cups[i] = int(cup - '0')
	}

	// perform 100 game rounds
	for round := 0; round < 100; round++ {
		newCups := make([]int, len(cups))
		pickedUp := cups[1:4]
		// determine destination cup label
		destinationCup := cups[0] - 1
		if destinationCup == 0 {
			destinationCup = 9
		}
		for destinationCup == pickedUp[0] || destinationCup == pickedUp[1] || destinationCup == pickedUp[2] {
			destinationCup--
			if destinationCup == 0 {
				destinationCup = 9
			}
		}
		// build the new cup arrangement
		newCups[0] = cups[4]
		for i, added := 1, 0; i < len(cups); i++ {
			modI := (4+i)%len(cups)
			if modI >= 1 && modI <= 3 {
				continue
			} else if cups[(3+i)%len(cups)] == destinationCup {
				copy(newCups[1+added:4+added], cups[1:4])
				added += 3
			}
			newCups[1+added] = cups[modI]
			added++
		}
		fmt.Printf("%v\n", newCups)
		copy(cups, newCups)
	}

	// determine final order
	var onePos int
	for pos, cup := range cups {
		if cup == 1 {
			onePos = pos
		}
	}
	for i := 1; i < len(cups); i++ {
		fmt.Printf("%v", cups[(onePos+i)%len(cups)])
	}
	fmt.Printf("\n")
}
