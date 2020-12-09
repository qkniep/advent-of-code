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

	// read program instructions
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var words = strings.Fields(scanner.Text())
		num, _ := strconv.Atoi(words[1])
		program = append(program, line{ cmd: words[0], arg: num})
	}

	acc, _ := emulate(program)
	fixedAcc := fixProgram(program)

	fmt.Printf("Acc value on loop: %d\n", acc)
	fmt.Printf("Fixed acc value: %d\n", fixedAcc)
}

// Emulates the program, keeping track of the accumulator value.
// Returns the current accumulator value and a bool, which is true iff the emulation ended normally,
// i.e. we tried to run the line immediately after the last line in the program.
// The other case (returned bool is false) happens when we encounter an infinite loop.
func emulate(program []line) (int, bool) {
	var executed = make([]bool, len(program))
	var instruction, accumulator = 0, 0

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

// Try fixing the program by randomly changing nop/jmp and emulating the changed program.
// Returns the accumulator value of the first successful run.
// Runs O(n+j) full emulations of the program, where n/j are the number of nop/jmp instructions.
func fixProgram(program []line) int {
	for i := 0; i < len(program); i++ {
		if !swapInstructions(&program[i], "nop", "jmp") {
			continue
		}
		accumulator, terminatedNormally := emulate(program)
		if terminatedNormally {
			return accumulator
		}
		swapInstructions(&program[i], "nop", "jmp")
	}
	return -1
}

// Checks whether source is instruction of type `instA` or `instB`, if so swaps it to the other.
// Returns true iff the instruction was swapped.
func swapInstructions(source *line, instA string, instB string) bool {
	if source.cmd == instA {
		source.cmd = instB
	} else if source.cmd == instB{
		source.cmd = instA
	} else {
		return false
	}
	return true
}
