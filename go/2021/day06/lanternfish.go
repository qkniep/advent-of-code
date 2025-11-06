package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// keeps track of number of fish for each timer value
	var fish = make([]int, 9)

	// read initial timers
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for _, t := range strings.Split(scanner.Text(), ",") {
		timer, _ := strconv.Atoi(t)
		fish[timer]++
	}

	fmt.Println("Number of lanternfish after 80 days:", simulate(fish, 80))
	fmt.Println("Number of lanternfish after 256 days:", simulate(fish, 256))
}

// Simulates the reproduction of lanternfish given an initial state of timers.
// Returns the total number of lanternfish after a given number of days.
func simulate(initialState []int, days int) int {
	var fish = make([]int, len(initialState))
	copy(fish, initialState)
	for d := 0; d < days; d++ {
		newFish := fish[0]
		for i := 0; i < len(fish)-1; i++ {
			fish[i] = fish[i+1]
		}
		fish[6] += newFish
		fish[8] = newFish
	}

	var sum int
	for i := 0; i < len(fish); i++ {
		sum += fish[i]
	}
	return sum
}
