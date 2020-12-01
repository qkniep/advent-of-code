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
		num, err := strconv.Atoi(scanner.Text())
		if err == nil {
			nums = append(nums, num)
		}
	}

	// search for 2 or 3 nums which sum to 2020
	for a := 0; a < len(nums); a++ {
		for b := a+1; b < len(nums); b++ {
			// check for a 2 number match
			if nums[a] + nums[b] == 2020 {
				fmt.Print("2 Entries: ")
				fmt.Printf("%d * %d = %d\n", nums[a], nums[b], nums[a] * nums[b])
			}

			// search for a 3rd number which makes this sum to 2020
			for c := b+1; c < len(nums); c++ {
				// check for a 3 number match
				if nums[a] + nums[b] + nums[c] == 2020 {
					fmt.Print("3 Entries: ")
					product := nums[a] * nums[b] * nums[c]
					fmt.Printf("%d * %d * %d = %d\n", nums[a], nums[b], nums[c], product)
				}
			}
		}
	}
}
