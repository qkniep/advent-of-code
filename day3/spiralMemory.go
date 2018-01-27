package main

import "fmt"
import "math"

func main() {
	input := 265149
	fmt.Printf("Part 1: %v\n", manhattanDistance(input))
	//fmt.Printf("Sum 2: %v\n", sumOfPairs(input, len(input)/2))
}

func manhattanDistance(n int) (distance int) {
	layer := math.Ceil((math.Sqrt(float64(n)) - 1) / 2)
	upperRight := 4*layer*layer - 2*layer + 1
	distance = int(layer + math.Abs(math.Mod(float64(n)-upperRight, 2*layer)-layer))
	return
}
