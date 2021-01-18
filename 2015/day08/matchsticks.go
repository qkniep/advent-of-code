package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var encodedOverhead, codeOverhead int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var escaped bool
		encodedOverhead += 2 // new double quotes
		for i, r := range scanner.Text() {
			if i == 0 || i == len(scanner.Text())-1 {
				encodedOverhead++
				codeOverhead++
			} else if escaped {
				if r == '\\' || r == '"' {
					encodedOverhead += 2
					codeOverhead++
				} else if r == 'x' {
					encodedOverhead++
					codeOverhead += 3
				}
				escaped = false
			} else if r == '\\' {
				escaped = true
			}
		}
	}

	fmt.Println("Overhead of string length in code:", codeOverhead)
	fmt.Println("Overhead of encoding code representation:", encodedOverhead)
}
