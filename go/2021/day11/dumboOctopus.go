package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	x, y int
}

func main() {
	var octopuses [][]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var octopusRow []int
		for _, octopus := range scanner.Text() {
			octopusRow = append(octopusRow, int(octopus-'0'))
		}
		octopuses = append(octopuses, octopusRow)
	}

	flashesIn100Round, firstSimultaneous := simulate(octopuses)

	fmt.Println("Total flashes in the first 100 steps:", flashesIn100Round)
	fmt.Println("Number of steps until all octopuses flash simultaneously:", firstSimultaneous)
}

func simulate(octopuses [][]int) (totalFlashes int, firstSimultaneous int) {
	for r := 0; r < 100 || firstSimultaneous == 0; r++ {
		var toFlash = []pos{}
		var flashed = make(map[pos]bool)
		for y, row := range octopuses {
			for x := range row {
				octopuses[y][x] += 1
				if octopuses[y][x] > 9 {
					toFlash = append(toFlash, pos{x, y})
				}
			}
		}
		for len(toFlash) > 0 {
			nextFlash := toFlash[0]
			toFlash = toFlash[1:]
			if flashed[nextFlash] {
				continue
			}
			flashed[nextFlash] = true
			if r < 100 {
				totalFlashes++
			}
			for y := nextFlash.y - 1; y <= nextFlash.y+1; y++ {
				for x := nextFlash.x - 1; x <= nextFlash.x+1; x++ {
					if x < 0 || y < 0 || x >= len(octopuses[0]) || y >= len(octopuses) {
						continue
					}
					octopuses[y][x] += 1
					if octopuses[y][x] > 9 {
						toFlash = append(toFlash, pos{x, y})
					}
				}
			}
		}
		for pos := range flashed {
			octopuses[pos.y][pos.x] = 0
		}
		if firstSimultaneous == 0 && len(flashed) == len(octopuses)*len(octopuses[0]) {
			firstSimultaneous = r + 1
		}
	}
	return
}
