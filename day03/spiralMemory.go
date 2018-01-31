package main

import "fmt"
import "math"

func main() {
	input := 265149
	fmt.Printf("Part 1: %v\n", manhattanDistance(input))
	//fmt.Printf("Sum 2: %v\n", sumOfPairs(input, len(input)/2))
}

func manhattanDistance(n int) (distance int) {
	if n == 1 {
		return 0
	} else {
		layer := math.Ceil((math.Sqrt(float64(n)) - 1) / 2)
		offset := math.Abs(math.Mod(float64(n)-1, 2*layer) - layer)
		return int(layer + offset)
	}
}
