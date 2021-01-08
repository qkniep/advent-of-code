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
	var count = 0

	// determine initial layout
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pos := pathStrToPos(scanner.Text())
		flipped[pos] = !flipped[pos]
	}

	for _, isFlipped := range flipped {
		if isFlipped {
			count++
		}
	}
	fmt.Printf("Tiles initially black: %v\n", count)

	for day := 0; day < 100; day++ {
		surroundingFlippedTiles := make(map[point2d]int, 0)
		// count surrounding
		for pos := range flipped { // ensure that black tiles with 0 neighbors are flipped
			surroundingFlippedTiles[pos] = 0
		}
		for pos, isFlipped := range flipped {
			if isFlipped {
				for _, neighbor := range surroundingTiles(pos) {
					surroundingFlippedTiles[neighbor]++
				}
			}
		}
		// perform flips
		for pos, numNeighbors := range surroundingFlippedTiles {
			if flipped[pos] && (numNeighbors == 0 || numNeighbors > 2) {
				flipped[pos] = false
			} else if !flipped[pos] && numNeighbors == 2 {
				flipped[pos] = true
			}
		}
	}

	count = 0
	for _, isFlipped := range flipped {
		if isFlipped {
			count++
		}
	}
	fmt.Printf("Tiles black after 100 days: %v\n", count)
}

func pathStrToPos(path string) point2d {
	var mode = 0
	var pos = point2d{x: 0, y: 0}

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

	return pos
}

func surroundingTiles(tile point2d) []point2d {
	return []point2d{
		{tile.x + 1, tile.y},
		{tile.x - 1, tile.y},
		{tile.x, tile.y + 1},
		{tile.x, tile.y - 1},
		{tile.x + 1, tile.y - 1},
		{tile.x - 1, tile.y + 1},
	}
}
