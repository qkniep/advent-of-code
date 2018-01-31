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
	childPosition := readInput()
	fmt.Printf("Distance: %v\n", calculateDistance(childPosition))
	//fmt.Printf("Garbaged Characters: %v\n", garbaged)
}

func readInput() (p pos) {
	dirs := map[string]pos{
		"n":  pos{1, 0, -1},
		"s":  pos{-1, 0, 1},
		"ne": pos{1, -1, 0},
		"nw": pos{0, 1, -1},
		"se": pos{0, -1, 1},
		"sw": pos{-1, 1, 0},
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
				direction = ""
			} else {
				direction += string(char)
			}
		}
		d := dirs[direction]
		p.x, p.y, p.z = p.x+d.x, p.y+d.y, p.z+d.z
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
