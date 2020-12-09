package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var nums = make([]int, 0)

	// read all numbers into an array
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}

	nonSum := findNonSum(nums, 25)
	encWeakness := findEncWeakness(nums, nonSum)

	fmt.Printf("Acc value: %d\n", nonSum)
	fmt.Printf("Fixed Acc value: %d\n", encWeakness)
}

// Finds the first entry in `nums` which is not the sum of any two of the `slide` entries
// immediately before it, ignoring the `slide` first entries.
// Runs in: O(n*(slide^2)), where n=len(nums).
func findNonSum(nums []int, slide int) int {
	for i := slide; i < len(nums); i++ {
		sumOfTwo := false
		for a := 1; a < slide; a++ {
			for b := 0; b < a; b++ {
				if nums[i+a] + nums[i+b] == nums[i+slide] {
					sumOfTwo = true
				}
			}
		}
		if !sumOfTwo {
			return nums[i+slide]
		}
	}
	return -1
}

// Finds the first contiguous range of entries in `nums` which sum to `sum`, returning the sum of
// the smallest and largest number in that range.
// Runs in: O(n^2), where n=len(nums).
func findEncWeakness(nums []int, sum int) int {
	for i := 0; i < len(nums); i++ {
		runningSum := nums[i]
		for j := i+1; j < len(nums); j++ {
			runningSum += nums[j]
			if runningSum == sum {
				sorted := make([]int, j-i+1)
				copy(sorted, nums[i:j+1])
				sort.Ints(sorted)
				return sorted[0]+sorted[len(sorted)-1]
			} else if runningSum > sum {
				break
			}
		}
	}
	return -1
}
