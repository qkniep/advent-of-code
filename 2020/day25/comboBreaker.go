package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	publicKeyCard, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	publicKeyDoor, _ := strconv.Atoi(scanner.Text())

	//loopSizeCard := guessLoopSize(publicKeyCard)
	loopSizeDoor := guessLoopSize(publicKeyDoor)

	fmt.Printf("Final encryption key: %v\n", transformSubjectNumber(publicKeyCard, loopSizeDoor))
}

func guessLoopSize(publicKey int) int {
	var value, loopSize = 1, 0
	for value != publicKey {
		value = (value * 7) % 20201227
		loopSize++
	}
	return loopSize
}

func transformSubjectNumber(subjectNum int, loopSize int) int {
	var value = 1
	for i := 0; i < loopSize; i++ {
		value = (value * subjectNum) % 20201227
	}
	return value
}
