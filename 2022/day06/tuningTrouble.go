package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var receivedSignals string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	receivedSignals = scanner.Text()

	firstPacketMarker := findFirstMarker(receivedSignals, 4)
	firstMessageMarker := findFirstMarker(receivedSignals, 14)

	fmt.Println("End position of first packet marker:", firstPacketMarker)
	fmt.Println("End position of first message marker:", firstMessageMarker)
}

// Finds the first marker of a given length, i.e. substring of different characters.
// Returns the end position of the marker.
func findFirstMarker(signal string, length int) int {
	for start := 0; start < len(signal)-4; start++ {
		var isPacketStart = true
		for i := 0; i < length && isPacketStart; i++ {
			for j := 0; j < length && isPacketStart; j++ {
				if i == j {
					continue
				}
				if signal[start+i] == signal[start+j] {
					isPacketStart = false
				}
			}
		}
		if isPacketStart {
			return start + length
		}
	}
	return -1
}
