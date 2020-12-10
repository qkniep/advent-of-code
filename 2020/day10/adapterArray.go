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

	fmt.Printf("Product (diff1 * diff3): %d\n", diff1 * diff3)
	fmt.Printf("Number of arrangements: %d\n", numArrangements)
}

// Runs in: O(n), where n=len(nums).
func countDiffs(nums []int) (int, int) {
	var diff1, diff3 = 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] + 1 {
			diff1++
		} else if nums[i] == nums[i-1] + 3 {
			diff3++
		}
	}
	return diff1, diff3
}

// Runs in: O(n^2), where n=len(nums).
func countArrangements(nums []int) int {
	var vals = []int{1}
	for i := 1; i < len(nums); i++ {
		inner := 0
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] + 3 {
				inner += vals[j]
			}
		}
		vals = append(vals, inner)
	}
	return vals[len(vals)-1]
}

// TODO: make this work
// Runs in: O(n), where n=len(nums).
func countArrangements2(nums []int) int {
	var outerArr, innerArr, segmentLen = 1, 0, 1
	var onesMode = false
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] + 3 {
			if segmentLen > 2 {
				innerArr += 2 * (segmentLen-2)
			} else if segmentLen == 2 || !onesMode {
				innerArr++
			}
			outerArr *= innerArr
			segmentLen, innerArr, onesMode = 1, 0, false
		} else if nums[i] == nums[i-1] + 2 {
			if segmentLen > 2 {
				innerArr += 2 * (segmentLen-2)
			} else if segmentLen == 2 || !onesMode {
				innerArr++
			}
			onesMode = (segmentLen == 1)
			innerArr++ // +1 for the 2-width gap
			segmentLen = 1
		} else {
			segmentLen++
		}
	}
	return outerArr
}

// 0,1,2,4,5,7,8,9,12
// [1]
// [1,1]
