package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	up = iota
	down
	left
	right
)

func main() {
	var diagram []string
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		diagram = append(diagram, scanner.Text())
	}

	start := findStart(diagram)
	letterSeq, steps := followPath(diagram, start)

	fmt.Printf("Sequence of letters: %v\n", letterSeq)
	fmt.Printf("Number of steps: %v\n", steps)
}

func findStart(diagram []string) int {
	return strings.Index(diagram[0], "|")
}

func followPath(diagram []string, start int) (string, int) {
	var lettersFound string
	var dir = down
	var steps = 0
	for x, y := start, 0; ; x, y = nextPos(x, y, dir) {
		if diagram[y][x] == '|' || diagram[y][x] == '-' {
			// do nothing
		} else if diagram[y][x] == '+' {
			if dir != up && y < len(diagram)-1 && diagram[y+1][x] != '-' && diagram[y+1][x] != ' ' {
				dir = down
			} else if dir != down && y > 0 && diagram[y-1][x] != '-' && diagram[y-1][x] != ' ' {
				dir = up
			} else if dir != right && x > 0 && diagram[y][x-1] != '|' && diagram[y][x-1] != ' ' {
				dir = left
			} else if dir != left && x < len(diagram[y])-1 && diagram[y][x+1] != '|' && diagram[y][x+1] != ' ' {
				dir = right
			}
		} else if diagram[y][x] >= 'A' && diagram[y][x] <= 'Z' {
			lettersFound += string(rune(diagram[y][x]))
		} else {
			break
		}
		steps++
	}
	return lettersFound, steps
}

func nextPos(x, y, direction int) (int, int) {
	switch direction {
	case down:
		return x, y+1
	case up:
		return x, y-1
	case left:
		return x-1, y
	case right:
		return x+1, y
	}
	panic("invalid direction")
}
