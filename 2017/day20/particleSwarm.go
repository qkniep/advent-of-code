package main

import (
	"bufio"
	"fmt"
	"os"
)

type vec3d struct {
	x, y, z int
}

const inFormat = "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>"

func main() {
	var candidates []int
	var minAcc vec3d
	var scanner = bufio.NewScanner(os.Stdin)

	for i := 0; scanner.Scan(); i++ {
		var p, v, a vec3d
		fmt.Sscanf(scanner.Text(), inFormat, &p.x, &p.y, &p.z, &v.x, &v.y, &v.z, &a.x, &a.y, &a.z)
		if len(candidates) == 0 || vecAbs(a) < vecAbs(minAcc) {
			candidates = []int{i}
			minAcc = a
		} else if vecAbs(a) == vecAbs(minAcc) {
			candidates = append(candidates, i)
		}
	}

	fmt.Printf("Squares used on disk: %v\n", candidates)
	//fmt.Printf("Connected regions: %v\n", connectedRegions)
}

func vecAbs(v vec3d) int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
