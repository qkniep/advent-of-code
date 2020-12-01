package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var nums []int
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for a := 0; a < len(nums); a++ {
		for b := a+1; b < len(nums); b++ {
			if nums[a] + nums[b] == 2020 {
				fmt.Print("2 Entries: ")
				fmt.Printf("%d * %d = %d\n", nums[a], nums[b], nums[a] * nums[b])
			}
		}
	}

	for a := 0; a < len(nums); a++ {
		for b := a+1; b < len(nums); b++ {
			for c := b+1; c < len(nums); c++ {
				if nums[a] + nums[b] + nums[c] == 2020 {
					fmt.Print("3 Entries: ")
					product := nums[a] * nums[b] * nums[c]
					fmt.Printf("%d * %d * %d = %d\n", nums[a], nums[b], nums[c], product)
				}
			}
		}
	}
}
