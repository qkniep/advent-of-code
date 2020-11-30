package main

import {
	"fmt"
	"math"
}

func main() {
	input := 265149
	fmt.Printf("Part 1: %v\n", manhattanDistance(input))
	//fmt.Printf("Part 2: %v\n", nextHighestSum(input))
}

func manhattanDistance(n int) (distance int) {
	if n != 1 {
		layer := math.Ceil((math.Sqrt(float64(n)) - 1) / 2)
		offset := math.Abs(math.Mod(float64(n)-1, 2*layer) - layer)
		distance = int(layer + offset)
	}
	return
}

/*func nextHighestSum(n int) int {
	spiral := make([]int, 0)
	x := 1
	for n <= x {
		spiral.append(x)
		x += ...
	}
	return x
}*/
