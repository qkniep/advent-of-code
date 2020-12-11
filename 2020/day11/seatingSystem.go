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

	occupied := simulateSeating(seatLayout)
	occupied2 := simulateSeating2(seatLayout2)

	fmt.Printf("Occupied seats: %d\n", occupied)
	fmt.Printf("Occupied visible: %d\n", occupied2)
}

func simulateSeating(seats [][]byte) int {
	var occupied, dirty = 0, true
	var newSeats = make([][]byte, len(seats))
	deepCopy(newSeats, seats)

	for dirty {
		dirty = false
		for y := 0; y < len(seats); y++ {
			for x := 0; x < len(seats[0]); x++ {
				occupiedNextTo := countOccupiedNextTo(seats, x, y)
				if occupiedNextTo == 0 && seats[y][x] == byte('L') {
					newSeats[y][x] = byte('#')
					occupied++
					dirty = true
				} else if occupiedNextTo >= 4 && seats[y][x] == '#' {
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

func simulateSeating2(seats [][]byte) int {
	var occupied, dirty = 0, true
	var newSeats = make([][]byte, len(seats))
	deepCopy(newSeats, seats)

	for dirty {
		dirty = false
		for y := 0; y < len(seats); y++ {
			for x := 0; x < len(seats[0]); x++ {
				occupiedNextTo := countOccupiedVisible(seats, x, y)
				if occupiedNextTo == 0 && seats[y][x] == byte('L') {
					newSeats[y][x] = byte('#')
					occupied++
					dirty = true
				} else if occupiedNextTo >= 5 && seats[y][x] == '#' {
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

func deepCopy(dst [][]byte, src [][]byte) {
	for i := range src {
		dst[i] = make([]byte, len(src[i]))
		copy(dst[i], src[i])
	}
}
