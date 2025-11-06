package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	passphrases := readInput()
	fmt.Printf("Part 1: %v\n", countValid(passphrases, false))
	fmt.Printf("Part 2: %v\n", countValid(passphrases, true))
}

func readInput() (pp [][]string) {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passphrase := make([]string, 0)
		fmt.Println(scanner.Text())
		for _, s := range strings.Fields(scanner.Text()) {
			passphrase = append(passphrase, s)
		}
		pp = append(pp, passphrase)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func countValid(pp [][]string, anagrams bool) (num int) {
	for _, passphrase := range pp {
		invalid := false
		for i, word1 := range passphrase {
			for _, word2 := range passphrase[i+1:] {
				if anagrams && areAnagrams(word1, word2) {
					invalid = true
					break
				} else if word1 == word2 {
					invalid = true
					break
				}
			}
			if invalid {
				break
			}
		}
		if !invalid {
			num++
		}
	}
	return
}

// RuneSlice is a type alias for a slice of runes
type RuneSlice []rune

func (r RuneSlice) Len() int           { return len(r) }
func (r RuneSlice) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r RuneSlice) Less(i, j int) bool { return r[i] < r[j] }

func areAnagrams(a string, b string) (is bool) {
	var r1 RuneSlice = []rune(a)
	var r2 RuneSlice = []rune(b)

	sort.Sort(r1)
	sort.Sort(r2)

	return string(r1) == string(r2)
}
