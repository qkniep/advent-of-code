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
	registers, maxValueOccured := readInput()
	fmt.Printf("Highest (final) Value: %v\n", findHighestValue(registers))
	fmt.Printf("Highest Value Occured: %v\n", maxValueOccured)
}

func readInput() (registers map[string]int, maxValueOccured int) {
	registers = make(map[string]int)
	maxValueOccured = -1 << 31

	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		dest := fields[0]
		value, _ := strconv.Atoi(fields[2])
		if fields[1] == "dec" {
			value = -value
		}
		condition := false
		a := registers[fields[4]]
		b, _ := strconv.Atoi(fields[6])
		switch fields[5] {
		case "<":
			condition = a < b
		case ">":
			condition = a > b
		case "<=":
			condition = a <= b
		case ">=":
			condition = a >= b
		case "==":
			condition = a == b
		case "!=":
			condition = a != b
		}
		if condition {
			registers[dest] += value
			if registers[dest] > maxValueOccured {
				maxValueOccured = registers[dest]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func findHighestValue(registers map[string]int) (max int) {
	max = -1 << 31
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	return
}
