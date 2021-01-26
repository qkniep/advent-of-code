package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var program [][]string
	var scanner = bufio.NewScanner(os.Stdin)

	// read program instructions
	for scanner.Scan() {
		program = append(program, strings.Fields(scanner.Text()))
	}

	var bStartA0 = emulateProgramBehaviour(program, 0, 0, 0)
	var bStartA1 = emulateProgramBehaviour(program, 0, 1, 0)

	fmt.Println("Final value of register b:", bStartA0)
	fmt.Println("Final value of register b (start with a=1):", bStartA1)
}

func emulateProgramBehaviour(program [][]string, ip, ra, rb int) int {
	for ip < len(program) {
		switch program[ip][0] {
		case "hlf":
			ra /= 2
			ip++
		case "tpl":
			ra *= 3
			ip++
		case "inc":
			if program[ip][1] == "a" {
				ra++
			} else if program[ip][1] == "b" {
				rb++
			}
			ip++
		case "jmp":
			offset, _ := strconv.Atoi(program[ip][1])
			ip += offset
		case "jie":
			if ra%2 == 0 {
				offset, _ := strconv.Atoi(program[ip][2])
				ip += offset
			} else {
				ip++
			}
		case "jio":
			if ra == 1 {
				offset, _ := strconv.Atoi(program[ip][2])
				ip += offset
			} else {
				ip++
			}
		}
	}
	return rb
}
