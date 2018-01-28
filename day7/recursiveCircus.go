package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reachable, weights, children := readInput()
	fmt.Printf("Root Program: %v\n", findRoot(reachable))
	fmt.Printf("Correct Weight: %v\n", fixWrongWeight(weights, children))
}

func readInput() (reachable map[string]bool, weights map[string]int, children map[string][]string) {
	reachable = make(map[string]bool)
	weights = make(map[string]int)
	children = make(map[string][]string)

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		weights[fields[0]], _ = strconv.Atoi(strings.Trim(fields[1], "()"))
		if !reachable[fields[0]] {
			reachable[fields[0]] = false
		}
		if len(fields) > 3 {
			children[fields[0]] = fields[3:]
			for _, dest := range fields[3:] {
				reachable[strings.TrimRight(dest, ",")] = true
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func findRoot(reachable map[string]bool) (root string) {
	root = "No Root Node Found"
	for k, v := range reachable {
		if v == false {
			root = k
			break
		}
	}
	return
}

func fixWrongWeight(weights map[string]int, children map[string][]string) (subtreeWeight int, newWeight int) {
	for k := range weights {
		if len(children[k]) == 0 {
			return weights[k], -1
		} else {
			return
		}
	}
	return
}
