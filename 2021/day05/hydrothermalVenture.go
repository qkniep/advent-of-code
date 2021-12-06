package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type line struct {
	from, to, dir [2]int
}

type diagram struct {
	data                   map[[2]int]int
	minX, minY, maxX, maxY int
}

func main() {
	var lines []line

	// read lines
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// read from and to positions
		var l line
		tokens := strings.Fields(scanner.Text())
		fmt.Sscanf(tokens[0], "%d,%d", &l.from[0], &l.from[1])
		fmt.Sscanf(tokens[2], "%d,%d", &l.to[0], &l.to[1])

		// determine pipe direction
		if l.from[0] > l.to[0] {
			l.dir[0] = -1
		} else if l.to[0] > l.from[0] {
			l.dir[0] = 1
		}
		if l.from[1] > l.to[1] {
			l.dir[1] = -1
		} else if l.to[1] > l.from[1] {
			l.dir[1] = 1
		}

		lines = append(lines, l)
	}

	// build diagrams and count overlaps
	diagram1 := buildDiagram(lines, true)
	count1 := countOverlapping(diagram1)
	diagram2 := buildDiagram(lines, false)
	count2 := countOverlapping(diagram2)

	fmt.Println("Number of overlapping points without diagonal pipes:", count1)
	fmt.Println("Number of overlapping points with diagonal pipes", count2)
}

// Builds a diagram of the pipes with the number of overlapping pipes in each grid position.
func buildDiagram(lines []line, skipDiagonal bool) (dia diagram) {
	dia.data = make(map[[2]int]int, 0)

	for _, l := range lines {
		// filter out invalid (non-straight) lines
		if skipDiagonal && l.dir[0] != 0 && l.dir[1] != 0 {
			continue
		}

		// mark fields
		for l.from[0] != l.to[0]+l.dir[0] || l.from[1] != l.to[1]+l.dir[1] {
			// adapt x,y limits
			if l.from[0] < dia.minX {
				dia.minX = l.from[0]
			} else if l.from[0] > dia.maxX {
				dia.maxX = l.from[0]
			}
			if l.from[1] < dia.minY {
				dia.minY = l.from[1]
			} else if l.from[1] > dia.maxY {
				dia.maxY = l.from[1]
			}

			dia.data[l.from]++
			l.from[0] += l.dir[0]
			l.from[1] += l.dir[1]
		}
	}
	return
}

// Counts the number of grid positions in which at least 2 pipes overlap.
func countOverlapping(dia diagram) (count int) {
	for y := dia.minY; y <= dia.maxY; y++ {
		for x := dia.minX; x <= dia.maxX; x++ {
			val, exists := dia.data[[2]int{x, y}]
			if exists && val > 1 {
				count++
			}
		}
	}
	return
}
