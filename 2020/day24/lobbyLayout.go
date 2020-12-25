package main

import (
	"bufio"
	"fmt"
	"os"
)

type point2d struct {
	x int
	y int
}

func main() {
	var flipped = make(map[point2d]bool, 0)
	var mode = 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		path := scanner.Text()
		pos := point2d{x: 0, y: 0}
		for _, r := range path {
			if r == 'n' {
				mode = 1
			} else if r == 's' {
				mode = 2
			} else {
				if r == 'e' {
					if mode == 1 {
						pos.y++
					} else if mode == 2 {
						pos.x++
						pos.y--
					} else {
						pos.x++
					}
				} else if r == 'w' {
					if mode == 1 {
						pos.x--
						pos.y++
					} else if mode == 2 {
						pos.y--
					} else {
						pos.x--
					}
				}
				mode = 0
			}
		}
		flipped[pos] = !flipped[pos]
	}

	count := 0
	for _, isFlipped := range flipped {
		if isFlipped {
			count++
		}
	}

	fmt.Printf("Tiles switched to black: %v\n", count)
}
