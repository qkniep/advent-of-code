package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pos struct {
	x, y, z int
}

func main() {
	childPosition, furthestDistance := readInput()
	fmt.Printf("Distance: %v\n", calculateDistance(childPosition))
	fmt.Printf("Furthest Distance: %v\n", furthestDistance)
}

func readInput() (p pos, furthest int) {
	dirs := map[string]pos{
		"n":  {1, 0, -1},
		"s":  {-1, 0, 1},
		"ne": {1, -1, 0},
		"nw": {0, 1, -1},
		"se": {0, -1, 1},
		"sw": {-1, 1, 0},
	}

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		direction := ""
		for _, char := range scanner.Text() {
			if char == ',' {
				d := dirs[direction]
				p.x, p.y, p.z = p.x+d.x, p.y+d.y, p.z+d.z
				dist := calculateDistance(p)
				if dist > furthest {
					furthest = dist
				}
				direction = ""
			} else {
				direction += string(char)
			}
		}
		d := dirs[direction]
		p.x, p.y, p.z = p.x+d.x, p.y+d.y, p.z+d.z
		dist := calculateDistance(p)
		if dist > furthest {
			furthest = dist
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func calculateDistance(p pos) int {
	if p.x < 0 {
		p.x = -p.x
	}
	if p.y < 0 {
		p.y = -p.y
	}
	if p.z < 0 {
		p.z = -p.z
	}
	return (p.x + p.y + p.z) / 2
}
