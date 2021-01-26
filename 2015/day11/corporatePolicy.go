package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nextPassword := findNextPassword(scanner.Text())
	secondNextPassword := findNextPassword(nextPassword)

	fmt.Println("Next password for Santa:", nextPassword)
	fmt.Println("Password after that:", secondNextPassword)
}

// Finds the next password starting from pwd, which satisfies the password policy.
func findNextPassword(pwd string) string {
	var isPolicyCompliant bool

	for !isPolicyCompliant {
		pwd = incrementPassword(pwd)
		isPolicyCompliant = satisfiesPolicy(pwd)
	}

	return pwd
}

// Checks whether a given password satisfies the Security-Elf's password policy.
func satisfiesPolicy(pwd string) bool {
	var lastChar, secondLastChar, pairChar rune
	var pairsFound, straightFound bool

	for _, r := range pwd {
		if r == 'i' || r == 'o' || r == 'l' {
			return false
		} else if lastChar == secondLastChar+1 && r == lastChar+1 {
			straightFound = true
		} else if r == lastChar {
			if pairChar == 0 {
				pairChar = r
			} else if pairChar != r {
				pairsFound = true
			}
		}
		secondLastChar = lastChar
		lastChar = r
	}

	return pairsFound && straightFound
}

// Increments the password string as if it were a base-26 number (with digits a-z).
func incrementPassword(pwd string) string {
	var needToIncrement = true
	var pwdSlice = []byte(pwd)

	for i := len(pwd) - 1; i >= 0 && needToIncrement; i-- {
		if pwdSlice[i] == byte('z') {
			pwdSlice[i] = byte('a')
		} else {
			pwdSlice[i]++
			needToIncrement = false
		}
	}

	return string(pwdSlice)
}
