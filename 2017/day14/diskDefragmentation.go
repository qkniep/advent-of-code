package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	keyString := scanner.Text()

	squaresUsed, connectedRegions := analyzeUsedSpace(keyString)

	fmt.Printf("Squares used on disk: %v\n", squaresUsed)
	fmt.Printf("Connected regions: %v\n", connectedRegions)
}

func analyzeUsedSpace(key string) (squaresUsed int, regions int) {
	for i := 0; i < 128; i++ {
		knotHash := createList(256)
		input := []byte(fmt.Sprintf("%s-%d", key, i))
		input = append(input, 17, 31, 73, 47, 23)
		hash(knotHash, input, 64)
		squaresUsed += denseHashOneBits(knotHash)
	}
	return
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
func hash(list []int, lengths []byte, iterations int) {
	currentPos := 0
	skipSize := 0
	for i := 0; i < iterations; i++ {
		for _, l := range lengths {
			a, b := currentPos, currentPos+int(l)-1
			for a < b {
				tmp := list[a%len(list)]
				list[a%len(list)] = list[b%len(list)]
				list[b%len(list)] = tmp
				a, b = a+1, b-1
			}
			currentPos += int(l) + skipSize
			skipSize++
		}
	}
}

// Generates the dense hash representation as hex string.
func denseHashOneBits(sparse []int) (oneBits int) {
	for block := 0; block < 16; block++ {
		char := 0
		for i := 0; i < 16; i++ {
			char ^= sparse[16*block+i]
		}
		oneBits += bits.OnesCount(uint(char))
	}
	return
}
