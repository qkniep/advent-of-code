package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var countIngredient = make(map[string]int, 0)
	var countAllergen = make(map[string]int, 0)
	var possibleForAllergen = make(map[string][]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var allergens []string = nil
		ingredients := strings.Fields(scanner.Text())
		for i, ingr := range ingredients {
			if ingr == "(contains" {
				allergens = ingredients[i+1:]
				ingredients = ingredients[:i]
				break
			}
			countIngredient[ingr]++
		}
		for i, allergen := range allergens {
			allergens[i] = allergen[:len(allergen)-1]
			countAllergen[allergens[i]]++
			if countAllergen[allergens[i]] == 1 {
				possibleForAllergen[allergens[i]] = ingredients
				break
			}
			// intersection
			newPossible := make([]string, 0)
			for _, p1 := range possibleForAllergen[allergens[i]] {
				inBoth := false
				for _, p2 := range ingredients {
					if p1 == p2 {
						inBoth = true
					}
				}
				if inBoth {
					newPossible = append(newPossible, p1)
				}
			}
			possibleForAllergen[allergens[i]] = newPossible
			fmt.Printf("poss[%v] = %v\n", allergens[i], newPossible)
		}
	}

	var totalOccurrences = 0
	for ingr, num := range countIngredient {
		defNoAllergen := true
		for i := range possibleForAllergen {
			for _, ingr0 := range possibleForAllergen[i] {
				if ingr0 == ingr {
					defNoAllergen = false
				}
			}
		}
		if defNoAllergen {
			totalOccurrences += num
		}
	}

	fmt.Printf("%v\n", totalOccurrences)
}
