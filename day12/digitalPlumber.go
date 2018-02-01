package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	graph := readInput()
	fmt.Printf("Size of 0-Group: %v\n", sizeOfGroup("0", graph))
	//fmt.Printf("Furthest Distance: %v\n", furthestDistance)
}

func readInput() (graph map[string][]string) {
	graph = make(map[string][]string, 0)

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		source := fields[0]
		dests := fields[2:]
		for _, d := range dests {
			graph[source] = append(graph[source], strings.TrimRight(d, ","))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func sizeOfGroup(element string, graph map[string][]string) (size int) {
	visited := make(map[string]bool, 0)
	toVisit := []string{element}
	for len(toVisit) > 0 {
		size++
		neighbors := graph[toVisit[0]]
		visited[toVisit[0]] = true
		toVisit = toVisit[1:]
		for _, n := range neighbors {
			if !visited[n] {
				toVisit = append(toVisit, n)
			}
		}
	}
	return
}
