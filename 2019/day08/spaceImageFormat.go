package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	width  = 25
	height = 6
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)

	scanner.Scan()
	data := scanner.Text()

	fmt.Println("The image checksum is:", imageChecksum(data))
	fmt.Println("The decoded image is:")
	drawImage(decodeImage(data))
}

func imageChecksum(data string) int {
	var minDigits = []int{(width * height) + 1, 0, 0}

	digits := []int{0, 0, 0}
	for i, r := range data {
		if int(r-'0') < 3 {
			digits[int(r-'0')]++
		}

		if (i+1)%(width*height) == 0 {
			if digits[0] < minDigits[0] {
				copy(minDigits, digits)
			}
			digits = []int{0, 0, 0}
		}
	}

	return minDigits[1] * minDigits[2]
}

func decodeImage(data string) (finalImage [][]int) {
	finalImage = make([][]int, height)
	for y := 0; y < height; y++ {
		finalImage[y] = make([]int, width)
		for x := 0; x < width; x++ {
			finalImage[y][x] = -1
		}
	}

	y := 0
	x := 0
	for i, r := range data {
		if finalImage[y][x] < 0 && int(r-'0') != 2 {
			finalImage[y][x] = int(r - '0')
		}

		if (i+1)%(width*height) == 0 {
			x = 0
			y = 0
		} else if (i+1)%width == 0 {
			x = 0
			y++
		} else {
			x++
		}
	}
	return
}

func drawImage(image [][]int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if image[y][x] == 0 {
				fmt.Printf("\u2591")
			} else if image[y][x] == 1 {
				fmt.Printf("\u2588")
			}
		}
		fmt.Println()
	}
}
