package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	score, garbaged := readInput()
	fmt.Printf("Score: %v\n", score)
	fmt.Printf("Garbaged Characters: %v\n", garbaged)
}

func readInput() (score int, garbagedChars int) {
	depth := 0
	skipChar := false
	garbage := false

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		for _, char := range scanner.Text() {
			if skipChar {
				skipChar = false
			} else if char == '!' {
				skipChar = true
			} else if garbage {
				if char == '>' {
					garbage = false
				} else {
					garbagedChars += 1
				}
			} else {
				if char == '<' {
					garbage = true
				} else if char == '{' {
					depth += 1
				} else if char == '}' {
					score += depth
					depth -= 1
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}
