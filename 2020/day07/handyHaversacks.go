package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bags struct {
	num int;
	color string;
}

func main() {
	var rules = make(map[string][]string, 0)
	var rulesCount = make(map[string][]bags, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var words = strings.Fields(scanner.Text())
		bagColor := fmt.Sprintf("%s %s", words[0], words[1])
		for i := 5; i + 2 < len(words); i += 4 {
			color := fmt.Sprintf("%s %s", words[i], words[i+1])
			rules[color] = append(rules[color], bagColor)
			num, _ := strconv.Atoi(words[i-1])
			rulesCount[bagColor] = append(rulesCount[bagColor], bags{num, color})
		}
	}

	var numPossibleOuter = len(possibleOuter(rules, "shiny gold"))
	var numTotalInner = totalInner(rulesCount, "shiny gold")

	fmt.Printf("Possible outer colors: %d\n", numPossibleOuter)
	fmt.Printf("Everyone answered yes: %d\n", numTotalInner)
}

func possibleOuter(rules map[string][]string, color string) []string {
	added := make(map[string]bool, 0)
	outer, dirty := make([]string, 0), rules[color]

	for _, c := range rules[color] {
		added[c] = true
	}

	for len(dirty) > 0 {
		d := dirty[0]
		outer = append(outer, d)
		for _, c := range rules[d] {
			if !added[c] {
				dirty = append(dirty, c)
				added[c] = true
			}
		}
		dirty = dirty[1:]
	}
	return outer
}

func totalInner(rules map[string][]bags, color string) int {
	sum := 0
	for _, b := range rules[color] {
		sum += b.num + b.num * totalInner(rules, b.color)
	}
	return sum
}
