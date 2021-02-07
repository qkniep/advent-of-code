package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()
	valueRange := strings.Split(line, "-")

	low, _ := strconv.Atoi(valueRange[0])
	high, _ := strconv.Atoi(valueRange[1])

	fmt.Println("Number of conforming passwords in range:", countConforming(low, high, false))
	fmt.Println("Without counting larger groups as doubles:", countConforming(low, high, true))
}

// Count the number of conforming passwords within the range, borders inclusive.
func countConforming(low, high int, exactlyTwo bool) (count int) {
	for i := low; i <= high; i++ {
		if doesConform(i, exactlyTwo) {
			count++
		}
	}
	return
}

// Checks whether the password conforms to the following rules:
//  - All digits are monotonically non-decreasing.
//  - Contains a double: Two equal digits next to each other.
// If the parameter needExactlyTwo is set, 3 or more digits will not count as a double.
func doesConform(n int, needExactlyTwo bool) bool {
	var monotonic, doubleDigit, exactlyTwo = true, false, false
	var lastDigit, runLength = -1, 0

	for ; n > 0; n /= 10 {
		digit := n % 10
		if digit == lastDigit {
			doubleDigit = true
			runLength++
		} else {
			if runLength == 2 {
				exactlyTwo = true
			}
			runLength = 1
			if lastDigit != -1 && digit > lastDigit {
				monotonic = false
			}
		}
		lastDigit = digit
	}

	if !needExactlyTwo {
		return monotonic && doubleDigit
	}
	return monotonic && (exactlyTwo || runLength == 2)
}
