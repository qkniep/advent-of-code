package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const vowels = "aeiou"

var naughtyCombinations = map[rune]rune{'a': 'b', 'c': 'd', 'p': 'q', 'x': 'y'}

func main() {
	var niceStrings, betterNiceStrings int

	scanner := bufio.NewScanner(os.Stdin)

	// go through all the strings on the list, checking whether they are nice or naughty
	for scanner.Scan() {
		var secondLastLetter, lastLetter rune
		var numVowels int
		var doubleLetter, naughtyCombination, pair, repetition bool
		var pairs = make(map[string]int, 0)

		for i, r := range scanner.Text() {
			// check old rules
			if strings.ContainsRune(vowels, r) {
				numVowels++
			}
			if lastLetter == r {
				doubleLetter = true
			} else if naughtyCombinations[lastLetter] == r {
				naughtyCombination = true
			}

			// check new rules
			if secondLastLetter == r {
				repetition = true
			}
			pairStr := string(lastLetter) + string(r)
			if pairs[pairStr] == 0 {
				pairs[pairStr] = i
			} else if pairs[pairStr] < i-1 {
				pair = true
			}

			secondLastLetter = lastLetter
			lastLetter = r
		}

		// count strings towards total if rules apply
		if numVowels >= 3 && doubleLetter && !naughtyCombination {
			niceStrings++
		}
		if pair && repetition {
			betterNiceStrings++
		}
	}

	fmt.Println("Number of nice strings:", niceStrings)
	fmt.Println("New better nice strings:", betterNiceStrings)
}
