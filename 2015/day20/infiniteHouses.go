package main

import (
	"fmt"
	"math"
)

func main() {
	var presents, house1, house2 int64
	fmt.Scanf("%d", &presents)

	house1 = findSmallestWithDivisorSum(presents/10, math.MaxInt64)
	house2 = findSmallestWithDivisorSum(presents/11, 50)

	fmt.Printf("First house (%d presents, 10 per Elf, no max): %d\n", presents, house1)
	fmt.Printf("First house (%d presents, 11 per Elf, 50 houses max): %d\n", presents, house2)
}

// Returns the smallest number where all its divisors added together produce the given sum.
// Also ignores all divisors for which the quotient is larger than max.
// Simply set max to math.MaxInt64 to not use this feature.
func findSmallestWithDivisorSum(divisorSum int64, max int64) int64 {
	for i := int64(2); ; i++ {
		// calculate sum of all factors
		var sum int64
		for j := int64(1); j*j <= i; j++ {
			if i%j == 0 {
				if i/j <= max {
					sum += j
				}
				if j*j < i && j <= max {
					sum += i / j
				}
			}
		}
		if sum >= divisorSum {
			return i
		}
	}
}
