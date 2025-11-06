package main

import "fmt"

const bufferSize = 2018
const angryBufferSize = 50_000_000

func main() {
	var stepSize int

	fmt.Scanf("%d", &stepSize)

	fmt.Println("Value after the position where 2017 was inserted:", valAfter(bufferSize, stepSize))
	fmt.Println("Value after 0, after inserting 50,000,000 values:", valAfter0(angryBufferSize, stepSize))
}

func valAfter(bufSize int, stepSize int) int {
	var buffer = []int{0}
	var pos int
	for i := 1; i < bufSize; i++ {
		pos = (pos + stepSize)%i + 1
		if pos == len(buffer) {
			buffer = append(buffer, i)
		} else {
			buffer = append(buffer, 0)
			copy(buffer[pos+1:], buffer[pos:])
			buffer[pos] = i
		}
	}
	return buffer[(pos+1)%len(buffer)]
}

func valAfter0(bufSize int, stepSize int) int {
	var pos, val int
	for i := 1; i < bufSize; i++ {
		pos = (pos + stepSize)%i + 1
		if pos == 1 {
			val = i
		}
	}
	return val
}
