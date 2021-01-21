package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var replacements = make(map[string][]string, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var from, to string
		if len(scanner.Text()) == 0 {
			break
		}
		fmt.Sscanf(scanner.Text(), "%s => %s", &from, &to)
		replacements[from] = append(replacements[from], to)
	}

	scanner.Scan()
	var medicineMolecule = scanner.Text()

	oneStep := numDistinctReplacements(medicineMolecule, replacements, 1)

	fmt.Println("Number of distinct one-step replacements:", oneStep)
}

func numDistinctReplacements(start string, replacements map[string][]string, steps int) (distinct int) {
	var molecules = make(map[string]bool, 0)
	for from, toSlice := range replacements {
		for _, to := range toSlice {
			var splits = strings.Split(start, from)
			for i := 0; i < len(splits)-1; i++ {
				var molecule string
				for j := 0; j < i; j++ {
					molecule += splits[j] + from
				}
				molecule += splits[i] + to
				for j := i+1; j < len(splits); j++ {
					if j > i+1 {
						molecule += from
					}
					molecule += splits[j]
				}
				if !molecules[molecule] {
					distinct++
				}
				molecules[molecule] = true
			}
		}
	}
	return
}
