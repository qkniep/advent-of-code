package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var countIngredient = make(map[string]int, 0)
	var countAllergen = make(map[string]int, 0)
	var possibleForAllergen = make(map[string][]string, 0)
	var allAllergens = make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// read food line (list of ingredients, followed by list of allergens)
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
			if countAllergen[allergens[i]] == 1 { // first time this allergen appears
				allAllergens = append(allAllergens, allergens[i])
				possibleForAllergen[allergens[i]] = ingredients
				continue
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

	// eliminate ingredients known to belong to one allergen from the others
	known := make(map[string]bool, 0)
	for len(known) < len(possibleForAllergen) {
		for i := range possibleForAllergen {
			count, value := 0, ""
			for _, ingr := range possibleForAllergen[i] {
				if !known[ingr] {
					count++
					value = ingr
				}
			}
			if count == 1 {
				possibleForAllergen[i] = []string{value}
				known[value] = true
			}
		}
	}

	var totalOccurrences = 0
	for ingr, num := range countIngredient {
		if !known[ingr] {
			totalOccurrences += num
		}
	}

	dangerousIngredients := make([]string, 0)
	sort.Strings(allAllergens)
	for _, allergen := range allAllergens {
		dangerousIngredients = append(dangerousIngredients, possibleForAllergen[allergen][0])
	}

	fmt.Printf("%v\n", totalOccurrences)
	fmt.Printf("%v\n", strings.Join(dangerousIngredients, ","))
}
