package main

import (
	"bufio"
	"fmt"
	"os"
)

type action struct {
	cmd rune
	val int
}

func main() {
	var actions = make([]action, 0)

	// read the instructions as cmd, val pairs
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd, val := 'X', 0
		fmt.Sscanf(scanner.Text(), "%c%d", &cmd, &val)
		actions = append(actions, action{cmd, val})
	}

	x, y := performActions(actions, false)
	noWaypoint := abs(x) + abs(y)
	x, y = performActions(actions, true)
	withWaypoint := abs(x) + abs(y)

	fmt.Printf("No waypoint: %d\n", noWaypoint)
	fmt.Printf("With waypoint: %d\n", withWaypoint)
}

// Performs all `actions`, the meaning of these depend on whether `useWaypoint` is set.
// Returns the final x,y coordinates of the ship.
func performActions(actions []action, useWaypoint bool) (int, int) {
	var shipX, shipY, waypX, waypY = 0, 0, 1, 0
	var cardX, cardY = &shipX, &shipY

	if useWaypoint {
		waypX, waypY = 10, 1
		cardX, cardY = &waypX, &waypY
	}

	for _, act := range actions {
		switch act.cmd {
		case 'L':
			for s := 0; s < act.val / 90; s++ {
				oldX := waypX
				waypX = -waypY
				waypY = oldX
			}
		case 'R':
			for s := 0; s < act.val / 90; s++ {
				oldY := waypY
				waypY = -waypX
				waypX = oldY
			}
		case 'F':
			shipX += act.val * waypX
			shipY += act.val * waypY
		case 'E':
			*cardX += act.val
		case 'S':
			*cardY -= act.val
		case 'W':
			*cardX -= act.val
		case 'N':
			*cardY += act.val
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
