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

	for scanner.Scan() {
		program = append(program, strings.Fields(scanner.Text()))
	}

	fmt.Printf("Squares used on disk: %v\n", performDuet(program))
	//fmt.Printf("Connected regions: %v\n", connectedRegions)
}

func performDuet(program [][]string) int {
	var registers = make(map[string]int, 0)
	var ip, x, y, lastPlayed int
	for {
		x = literalOrRegister(program[ip][1], registers)
		if len(program[ip]) == 3 {
			y = literalOrRegister(program[ip][2], registers)
		}

		switch program[ip][0] {
		case "snd":
			lastPlayed = x
		case "set":
			registers[program[ip][1]] = y
		case "add":
			registers[program[ip][1]] += y
		case "mul":
			registers[program[ip][1]] *= y
		case "mod":
			registers[program[ip][1]] %= y
		case "rcv":
			if x != 0 {
				return lastPlayed
			}
		case "jgz":
			if x > 0 {
				ip += y-1
			}
		}
		ip++
	}
}

func literalOrRegister(v string, registers map[string]int) int {
	if n, err := strconv.Atoi(v); err == nil {
		return n
	}
	return registers[v]
}
