package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var memory1, memory2 = make(map[int]int, 0), make(map[int]int, 0)
	var mask string
	var addr, value int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "mask = %s", &mask)
		if err != nil {
			fmt.Sscanf(scanner.Text(), "mem[%d] = %d", &addr, &value)
			memory1[addr] = applyBitmask(value, mask)
			addresses := findAddressesFromBitmask(&addr, mask)
			for _, a := range addresses {
				memory2[a] = value
			}
		}
	}

	fmt.Printf("Version 1 decoder chip: %v\n", sumOfValues(memory1))
	fmt.Printf("Version 2 decoder chip: %v\n", sumOfValues(memory2))
}

func applyBitmask(value int, mask string) int {
	for i, b := range mask {
		singleBitMask := 1 << (35 - i)
		if b == '0' {
			value &= ^singleBitMask
		} else if b == '1' {
			value |= singleBitMask
		}
	}
	return value
}

func findAddressesFromBitmask(addr *int, mask string) []int {
	var floating []int
	for i, b := range mask {
		singleBitMask := 1 << (35 - i)
		if b == '1' {
			*addr |= singleBitMask
		} else if b == 'X' {
			floating = append(floating, i)
		}
	}

	var addresses = []int{*addr}
	for _, bitPos := range floating {
		singleBitMask := 1 << (35 - bitPos)
		numOldAddr := len(addresses)
		for _, a := range addresses {
			addresses = append(addresses, a|singleBitMask)
			addresses = append(addresses, a&(^singleBitMask))
		}
		addresses = addresses[numOldAddr:]
	}

	return addresses
}

func sumOfValues(memory map[int]int) (sum int) {
	for _, v := range memory {
		sum += v
	}
	return
}
