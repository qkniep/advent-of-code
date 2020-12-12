package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var seatLayout = make([][]byte, 0)

	// read the seat layout line by line (row by row)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := []byte(scanner.Text())
		seatLayout = append(seatLayout, row)
	}
	seatLayout2 := make([][]byte, len(seatLayout))
	deepCopy(seatLayout2, seatLayout)

	occupiedNxt := simulateSeating(seatLayout, 4, false)
	occupiedVis := simulateSeating(seatLayout2, 5, true)

	fmt.Printf("Occupied seats (next to): %d\n", occupiedNxt)
	fmt.Printf("Occupied seats (visible): %d\n", occupiedVis)
}

// Simulates the seating process until equilibrium is reached, i.e. no more changes happen.
// `occupiedThreshold`: Number of seats that need to be occupied for a seat to become empty.
// `byVisibility`: If true, check for visible occupied seats, otherwise check surrounding seats.
// Returns the final number of occupied seats.
func simulateSeating(seats [][]byte, occupiedThreshold int, byVisibility bool) int {
	var occupied, occupiedNextTo, dirty = 0, 0, true
	var newSeats = make([][]byte, len(seats))
	deepCopy(newSeats, seats)

	for dirty {
		dirty = false
		for y := 0; y < len(seats); y++ {
			for x := 0; x < len(seats[0]); x++ {
				if byVisibility {
					occupiedNextTo = countOccupiedVisible(seats, x, y)
				} else {
					occupiedNextTo = countOccupiedNextTo(seats, x, y)
				}

				if occupiedNextTo == 0 && seats[y][x] == byte('L') {
					newSeats[y][x] = byte('#')
					occupied++
					dirty = true
				} else if occupiedNextTo >= occupiedThreshold && seats[y][x] == '#' {
					newSeats[y][x] = byte('L')
					occupied--
					dirty = true
				}
			}
		}
		deepCopy(seats, newSeats)
	}

	return occupied
}

// Counts the number of occupied seats directly next to (x,y), on any of the 8 possible spots.
func countOccupiedNextTo(seats [][]byte, x int, y int) int {
	var occupied = 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if (dx == 0 && dy == 0) || y+dy < 0 || x+dx < 0 ||
					y+dy >= len(seats) || x+dx >= len(seats[0]) {
				continue
			} else if seats[y+dy][x+dx] == '#' {
				occupied++
			}
		}
	}
	return occupied
}

// Counts the number of occupied seats visible from (x,y) by following lines in all 8 directions.
func countOccupiedVisible(seats [][]byte, x int, y int) int {
	var occupied = 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if (dx == 0 && dy == 0) {
				continue
			}
			for s := 1;; s++ {
				if y+dy*s < 0 || x+dx*s < 0 || y+dy*s >= len(seats) || x+dx*s >= len(seats[0]) {
					break
				} else if seats[y+dy*s][x+dx*s] == '#' {
					occupied++
					break
				} else if seats[y+dy*s][x+dx*s] == 'L' {
					break
				}
			}
		}
	}
	return occupied
}

// Perform a deep copy of the 2D slice in src into the 2D slice dst.
// Assumes that dst has been initialized to have the same length as src, inner slices can be empty.
func deepCopy(dst [][]byte, src [][]byte) {
	for i := range src {
		dst[i] = make([]byte, len(src[i]))
		copy(dst[i], src[i])
	}
}
