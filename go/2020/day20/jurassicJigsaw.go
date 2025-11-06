package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	top    = 0
	bottom = 1
	left   = 2
	right  = 3
)

type tile struct {
	id    int
	edges []string // 0: top, 1: bottom, 2: left, 3: right
	inner []string // horizontal lines of the binary image
}

type tileMatch struct {
	edge       int
	otherTile  int
	otherEdge  int
	needToFlip bool
}

func main() {
	var tiles = make([]tile, 150)
	var tile, tileLine = 0, 0

	// read input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if tileLine == 0 {
			// read tile ID and create slices for edges
			fmt.Sscanf(scanner.Text(), "Tile %d:", &tiles[tile].id)
			tiles[tile].edges = make([]string, 4)
			tileLine++
		} else if tileLine < 11 {
			// read tile edges
			if tileLine == 1 {
				tiles[tile].edges[top] = scanner.Text()
			} else if tileLine == 10 {
				tiles[tile].edges[bottom] = scanner.Text()
			} else {
				tiles[tile].inner = append(tiles[tile].inner, scanner.Text()[1:len(scanner.Text())-1])
			}
			tiles[tile].edges[left] += scanner.Text()[0:1]
			tiles[tile].edges[right] += scanner.Text()[len(scanner.Text())-1:]
			tileLine++
		} else { // empty input line between tiles
			tile++
			tileLine = 0
		}
	}

	// find corner tiles
	var result = 1
	var cornerTile int
	for tileIndex, tile := range tiles {
		if tile.id == 0 {
			continue
		}
		if len(findMatchingEdges(tile, tiles)) == 2 {
			result *= tile.id
			cornerTile = tileIndex
		}
	}
	fmt.Printf("Product of corner tile IDs: %v\n", result)

	// solve jigsaw, generate final image, and find sea monsters
	var solvedJigsaw = solveJigsaw(tiles, cornerTile)
	var monsters = 0
	for i := 0; monsters == 0 && i < 8; i++ {
		monsters = countSeaMonsters(solvedJigsaw)
		solvedJigsaw = rotateCW(solvedJigsaw, 90)
		if i == 4 {
			flipVertically(solvedJigsaw)
		}
	}
	if monsters == 0 {
		panic("No sea monsters found!")
	}
	roughness := calculateRoughness(solvedJigsaw, monsters)
	fmt.Printf("Sea monster habitat roughness: %v\n", roughness)
}

// The number of ways in which we can attach the given tile to other tiles.
func findMatchingEdges(t tile, tiles []tile) (matches []tileMatch) {
	for t0Index, t0 := range tiles {
		if t0.id == 0 || t0.id == t.id {
			continue
		}
		for eID, edge := range t.edges {
			for e0ID, edge0 := range t0.edges {
				if edge == edge0 || edge == reverse(edge0) {
					needToFlip := edge == reverse(edge0)
					if eID == e0ID ||
						(eID == bottom && e0ID == left) || (eID == left && e0ID == bottom) ||
						(eID == top && e0ID == right) || (eID == right && e0ID == top) {
						needToFlip = edge == edge0
					}
					matches = append(matches, tileMatch{eID, t0Index, e0ID, needToFlip})
				}
			}
		}
	}
	return
}

// The startingTile has to be the index in the tiles slice of a corner tile.
func solveJigsaw(tiles []tile, startingTile int) (solved []string) {
	var current = startingTile
	var row []int

	// rotate starting tile to make it the top-left corner
	matches := findMatchingEdges(tiles[startingTile], tiles)
	if matches[0].edge == top || matches[1].edge == top {
		if matches[0].edge == left || matches[1].edge == left {
			rotateTile(&tiles[startingTile], 180)
		} else {
			rotateTile(&tiles[startingTile], 90)
		}
	} else if matches[0].edge == left || matches[1].edge == left {
		rotateTile(&tiles[startingTile], 270)
	}

	// build first row of final image
	solved = placeTile(solved, tiles, startingTile, true)
	for x := 0; ; x++ {
		row = append(row, current)
		matches := findMatchingEdges(tiles[current], tiles)
		if matches[0].edge != right && matches[1].edge != right && (len(matches) < 3 || matches[2].edge != right) {
			break
		}
		// find the correct match object
		match := matches[0]
		if matches[1].edge == right {
			match = matches[1]
		} else if len(matches) >= 3 && matches[2].edge == right {
			match = matches[2]
		}
		// rotate the new tile accordingly
		if match.otherEdge == bottom {
			rotateTile(&tiles[match.otherTile], 90)
		} else if match.otherEdge == right {
			rotateTile(&tiles[match.otherTile], 180)
		} else if match.otherEdge == top {
			rotateTile(&tiles[match.otherTile], 270)
		}
		if match.needToFlip {
			flipTileVert(&tiles[match.otherTile])
		}
		// add the tile into the final image
		solved = placeTile(solved, tiles, match.otherTile, false)
		current = match.otherTile
	}

	// build remaining rows of final image
	for y := 1; y < len(row); y++ {
		for x := 0; x < len(row); x++ {
			matches := findMatchingEdges(tiles[row[x]], tiles)
			for _, match := range matches {
				if match.edge == bottom {
					row[x] = match.otherTile
					// rotate the new tile accordingly
					if match.otherEdge == left {
						rotateTile(&tiles[match.otherTile], 90)
					} else if match.otherEdge == bottom {
						rotateTile(&tiles[match.otherTile], 180)
					} else if match.otherEdge == right {
						rotateTile(&tiles[match.otherTile], 270)
					}
					if match.needToFlip {
						flipTileHori(&tiles[match.otherTile])
					}
					solved = placeTile(solved, tiles, match.otherTile, (x == 0))
					break
				}
			}
		}
	}

	return
}

