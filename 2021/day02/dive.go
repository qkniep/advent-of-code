package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// y1 corresponds to part 1, y2 to part 2, x behaves the same for both
	var x, y1, y2, aim int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		d, _ := strconv.Atoi(s[1])
		switch s[0] {
		case "forward":
			x += d
			y2 += d * aim
		case "down":
			y1 += d
			aim += d
		case "up":
			y1 -= d
			aim -= d
		}
	}

	fmt.Println("Product of x and y without aim:", x*y1)
	fmt.Println("Product of x and y with aim:", x*y2)
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
