package main

import "fmt"

func main() {
	var presents, house int

	fmt.Scanf("%d", &presents)

	house = findSmallestWithDivisorSum(presents / 10)

	fmt.Printf("First house to receive %d presents: %d\n", presents, house)
	//fmt.Println("Total brightness:", totalBrightness)
}

// Returns the smallest number where all its divisors added together produce the given sum.
func findSmallestWithDivisorSum(divisorSum int) int {
	//var primes []int
	for i := 2; ; i++ {
		if i % 100000 == 0 {
			fmt.Println(i)
		}
		// find prime factors
		/*for _, p := range primes {
			if i % p == 0 {
				for a := p; a <= i; a += p {
				sum += p
				if sum >= factorSum {
					break
				}
			}
		}*/
		// calculate sum of all factors
		var sum = i+1 // 1 and i itself are divisors of every i
		for j := 2; j*j <= i; j++ {
			if i % j == 0 {
				sum += j
				if j*j < i {
					sum += i/j
				}
			}
		}
		if i < 20 {
			fmt.Printf("Divisor sum of %v: %v\n", i, sum)
		}
		if sum >= divisorSum {
			return i
		}/* else if sum == i+1 {
			primes = append(primes, i)
		}*/
	}
}
