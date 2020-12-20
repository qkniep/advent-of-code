package main

import (
	"bufio"
	"fmt"
	"os"
)

type tile struct {
	id    int
	edges []string // 0: top, 1: bottom, 2: left, 3: right
}

func main() {
	var tiles = make([]tile, 150)
	var tile, tileLine = 0, 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if tileLine == 0 {
			fmt.Sscanf(scanner.Text(), "Tile %d:", &tiles[tile].id)
			tiles[tile].edges = make([]string, 4)
			tileLine++
		} else if tileLine < 11 {
			if tileLine == 1 {
				tiles[tile].edges[0] = scanner.Text()
			} else if tileLine == 10 {
				tiles[tile].edges[1] = scanner.Text()
			}
			tiles[tile].edges[2] += scanner.Text()[0:1]
			tiles[tile].edges[3] += scanner.Text()[len(scanner.Text())-1:]
			tileLine++
		} else {
			tile++
			tileLine = 0
		}
	}

	// find corner tiles
	var result = 1
	for _, tile := range tiles {
		if tile.id == 0 {
			continue
		}
		if countMatchingEdges(tile, tiles) == 2 {
			result *= tile.id
		}
	}

	fmt.Printf("Product of corner tile IDs: %v\n", result)
}

func countMatchingEdges(t tile, tiles []tile) (matches int) {
	for _, t0 := range tiles {
		if t0.id == 0 || t0.id == t.id {
			continue
		}
		for _, edge := range t.edges {
			for _, edge0 := range t0.edges {
				if edge == edge0 || reverse(edge) == reverse(edge0) || edge == reverse(edge0) || edge0 == reverse(edge) {
					matches++
				}
			}
		}
	}
	return
}

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, r := range s {
		n--
		runes[n] = r
	}
	return string(runes[n:])
}
