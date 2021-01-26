package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var weights []int
	var weightSum int
	var scanner = bufio.NewScanner(os.Stdin)

	// read sorted list of package weights, maintaining a running sum
	for scanner.Scan() {
		weight, _ := strconv.Atoi(scanner.Text())
		weights = append(weights, weight)
		weightSum += weight
	}

	fmt.Println("Minimal quantum entanglement for 3 groups:", findMinimalQE(weights, weightSum/3))
	fmt.Println("Minimal quantum entanglement for 4 groups:", findMinimalQE(weights, weightSum/4))
}

// TODO Drop the assumption about it just being the smallest QE possible for the group1 sum,
// i.e. consider the case that the remaining elements can't be split into two groups of equal weight.
func findMinimalQE(weights []int, targetSum int) int {
	var groupSize = smallestSubsetSum(weights, targetSum)
	var cs = combinations(weights, groupSize)
	var qes []int
	var result = make(map[int][][]int, 0)

	for _, c := range cs {
		// drop unqualifying answers
		var sum int
		for _, e := range c {
			sum += e
		}
		if sum != targetSum {
			continue
		}

		qe := quantumEntanglement(c)
		if len(result[qe]) == 0 {
			qes = append(qes, qe)
		}
		result[qe] = append(result[qe], c)
	}

	sort.Ints(qes)
	return qes[0]
}

// source: https://prtamil.github.io/posts/powersets-go/
func combinations(L []int, r int) [][]int {
	if r == 1 {
		temp := make([][]int, 0)
		for _, rr := range L {
			t := make([]int, 0)
			t = append(t, rr)
			temp = append(temp, [][]int{t}...)
		}
		return temp
	}
	res := make([][]int, 0)
	for i := 0; i < len(L); i++ {
		perms := make([]int, 0)
		perms = append(perms, L[:i]...)
		for _, x := range combinations(perms, r-1) {
			t := append(x, L[i])
			res = append(res, [][]int{t}...)
		}
	}
	return res
}

// Returns the smallest size of a subset in `set` that sums to `sum`.
// If there is no such subset this returns -1 instead.
func smallestSubsetSum(set []int, sum int) int {
	for size := 1; size <= len(set); size++ {
		if subsetSumExists(set, sum, size) {
			return size
		}
	}
	return -1
}

// Checks whether a subset of the given size exists that sums to the given sum.
// TODO speed up through memoization?
func subsetSumExists(set []int, sum int, subsetSize int) bool {
	if subsetSize == 0 {
		return sum == 0
	}
	for i := range set {
		var newSet = make([]int, len(set)-1)
		copy(newSet[:i], set[:i])
		copy(newSet[i:], set[i+1:])
		if subsetSumExists(newSet, sum-set[i], subsetSize-1) {
			return true
		}
	}
	return false
}

// Returns the quantum entanglement of a group of packages, i.e. the product of their weights.
func quantumEntanglement(group []int) int {
	var product = 1
	for _, item := range group {
		product *= item
	}
	return product
}
