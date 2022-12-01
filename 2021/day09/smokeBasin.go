package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// read heightmap
	var heightmap [][]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		heightmap = append(heightmap, []int{})
		for _, r := range scanner.Text() {
			heightmap[len(heightmap)-1] = append(heightmap[len(heightmap)-1], int(r)-48)
		}
	}

	// identify low points, sum up their risk levels, and find their basin
	var riskSum = 0
	var basinSizes []int
	for y, row := range heightmap {
		for x, cell := range row {
			if y-1 >= 0 && cell >= heightmap[y-1][x] {
				continue
			} else if y+1 < len(heightmap) && cell >= heightmap[y+1][x] {
				continue
			} else if x-1 >= 0 && cell >= row[x-1] {
				continue
			} else if x+1 < len(row) && cell >= row[x+1] {
				continue
			}
			// found low point
			riskSum += cell + 1
			basinSizes = append(basinSizes, growBasin(heightmap, x, y))
		}
	}

	// multiply top 3 basin sizes
	var basinSizeProduct = 1
	sort.Ints(basinSizes)
	for _, size := range basinSizes[len(basinSizes)-3:] {
		basinSizeProduct *= size
	}

	fmt.Println("Sum of risk levels:", riskSum)
	fmt.Println("Product of top 3 basin sizes:", basinSizeProduct)
}

func growBasin(heightmap [][]int, x int, y int) (size int) {
	var visited = make([]bool, len(heightmap)*len(heightmap[0]))
	var toCheck = [][]int{{x, y}}
	var point []int
	for len(toCheck) > 0 {
		point, toCheck = toCheck[0], toCheck[1:]
		x, y = point[0], point[1]

		if visited[y*len(heightmap[0])+x] {
			continue
		}
		visited[y*len(heightmap[0])+x] = true
		size++

		// add neighbors as to visit
		if y-1 >= 0 && heightmap[y-1][x] < 9 {
			toCheck = append(toCheck, []int{x, y - 1})
		}
		if y+1 < len(heightmap) && heightmap[y+1][x] < 9 {
			toCheck = append(toCheck, []int{x, y + 1})
		}
		if x-1 >= 0 && heightmap[y][x-1] < 9 {
			toCheck = append(toCheck, []int{x - 1, y})
		}
		if x+1 < len(heightmap[0]) && heightmap[y][x+1] < 9 {
			toCheck = append(toCheck, []int{x + 1, y})
		}
	}
	return
}
