package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var wrappingPaper int
	var ribbon int

	// read the instructions, keeping track of current floor
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		measures := strings.Split(scanner.Text(), "x")
		l, _ := strconv.Atoi(measures[0])
		w, _ := strconv.Atoi(measures[1])
		h, _ := strconv.Atoi(measures[2])
		wrappingPaper += 2*l*w + 2*l*h + 2*w*h
		wrappingPaper += min(l*w, l*h, w*h)
		ribbon += min(2*l + 2*w, 2*l + 2*h, 2*w + 2*h);
		ribbon += l*w*h;
	}

	fmt.Printf("Wrapping paper needed: %v\n", wrappingPaper)
	fmt.Printf("Ribbon need in total: %v\n", ribbon)
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}
