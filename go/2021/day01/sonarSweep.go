package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// read integers (each on a single line) from stdin
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}

	// calculate sums over all sliding windows of length 3
	var windowSums = make([]int, len(nums)-2)
	for i := 0; i < len(nums)-2; i++ {
		windowSums[i] = nums[i] + nums[i+1] + nums[i+2]
	}

	fmt.Println("Number of measurement increases:", countIncreases(nums))
	fmt.Println("Number of sliding window increases:", countIncreases(windowSums))
}

// Counts the number of values which are larger than their predecessor
func countIncreases(nums []int) (inc int) {
	var last int = nums[0]
	for _, num := range nums[1:] {
		if num > last {
			inc++
		}
		last = num
	}
	return
}
