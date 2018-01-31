package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	spreadsheet := readInput()
	fmt.Printf("Checksum: %v\n", calculateChecksum(spreadsheet))
	fmt.Printf("Sum of Quotients: %v\n", calculateQuotientSum(spreadsheet))
}

func readInput() (ss [][]int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowOfNums := make([]int, 0)
		fmt.Println(scanner.Text())
		for _, s := range strings.Fields(scanner.Text()) {
			i, _ := strconv.Atoi(s)
			rowOfNums = append(rowOfNums, i)
		}
		ss = append(ss, rowOfNums)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func calculateChecksum(ss [][]int) (checksum int) {
	for _, row := range ss {
		min, max := row[0], row[0]
		for _, cell := range row {
			if cell < min {
				min = cell
			} else if cell > max {
				max = cell
			}
		}
		checksum += max - min
	}
	return
}

func calculateQuotientSum(ss [][]int) (sum int) {
	for _, row := range ss {
		for _, cell := range row {
			for _, div := range row {
				quotient := cell / div
				if div != cell && (quotient*div) == cell {
					sum += quotient
				}
			}
		}
	}
	return
}