func rotateTile(t *tile, degrees int) {
	(*t).inner = rotateCW((*t).inner, degrees)
	for degLeft := degrees; degLeft > 0; degLeft -= 90 {
		oldEdges := make([]string, len((*t).edges))
		copy(oldEdges, (*t).edges)
		(*t).edges[top] = reverse(oldEdges[left])
		(*t).edges[bottom] = reverse(oldEdges[right])
		(*t).edges[left] = oldEdges[bottom]
		(*t).edges[right] = oldEdges[top]
	}
}

func flipTileVert(t *tile) {
	flipVertically((*t).inner)
	oldEdges := make([]string, len((*t).edges))
	copy(oldEdges, (*t).edges)
	(*t).edges[top] = oldEdges[bottom]
	(*t).edges[bottom] = oldEdges[top]
	(*t).edges[left] = reverse(oldEdges[left])
	(*t).edges[right] = reverse(oldEdges[right])
}

func flipTileHori(t *tile) {
	flipHorizontally((*t).inner)
	oldEdges := make([]string, len((*t).edges))
	copy(oldEdges, (*t).edges)
	(*t).edges[top] = reverse(oldEdges[top])
	(*t).edges[bottom] = reverse(oldEdges[bottom])
	(*t).edges[left] = oldEdges[right]
	(*t).edges[right] = oldEdges[left]
}

func placeTile(image []string, tiles []tile, tileIndex int, newRow bool) []string {
	if newRow {
		for _, row := range tiles[tileIndex].inner {
			image = append(image, row)
		}
	} else {
		for innerY, row := range tiles[tileIndex].inner {
			y := len(image) - len(tiles[tileIndex].inner) + innerY
			image[y] += row
		}
	}
	return image
}

// Mutates image internally.
func rotateCW(image []string, degrees int) []string {
	if degrees >= 180 {
		flipHorizontally(image)
		flipVertically(image)
	}
	if degrees%180 == 90 {
		outImage := make([]string, len(image))
		for y := len(image) - 1; y >= 0; y-- {
			for x := 0; x < len(image[0]); x++ {
				outImage[x] += string(rune(image[y][x]))
			}
		}
		return outImage
	}
	return image
}

// Mutates image in place.
func flipHorizontally(image []string) {
	for i, row := range image {
		image[i] = reverse(row)
	}
}

// Mutates image in place.
func flipVertically(image []string) {
	for i := 0; i < len(image)/2; i++ {
		j := len(image) - i - 1
		image[i], image[j] = image[j], image[i]
	}
}

var seaMonster = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
}

func countSeaMonsters(image []string) (monsters int) {
	for x := 0; x < len(image[0])-len(seaMonster[0]); x++ {
		for y := 0; y < len(image)-len(seaMonster); y++ {
			if matchSeaMonster(image, x, y) {
				monsters++
			}
		}
	}
	return
}

func matchSeaMonster(image []string, xOffset, yOffset int) bool {
	for x := 0; x < len(seaMonster[0]); x++ {
		for y := 0; y < len(seaMonster); y++ {
			if seaMonster[y][x] == byte('#') && image[yOffset+y][xOffset+x] != byte('#') {
				return false
			}
		}
	}
	return true
}

func calculateRoughness(image []string, monsters int) int {
	var roughness = 0
	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image[0]); x++ {
			if image[y][x] == byte('#') {
				roughness++
			}
		}
	}
	return roughness - monsters*15
}

// Reverses the given string, returning a string with all the runes in reverse order.
func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, r := range s {
		n--
		runes[n] = r
	}
	return string(runes[n:])
}
