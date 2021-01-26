package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var wantedSue = map[string]int{
	"children": 3,
	"cats": 7,
	"samoyeds": 2,
	"pomeranians": 3,
	"akitas": 0,
	"vizslas": 0,
	"goldfish": 5,
	"trees": 3,
	"cars": 2,
	"perfumes": 1,
}

func main() {
	var sue, correctSue1, correctSue2 int
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() && (correctSue1 == 0 || correctSue2 == 0) {
		fields := strings.Fields(scanner.Text())
		fmt.Sscanf(fields[1], "%d:", &sue)

		// compare with the wanted sue
		var correctFor1, correctFor2 = true, true
		for i := 3; i < len(fields); i += 2 {
			fact := strings.TrimRight(fields[i-1], ":")
			num, _ := strconv.Atoi(strings.TrimRight(fields[i], ","))
			if wantedSue[fact] != num {
				correctFor1 = false
			}
			if fact == "cats" || fact == "trees" {
				if wantedSue[fact] >= num {
					correctFor2 = false
				}
			} else if fact == "pomeranians" || fact == "goldfish" {
				if wantedSue[fact] <= num {
					correctFor2 = false
				}
			} else if wantedSue[fact] != num {
				correctFor2 = false
			}
		}

		if correctFor1 {
			correctSue1 = sue
		}
		if correctFor2 {
			correctSue2 = sue
		}
	}

	fmt.Println("Aunt Sue who matches exactly:", correctSue1)
	fmt.Println("Aunt sue who matches ranges:", correctSue2)
}
