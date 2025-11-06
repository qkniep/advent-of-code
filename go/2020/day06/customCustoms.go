package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var answered = make(map[rune]int, 26)
	var countAny, countEvery = 0, 0
	var people = 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			for _, n := range answered {
				if n == people {
					countEvery++
				}
			}
			answered = make(map[rune]int, 26)
			people = 0
			continue
		}
		for _, r := range line {
			if answered[r] == 0 {
				countAny++
			}
			answered[r]++
		}
		people++
	}
	for _, n := range answered {
		if n == people {
			countEvery++
		}
	}

	fmt.Printf("Anyone answered yes: %d\n", countAny)
	fmt.Printf("Everyone answered yes: %d\n", countEvery)
}
