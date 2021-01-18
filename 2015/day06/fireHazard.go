package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pos struct {
	x, y int
}

func main() {
	var lights, lights2 = make(map[pos]bool, 0), make(map[pos]int, 0)
	var lightsLit, totalBrightness int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// read one instruction
		var minX, minY, maxX, maxY int
		parsed := strings.Fields(scanner.Text())
		command := parsed[0]
		parsed = parsed[1:]
		if command == "turn" {
			command += " " + parsed[0]
			parsed = parsed[1:]
		}
		fmt.Sscanf(parsed[0], "%d,%d", &minX, &minY)
		fmt.Sscanf(parsed[2], "%d,%d", &maxX, &maxY)

		// perform the instruction
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				p := pos{x, y}
				switch command {
				case "turn on":
					if !lights[p] {
						lights[p] = true
						lightsLit++
					}
					lights2[p]++
					totalBrightness++
				case "turn off":
					if lights[p] {
						lights[p] = false
						lightsLit--
					}
					if lights2[p] > 0 {
						lights2[p]--
						totalBrightness--
					}
				case "toggle":
					lights[p] = !lights[p]
					if lights[p] {
						lightsLit++
					} else {
						lightsLit--
					}
					lights2[p] += 2
					totalBrightness += 2
				}
			}
		}
	}

	fmt.Println("Number of lights lit:", lightsLit)
	fmt.Println("Total brightness:", totalBrightness)
}
