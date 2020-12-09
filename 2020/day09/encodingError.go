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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}

	fmt.Printf("Acc value: %d\n", findNonSum(nums))
	fmt.Printf("Fixed Acc value: %d\n", findEncWeakness(nums, 90433990))
}

func findNonSum(nums []int) int {
	for i := 25; i < len(nums); i++ {
		sumOfTwo := false
		for a := 1; a < 25; a++ {
			for b := 0; b < a; b++ {
				if nums[i+a] + nums[i+b] == nums[i+25] {
					sumOfTwo = true
				}
			}
		}
		if !sumOfTwo {
			return nums[i+25]
		}
	}
	return -1
}

func findEncWeakness(nums []int, sum int) int {
	for i := 0; i < len(nums); i++ {
		runningSum := nums[i]
		for j := i+1; j < len(nums); j++ {
			runningSum += nums[j]
			if runningSum == sum {
				sorted := nums[i:j+1]
				sort.Ints(sorted)
				return sorted[0]+sorted[len(sorted)-1]
			} else if runningSum > sum {
				break
			}
		}
	}
	return -1
}
