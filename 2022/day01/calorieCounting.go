package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var elves [][]int
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			elves = append(elves, nums)
			nums = []int{}
			continue
		}
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	elves = append(elves, nums)

	fmt.Println("Most calories carried by one elf:", topSum(elves, 1))
	fmt.Println("Calories carried by top three elves:", topSum(elves, 3))
}

// Sums the numbers in each slice within `nums`.
// Outputs the sum of the largest `topN` such sums.
func topSum(nums [][]int, topN int) (res int) {
	var sums []int
	for _, nn := range nums {
		var sum int = 0
		for _, n := range nn {
			sum += n
		}
		sums = append(sums, sum)
	}
	sort.Ints(sums)
	for _, sum := range sums[len(sums)-topN:] {
		res += sum
	}
	return
}
