package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var containers []int
	var combinations, minContainers, minCombinations int
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		size, _ := strconv.Atoi(scanner.Text())
		containers = append(containers, size)
	}

	minContainers = 999999
	subsetSumComb(containers, 0, 0, 150, &combinations, &minContainers, &minCombinations)

	fmt.Println("Combinations with any number of containers:", combinations)
	fmt.Println("Combinations with smallest possible number:", minCombinations)
}

// Sets `count` to the number of subsets of `a` there are that sum to a total of `sum`.
// Sets `minNum` to the size of the smallest such subset.
// Sets `minCount` to the number of such smallest subsets.
func subsetSumComb(a []int, i, n, sum int, count, minNum, minCount *int) {
	if sum < 0 || i == len(a) {
		if sum == 0 {
			*count++
			if n < *minNum {
				*minNum = n
				*minCount = 1
			} else if n == *minNum {
				*minCount++
			}
		}
		return
	}

	subsetSumComb(a, i+1, n, sum, count, minNum, minCount)
	subsetSumComb(a, i+1, n+1, sum-a[i], count, minNum, minCount)
}
