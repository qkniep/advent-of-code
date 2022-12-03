package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var rucksacks []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
	}

	var rucksackPriority = 0
	for _, rucksack := range rucksacks {
		var compartment1 = []byte(rucksack[:len(rucksack)/2])
		var compartment2 = []byte(rucksack[len(rucksack)/2:])
		priority, _ := commonItems(compartment1, compartment2)
		rucksackPriority += priority
	}

	var groupPriority = 0
	for i := 0; i < len(rucksacks)/3; i++ {
		_, items := commonItems([]byte(rucksacks[3*i]), []byte(rucksacks[3*i+1]))
		priority, _ := commonItems(items, []byte(rucksacks[3*i+2]))
		groupPriority += priority
	}

	fmt.Println("Sum of priorities for each elf:", rucksackPriority)
	fmt.Println("Sum of priorities in group of 3 elves:", groupPriority)
}

func commonItems(items1 []byte, items2 []byte) (priority int, common []byte) {
	var countedItems = make(map[byte]bool)
	for _, item1 := range items1 {
		for _, item2 := range items2 {
			if item1 == item2 && !countedItems[item1] {
				priority += itemPriority(item1)
				common = append(common, item1)
				countedItems[item1] = true
			}
		}
	}
	return
}

func itemPriority(item byte) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	} else if item >= 'A' && item <= 'Z' {
		return int(item - 'A' + 27)
	}
	return 0
}
