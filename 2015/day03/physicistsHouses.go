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
	var visited = map[pos]bool{{0, 0}: true}
	var visitedRobo = map[pos]bool{{0, 0}: true}
	var currentPos, santaPos, roboPos = pos{0, 0}, pos{0, 0}, pos{0, 0}

	// perform the movements keeping track of the possible positions of:
	//   - santa moving alone (on every move)
	//   - santa moving only on even moves
	//   - robo santa moving only on odd moves
	// also mark houses as visited upon hitting their location
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for i, r := range scanner.Text() {
		applyChange(&currentPos, r)
		if i % 2 == 1 {
			applyChange(&roboPos, r)
			visitedRobo[roboPos] = true
		} else {
			applyChange(&santaPos, r)
			visitedRobo[santaPos] = true
		}
		visited[currentPos] = true
	}

	fmt.Printf("Number of houses (Santa alone): %v\n", len(visited))
	fmt.Printf("Number of houses (w/ Robo Santa): %v\n", len(visitedRobo))
}

func applyChange(p *pos, dir rune) {
	if dir == '<' {
		p.x--
	} else if dir == '>' {
		p.x++
	} else if dir == '^' {
		p.y++
	} else if dir == 'v' {
		p.y--
	}
}
