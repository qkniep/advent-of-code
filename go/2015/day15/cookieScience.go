package main

import (
	"bufio"
	"fmt"
	"os"
)

const inFormat = "%s capacity %d, durability %d, flavor %d, texture %d, calories %d"

func main() {
	var ingredients [][]int
	var scanner = bufio.NewScanner(os.Stdin)

	// read list of ingredients
	for scanner.Scan() {
		var _name string
		var stats = make([]int, 5)
		fmt.Sscanf(scanner.Text(), inFormat, &_name, &stats[0], &stats[1], &stats[2], &stats[3], &stats[4])
		ingredients = append(ingredients, make([]int, 5))
		copy(ingredients[len(ingredients)-1], stats)
	}

	var bestCookie, bestCookie500Calories int
	partitions(100, len(ingredients), func(recipe []int) {
		score, calories := cookieScoring(recipe, ingredients)
		if score > bestCookie {
			bestCookie = score
		}
		if score > bestCookie500Calories && calories == 500 {
			bestCookie500Calories = score
		}
	})

	fmt.Println("Best scoring cookie overall:", bestCookie)
	fmt.Println("Best cookie with exactly 500 calories:", bestCookie500Calories)
}

// Score a cookie according to the rules from the problem statement.
// Returns the score (based on capacity, durability, flavor, texture) and the total calories.
func cookieScoring(recipe []int, ingredients [][]int) (int, int) {
	var capacity, durability, flavor, texture, calories int

	for ingredient, amount := range recipe {
		capacity += ingredients[ingredient][0] * amount
		durability += ingredients[ingredient][1] * amount
		flavor += ingredients[ingredient][2] * amount
		texture += ingredients[ingredient][3] * amount
		calories += ingredients[ingredient][4] * amount
	}

	if capacity <= 0 || durability <= 0 || flavor <= 0 || texture <= 0 {
		return 0, calories
	}
	return capacity * durability * flavor * texture, calories
}

// Applies function f to all partitions of n into k parts.
// Such a partition is a slice of size k with a sum of n.
func partitions(n, k int, f func([]int)) {
	part(n, k, f, make([]int, k))
}

func part(n, k int, f func([]int), r []int) {
	if k == 1 {
		r[0] = n
		f(r)
		return
	}
	for i := 0; i <= n; i++ {
		r[k-1] = i
		part(n-i, k-1, f, r)
	}
}
