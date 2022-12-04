package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var charCorruptionScores map[byte]int = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var charAddScores map[byte]int = map[byte]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func main() {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	corruptionScore, completionScores := calculateScores(lines)
	sort.Ints(completionScores)
	middleScore := completionScores[len(completionScores)/2]

	fmt.Println("Sum of corruption scores of corrupted lines:", corruptionScore)
	fmt.Println("Middle completion score among non-corrupted lines:", middleScore)
}

func calculateScores(subsystem []string) (corruptionScore int, completionScores []int) {
	for _, line := range subsystem {
		var charStack []byte
		var corrupted = false
		var score = 0
		for _, char := range line {
			if char == '(' || char == '[' || char == '{' || char == '<' {
				charStack = append(charStack, byte(char))
			} else if char == ')' || char == ']' || char == '}' || char == '>' {
				if len(charStack) == 0 {
					corruptionScore += charCorruptionScores[byte(char)]
					corrupted = true
					break
				} else {
					lastChar := charStack[len(charStack)-1]
					if lastChar == '(' && char == ')' ||
						lastChar == '[' && char == ']' ||
						lastChar == '{' && char == '}' ||
						lastChar == '<' && char == '>' {
						charStack = charStack[:len(charStack)-1]
					} else {
						corruptionScore += charCorruptionScores[byte(char)]
						corrupted = true
						break
					}
				}
			}
		}

		// only try to complete the line if it wasn't detected as being corrupted
		if !corrupted {
			for i := len(charStack) - 1; i >= 0; i-- {
				score *= 5
				score += charAddScores[charStack[i]]
			}
			completionScores = append(completionScores, score)
		}
	}
	return
}
