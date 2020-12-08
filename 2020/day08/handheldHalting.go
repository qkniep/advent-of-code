package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type line struct {
	cmd string
	arg int
}

func main() {
	var program = make([]line, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var words = strings.Fields(scanner.Text())
		num, _ := strconv.Atoi(words[1])
		program = append(program, line{ cmd: words[0], arg: num})
	}

	acc, _ := emulate(program)
	fmt.Printf("Acc value: %d\n", acc)

	// try fixing by randomly changing nop/jmp
	acc, term := 0, false
	for i := 0; i < len(program); i++ {
		// switch
		if program[i].cmd == "nop" {
			program[i].cmd = "jmp"
		} else if program[i].cmd == "jmp" {
			program[i].cmd = "nop"
		}
		// emulate
		acc, term = emulate(program)
		if term {
			break
		}
		// switch back
		if program[i].cmd == "nop" {
			program[i].cmd = "jmp"
		} else if program[i].cmd == "jmp" {
			program[i].cmd = "nop"
		}
	}

	fmt.Printf("Fixed Acc value: %d\n", acc)
}

func emulate(program []line) (int, bool) {
	var executed = make([]bool, len(program))
	var instruction = 0
	var accumulator = 0

	for instruction < len(program) && !executed[instruction] {
		executed[instruction] = true
		if program[instruction].cmd == "jmp" {
			instruction += program[instruction].arg
		} else {
			if program[instruction].cmd == "acc" {
				accumulator += program[instruction].arg
			}
			instruction++
		}
	}

	return accumulator, instruction == len(program)
}
