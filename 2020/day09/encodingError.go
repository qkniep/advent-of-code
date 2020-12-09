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

	fmt.Printf("First non 2/25 sum: %d\n", nonSum)
	fmt.Printf("Encryption weakness: %d\n", encWeakness)
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
// Runs in: O(n), where n=len(nums).
func findEncWeakness(nums []int, sum int) int {
	min, max, runningSum := 0, 0, nums[0]
	for runningSum != sum {
		if runningSum < sum || min == max {
			max++
			runningSum += nums[max]
		}
		if runningSum > sum {
			runningSum -= nums[min]
			min++
		}
	}
	sorted := make([]int, max-min+1)
	copy(sorted, nums[min:max+1])
	sort.Ints(sorted)
	return sorted[0]+sorted[len(sorted)-1]
}
