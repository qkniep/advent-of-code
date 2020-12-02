package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var numValid1 int = 0
	var numValid2 int = 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// parse one line of input
		fields := strings.Fields(scanner.Text())
		var n1, n2 int
		fmt.Sscanf(fields[0], "%d-%d", &n1, &n2)
		char := fields[1][0]

		// check whether password conforms to the two password policies
		if checkPasswordPolicy1(fields[2], n1, n2, char) {
			numValid1++
		}
		if checkPasswordPolicy2(fields[2], n1, n2, char) {
			numValid2++
		}
	}

	fmt.Printf("Valid according to Policy 1: %d\n", numValid1)
	fmt.Printf("Valid according to Policy 2: %d\n", numValid2)
}

// Checks if `password` contains `char` at least `min` time and at most `max` times.
func checkPasswordPolicy1(password string, min int, max int, char byte) bool {
	count := 0
	for i := 0; i < len(password); i++ {
		if password[i] == char {
			count++
		}
	}
	return count >= min && count <= max
}

// Checks if `password` contains `char` at exactly one of `firstPos` or `secondPos`,
// both positions starting from 1 instead of 0.
func checkPasswordPolicy2(password string, firstPos int, secondPos int, char byte) bool {
	return (password[firstPos-1] == char) != (password[secondPos-1] == char)
}
