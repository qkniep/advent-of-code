package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bags struct {
	num   int
	color string
}

func main() {
	var rulesOuter = make(map[string][]string, 0)
	var rulesCount = make(map[string][]bags, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var words = strings.Fields(scanner.Text())
		outColor := words[0] + " " + words[1]
		// parse all inner bags for this one outer bag, add to rules
		for i := 5; i+2 < len(words); i += 4 {
			inColor := words[i] + " " + words[i+1]
			num, _ := strconv.Atoi(words[i-1])
			rulesOuter[inColor] = append(rulesOuter[inColor], outColor)
			rulesCount[outColor] = append(rulesCount[outColor], bags{num, inColor})
		}
	}

	fmt.Printf("Possible outer bag colors: %d\n", possibleOuter(rulesOuter, "shiny gold"))
	fmt.Printf("Total number of inner bags: %d\n", totalInner(rulesCount, "shiny gold"))
}

// Gives the number of possible outer bags which contain a `color` bag somewhere inside.
func possibleOuter(rules map[string][]string, color string) int {
	added := make(map[string]bool, 0)
	dirty := rules[color]
	count := 0

	for _, c := range rules[color] {
		added[c] = true
	}

	for d := dirty[0]; len(dirty) > 0; d, dirty = dirty[0], dirty[1:] {
		for _, c := range rules[d] {
			if !added[c] {
				dirty = append(dirty, c)
				added[c] = true
			}
		}
		count++
	}
	return count
}

// Gives the total number of bags which are needed inside one `color` bag, based on `rules`.
func totalInner(rules map[string][]bags, color string) int {
	sum := 0
	for _, b := range rules[color] {
		sum += b.num + b.num*totalInner(rules, b.color)
	}
	return sum
}
