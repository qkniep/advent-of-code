package main

import (
	"bufio"
	"fmt"
	"os"
)

const gridSize = 100
const numLights = gridSize * gridSize
const steps = 100

func main() {
	var lightsLit1, lightsLit2 int
	var lights1, lights2 = make([]bool, numLights), make([]bool, numLights)
	var scanner = bufio.NewScanner(os.Stdin)

	// load initial state
	for y := 0; scanner.Scan(); y++ {
		for x, r := range scanner.Text() {
			if r == '#' {
				lights1[y*gridSize+x] = true
				lights2[y*gridSize+x] = true
				lightsLit1++
				lightsLit2++
			}
		}
	}

	turnOnIfNotAlready(&lights2[0], &lightsLit2)
	turnOnIfNotAlready(&lights2[gridSize-1], &lightsLit2)
	turnOnIfNotAlready(&lights2[(gridSize-1)*gridSize], &lightsLit2)
	turnOnIfNotAlready(&lights2[gridSize*gridSize-1], &lightsLit2)

	performAnimation(lights1, lights2, &lightsLit1, &lightsLit2)

	fmt.Println("Number of lights that should be lit:", lightsLit1)
	fmt.Println("Number of lights lit with broken lights:", lightsLit2)
}

func performAnimation(lights1, lights2 []bool, lightsLit1, lightsLit2 *int) {
	for step := 0; step < steps; step++ {
		newLights1, newLights2 := make([]bool, len(lights1)), make([]bool, len(lights2))
		copy(newLights1, lights1)
		copy(newLights2, lights2)
		for y := 0; y < gridSize; y++ {
			for x := 0; x < gridSize; x++ {
				var neighborsOn1, neighborsOn2 int
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if (dx == 0 && dy == 0) || x+dx < 0 || y+dy < 0 || x+dx >= gridSize || y+dy >= gridSize {
							continue
						}
						if lights1[(y+dy)*gridSize+x+dx] {
							neighborsOn1++
						}
						if lights2[(y+dy)*gridSize+x+dx] {
							neighborsOn2++
						}
					}
				}
				updateLight(&newLights1[y*gridSize+x], lights1[y*gridSize+x], neighborsOn1, lightsLit1)
				if (x == 0 && y == 0) || (x == 0 && y == gridSize-1) ||
					(x == gridSize-1 && y == 0) || (x == gridSize-1 && y == gridSize-1) {
					continue
				}
				updateLight(&newLights2[y*gridSize+x], lights2[y*gridSize+x], neighborsOn2, lightsLit2)
			}
		}
		copy(lights1, newLights1)
		copy(lights2, newLights2)
	}
}

func updateLight(newLight *bool, oldLight bool, neighbors int, count *int) {
	if oldLight && neighbors != 2 && neighbors != 3 {
		*newLight = false
		*count--
	} else if !oldLight && neighbors == 3 {
		*newLight = true
		*count++
	}
}

func turnOnIfNotAlready(light *bool, count *int) {
	if !*light {
		*light = true
		*count++
	}
}
