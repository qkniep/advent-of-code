package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var fuelSum, recursiveFuelSum = 0, 0

	// read all mass numbers and calculate the fuel requirements, keeping track of the sums
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		fuelSum += fuelRequirement(mass)
		recursiveFuelSum += recursiveFuelRequirement(mass)
	}

	fmt.Printf("Fuel requirement for mass: %d\n", fuelSum)
	fmt.Printf("Recursive fuel requirements: %d\n", recursiveFuelSum)
}

func fuelRequirement(mass int) int {
	return mass/3 - 2
}

// Calculates the amount of fuel needed for the base mass and all the fuel.
func recursiveFuelRequirement(mass int) int {
	fuel := fuelRequirement(mass)
	if fuel <= 0 {
		return 0
	}
	return fuel + recursiveFuelRequirement(fuel)
}
