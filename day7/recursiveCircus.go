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
	rootNode := findRoot(reachable)
	fmt.Printf("Root Program: %v\n", rootNode)
	_, newWeight := fixWrongWeight(rootNode, weights, children)
	fmt.Printf("Correct Weight: %v\n", newWeight)
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
			for _, dest := range fields[3:] {
				dest = strings.TrimRight(dest, ",")
				reachable[dest] = true
				children[fields[0]] = append(children[fields[0]], dest)
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

func fixWrongWeight(node string, weights map[string]int, children map[string][]string) (stw int, fixed int) {
	if len(children[node]) == 0 {
		return weights[node], -1
	}
	var childWeightSum int = 0
	var aValue, aWeight int = -1, -1
	var foundDifferent bool = false
	for _, child := range children[node] {
		weight, newWeight := fixWrongWeight(child, weights, children)
		if newWeight >= 0 {
			return -1, newWeight
		}
		if aValue == -1 {
			aValue = weight
			aWeight = weights[child]
		} else if weight != aValue {
			if !foundDifferent {
				foundDifferent = true
				fixed = weights[child] + (aValue - weight)
			} else {
				return -1, aWeight + (weight - aValue)
			}
		}
		childWeightSum += weight
	}
	if foundDifferent {
		return -1, fixed
	}
	return weights[node] + childWeightSum, -1
}
