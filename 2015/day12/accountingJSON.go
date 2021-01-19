package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sumTotal, sumNoRed, level, number, levelRunningCount int
	var lastChar rune
	var currentlyIgnoring, readingColor bool
	var color string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for _, r := range scanner.Text() {
		if r == '"' && lastChar == ':' {
			readingColor = true
		} else if r == '"' {
			if color == "red" {
				level = 0
				currentlyIgnoring = true
			}
			readingColor = false
			color = ""
		} else if readingColor {
			color += string(r)
		} else if r >= '0' && r <= '9' {
			number *= 10
			if number < 0 {
				number -= int(r - '0')
			} else {
				number += int(r - '0')
			}
			if lastChar == '-' {
				number *= -1
			}
		} else {
			if number != 0 {
				levelRunningCount += number
				number = 0
			}
			if r == '[' || r == '{' {
				level++
			} else if r == ']' || r == '}' {
				sumTotal += levelRunningCount
				if !currentlyIgnoring {
					sumNoRed += levelRunningCount
				} else if level == 0 {
					currentlyIgnoring = false
				}
				level--
				levelRunningCount = 0
			}
		}
		lastChar = r
	}

	fmt.Println("Sum of all numbers in the JSON:", sumTotal)
	fmt.Println("Sum when ignoring everything red:", sumNoRed)
}
