package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var digits = make([]int, len(scanner.Text()))
	for i, r := range scanner.Text() {
		digits[i] = int(r - '0')
	}

	iteration40 := lookAndSay(digits, 40)
	iteration50 := lookAndSay(iteration40, 10)

	fmt.Println("Length after 40 iterations:", len(iteration40))
	fmt.Println("Length after 50 iterations:", len(iteration50))
}

func lookAndSay(start []int, iterations int) []int {
	current := start
	for i := 0; i < iterations; i++ {
		var newCurrent []int
		var digit, count int
		for i, d := range current {
			if d == digit {
				count++
			} else {
				if i > 0 {
					newCurrent = append(newCurrent, count, digit)
				}
				digit = d
				count = 1
			}
		}
		newCurrent = append(newCurrent, count, digit)
		current = newCurrent
	}
	return current
}
