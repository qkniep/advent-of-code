package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type updateFn func(*[]string, []string, []string, *int64) bool

func main() {
	var nums []string

	// read binary numbers and store as strings
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums = append(nums, scanner.Text())
	}

	// calculate gamma and epsilon rate
	gamma := master(nums, updateGamma)
	epsilon := master(nums, updateEpsilon)

	// calculate oxygen and co2 rating
	oxygen := master(nums, updateOxygen)
	co2 := master(nums, updateCO2)

	fmt.Println("Product of gamma and epsilon rate:", gamma*epsilon)
	fmt.Println("Product of oxygen and CO2 rating:", oxygen*co2)
}

func updateGamma(nums *[]string, ones, zeroes []string, res *int64) bool {
	*res *= 2
	if len(ones) > len(zeroes) {
		*res++
	}
	return false
}

func updateEpsilon(nums *[]string, ones, zeroes []string, res *int64) bool {
	return updateGamma(nums, zeroes, ones, res)
}

func updateOxygen(nums *[]string, ones, zeroes []string, res *int64) bool {
	if len(ones) >= len(zeroes) {
		*nums = ones
	} else {
		*nums = zeroes
	}
	*res, _ = strconv.ParseInt((*nums)[0], 2, 64)
	return len(*nums) == 1
}

func updateCO2(nums *[]string, ones, zeroes []string, res *int64) bool {
	if len(zeroes) <= len(ones) {
		*nums = zeroes
	} else {
		*nums = ones
	}
	*res, _ = strconv.ParseInt((*nums)[0], 2, 64)
	return len(*nums) == 1
}

func master(nums []string, update updateFn) (res int64) {
	for pos, done := 0, false; pos < len(nums[0]) && !done; pos++ {
		var ones, zeroes []string
		for _, n := range nums {
			if n[pos] == '0' {
				zeroes = append(zeroes, n)
			} else {
				ones = append(ones, n)
			}
		}
		done = update(&nums, ones, zeroes, &res)
	}
	return
}
