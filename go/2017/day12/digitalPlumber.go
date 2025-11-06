package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var graph = make(map[string][]string, 0)

	// interpret input as adjacency list for an undirected graph
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		source := fields[0]
		dests := fields[2:]
		for _, d := range dests {
			graph[source] = append(graph[source], strings.TrimRight(d, ","))
		}
	}

	fmt.Printf("Size of 0-group: %v\n", sizeOfGroup("0", graph, make(map[string]bool, 0)))
	fmt.Printf("Number of groups: %v\n", numberOfGroups(graph))
}

// Explore the group (connected component) containing element, using the graph adjacency list.
// All visited nodes are marked with true in visited slice.
// Returns the size of the group.
func sizeOfGroup(element string, graph map[string][]string, visited map[string]bool) (size int) {
	var toVisit = []string{element}

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

// Find out how many groups (connected components) there are in the graph.
func numberOfGroups(graph map[string][]string) (num int) {
	var visited = make(map[string]bool, 0)

	for k := range graph {
		if !visited[k] {
			sizeOfGroup(k, graph, visited)
			num++
		}
	}
	return
}
