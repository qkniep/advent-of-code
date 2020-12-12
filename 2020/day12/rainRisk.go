package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	east = iota
	south
	west
	north
)

type instruction struct {
	cmd rune
	val int
}

func main() {
	var instructions = make([]instruction, 0)

	// read the seat layout line by line (row by row)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd, val := 'X', 0
		fmt.Sscanf(scanner.Text(), "%c%d", &cmd, &val)
		instructions = append(instructions, instruction{cmd, val})
	}

	x, y := performInstructions(instructions)
	noWaypoint := abs(x) + abs(y)

	x, y = performInstructions2(instructions)
	withWaypoint := abs(x) + abs(y)

	fmt.Printf("Manhattan distance: %d\n", noWaypoint)
	fmt.Printf("Manhattan distance: %d\n", withWaypoint)
}

func performInstructions(inst []instruction) (int, int) {
	var direction, xPos, yPos = east, 0, 0
	var dx = [4]int{1, 0, -1, 0}
	var dy = [4]int{0, -1, 0, 1}

	for _, i := range inst {
		switch i.cmd {
		case 'L':
			direction = mod(direction - i.val / 90, 4)
		case 'R':
			direction = mod(direction + i.val / 90, 4)
		case 'F':
			xPos += i.val * dx[direction]
			yPos += i.val * dy[direction]
		case 'E':
			xPos += i.val
		case 'S':
			yPos -= i.val
		case 'W':
			xPos -= i.val
		case 'N':
			yPos += i.val
		}
	}

	return xPos, yPos
}

func performInstructions2(inst []instruction) (int, int) {
	var shipX, shipY, waypX, waypY = 0, 0, 10, 1

	for _, i := range inst {
		switch i.cmd {
		case 'L':
			for s := 0; s < i.val / 90; s++ {
				oldX := waypX
				waypX = -waypY
				waypY = oldX
			}
		case 'R':
			for s := 0; s < i.val / 90; s++ {
				oldY := waypY
				waypY = -waypX
				waypX = oldY
			}
		case 'F':
			shipX += i.val * waypX
			shipY += i.val * waypY
		case 'E':
			waypX += i.val
		case 'S':
			waypY -= i.val
		case 'W':
			waypX -= i.val
		case 'N':
			waypY += i.val
		}
	}

	return shipX, shipY
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func mod(a, b int) int {
    return (a % b + b) % b
}
