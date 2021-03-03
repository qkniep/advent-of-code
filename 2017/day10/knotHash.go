package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var intList, byteList = createList(256), createList(256)
	var intLengths, byteLengths = make([]int, 0), make([]int, 0)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// interpret input as comma separated numbers
	for _, n := range strings.Split(input, ",") {
		i, _ := strconv.Atoi(n)
		intLengths = append(intLengths, i)
	}
	// interpret input as ASCII byte values and append suffix
	for _, b := range strings.TrimSuffix(input, "\n") {
		byteLengths = append(byteLengths, int(byte(b)))
	}
	byteLengths = append(byteLengths, 17, 31, 73, 47, 23)

	hash(intList, intLengths, 1)
	hash(byteList, byteLengths, 64)
	denseHash := densify(byteList)

	fmt.Printf("Product of first two values: %v\n", intList[0]*intList[1])
	fmt.Printf("Dense hash hex string: %s\n", denseHash)
}

// Initialize a new list of given size where each element is equal to its index.
func createList(size int) (list []int) {
	list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}
	return
}

// Perform given number of iterations of the knot-tying hash algorithm, based on lengths.
// Results are written into the list slice.
func hash(list []int, lengths []int, iterations int) {
	currentPos := 0
	skipSize := 0
	for i := 0; i < iterations; i++ {
		for _, l := range lengths {
			a, b := currentPos, currentPos+l-1
			for a < b {
				tmp := list[a%len(list)]
				list[a%len(list)] = list[b%len(list)]
				list[b%len(list)] = tmp
				a, b = a+1, b-1
			}
			currentPos += l + skipSize
			skipSize++
		}
	}
}

// Generates the dense hash representation as hex string.
func densify(sparse []int) (dense string) {
	for block := 0; block < 16; block++ {
		char := 0
		for i := 0; i < 16; i++ {
			char ^= sparse[16*block+i]
		}
		dense += fmt.Sprintf("%02x", char)
	}
	return
}
