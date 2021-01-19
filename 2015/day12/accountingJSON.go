package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var jsonObj interface{}
	json.Unmarshal([]byte(scanner.Text()), &jsonObj)

	sumTotal, sumNoRed, _ := sum(jsonObj)

	fmt.Println("", sumTotal)
	fmt.Println("", sumNoRed)
}

// Sum all numbers in the given JSON entity, returning the total sum.
// Also returns the sum excluding all objects (and their children)
// where some attribute's value is set to "red".
func sum(obj interface{}) (total, noRed float64, red bool) {
	switch v := obj.(type) {
	case string:
		if v == "red" {
			return 0, 0, true
		}
	case float64:
		total, noRed = v, v
	case []interface{}:
		for _, e := range v {
			t, nr, _ := sum(e)
			total += t
			noRed += nr
		}
	case map[string]interface{}:
		var hasSomethingRed bool
		for _, e := range v {
			t, nr, isRed := sum(e)
			total += t
			noRed += nr
			if isRed {
				hasSomethingRed = true
			}
		}
		if hasSomethingRed {
			return total, 0, false
		}
	}
	return
}
