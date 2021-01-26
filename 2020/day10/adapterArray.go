package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var joltageRatings = make([]int, 0)

	// read all numbers into an array and sort them
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		joltage, _ := strconv.Atoi(scanner.Text())
		joltageRatings = append(joltageRatings, joltage)
	}
	joltageRatings = append(joltageRatings, 0)
	sort.Ints(joltageRatings)
	joltageRatings = append(joltageRatings, joltageRatings[len(joltageRatings)-1]+3)

	diff1, diff3 := countDiffs(joltageRatings)
	numArrangements := countArrangements(joltageRatings)

	fmt.Printf("Product (diff1 * diff3): %d\n", diff1*diff3)
	fmt.Printf("Number of arrangements: %d\n", numArrangements)
}

// Runs in: O(n), where n=len(nums).
func countDiffs(nums []int) (int, int) {
	var diff1, diff3 = 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			diff1++
		} else if nums[i] == nums[i-1]+3 {
			diff3++
		}
	}
	return diff1, diff3
}

// Runs in: O(n) with O(1) additional space, where n=len(nums).
func countArrangements(nums []int) int {
	var vals = []int{1, 1, 1}
	for i := 1; i < len(nums); i++ {
		inner := 0
		for j := i - 3; j < i; j++ {
			if j >= 0 && nums[i] <= nums[j]+3 {
				inner += vals[j%3]
			}
		}
		vals[i%3] = inner
	}
	return vals[(len(nums)-1)%3]
}
