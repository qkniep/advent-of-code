package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Must have either x1==x2 or y1==y2 to be a valid line.
type line struct {
	x1, y1, x2, y2 int
}

func (l line) parallel(o line) bool {
	return (l.x1 == l.x2) == (o.x1 == o.x2)
}

func (l line) length() int {
	if l.x1 == l.x2 {
		return abs(l.y1 - l.y2)
	}
	return abs(l.x1 - l.x2)
}

func main() {
	var wire1, wire2 []line
	var scanner = bufio.NewScanner(os.Stdin)

	scanner.Scan()
	wire1 = constructWire(scanner.Text())
	scanner.Scan()
	wire2 = constructWire(scanner.Text())

	fmt.Println("Distance to origin of closest crossing:", findMinCrossingDistance(wire1, wire2))
	fmt.Println("Number of steps until first crossing:", findMinCrossingSteps(wire1, wire2))
}

func constructWire(s string) []line {
	var segments []line
	var x, y, nx, ny = 0, 0, 0, 0

	for _, step := range strings.Split(s, ",") {
		stepSize, _ := strconv.Atoi(step[1:])
		switch rune(step[0]) {
		case 'R':
			nx += stepSize
		case 'L':
			nx -= stepSize
		case 'U':
			ny += stepSize
		case 'D':
			ny -= stepSize
		}
		segments = append(segments, line{x, y, nx, ny})
		x, y = nx, ny
	}

	return segments
}

func findMinCrossingSteps(wire1, wire2 []line) int {
	var minSteps = 99999
	var steps, oldSteps, newSteps int

	for _, segment1 := range wire1 {
		oldSteps = steps
		for _, segment2 := range wire2 {
			doIntersect, x, y := intersectLines(segment1, segment2)
			if segment1.x1 == segment1.x2 {
				newSteps = steps + abs(y - segment1.y1) + abs(x - segment2.x1)
			} else {
				newSteps = steps + abs(y - segment2.y1) + abs(x - segment1.x1)
			}
			if doIntersect && (x != 0 || y != 0) && newSteps < minSteps {
				minSteps = newSteps
			}
			steps += segment2.length()
		}
		steps = oldSteps + segment1.length()
	}

	return minSteps
}

func findMinCrossingDistance(wire1, wire2 []line) int {
	var minDistance = 99999

	for _, segment1 := range wire1 {
		for _, segment2 := range wire2 {
			doIntersect, x, y := intersectLines(segment1, segment2)
			if doIntersect && (x != 0 || y != 0) && abs(x) + abs(y) < minDistance {
				minDistance = abs(x) + abs(y)
			}
		}
	}

	return minDistance
}

// Checks two lines for intersection.
// Returns a boolean indication whether they intersect and the position of intersection.
func intersectLines(line1, line2 line) (bool, int, int) {
	if line1.parallel(line2) {
		return false, 0, 0
	}
	if line1.x1 != line1.x2 { // let line1 always be the vertical one
		line1, line2 = line2, line1
	}
	if inRange(line2.y1, line1.y1, line1.y2) && inRange(line1.x1, line2.x1, line2.x2) {
		return true, line1.x1, line2.y1
	}
	return false, 0, 0
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Checks whether n is in the given range. where we don't know the order of the bounds.
func inRange(n, bound1, bound2 int) bool {
	return (n >= bound1 && n <= bound2) || (n >= bound2 && n <= bound1)
}
