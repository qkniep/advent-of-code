package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var distances = make(map[string]int, 0)
	var orbitedBy = make(map[string][]string, 0)
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		planets := strings.Split(scanner.Text(), ")")
		/*d, ok := distances[planets[0]]
		if ok {
			distances[planets[1]] = d + 1
		} else {*/
			orbitedBy[planets[0]] = append(orbitedBy[planets[0]], planets[1])
		//}
	}

	fmt.Println(orbitedBy)

	sum := 0
	toVisit := []string{"COM"}
	depths := []int{0}
	currentDepth := 0
	for len(toVisit) > 0 {
		sum += depths[0]
		currentDepth, depths = depths[0], depths[1:]
		for i := 0; i < len(orbitedBy[toVisit[0]]); i++ {
			depths = append(depths, currentDepth + 1)
		}
		toVisit = append(toVisit[1:], orbitedBy[toVisit[0]]...)
	}

	fmt.Println("Distance to origin of closest crossing:", sum)
	//fmt.Println("Number of steps until first crossing:", findMinCrossingSteps(wire1, wire2))
}
