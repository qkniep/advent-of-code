package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const desiredOutput = 19690720

func main() {
	var intcode []int
	var scanner = bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()
	code := strings.Split(line, ",")

	// convert to integers
	intcode = make([]int, len(code))
	for i, s := range code {
		intcode[i], _ = strconv.Atoi(s)
	}

	correctNoun, correctVerb := findCorrectInput(intcode)

	fmt.Println("Value at address 0 after restoring 1202 state:", runProgramCopy(intcode, 12, 2))
	fmt.Println("Correct input (first two digits noun, second two verb):", 100 * correctNoun + correctVerb)
}

// Finds inputs (noun+verb) for the given intcode prgram that produce the desired output.
func findCorrectInput(intcode []int) (int, int) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			if runProgramCopy(intcode, noun, verb) == desiredOutput {
				return noun, verb
			}
		}
	}
	return -1, -1
}

// Loads a copy of the program into memory (copies slice).
// Runs the intcode program with the given input noun and verb.
// Finally, returns the value that is in memory address 0.
func runProgramCopy(intcode []int, noun int, verb int) int {
	memory := make([]int, len(intcode))
	copy(memory, intcode)

	memory[1] = noun
	memory[2] = verb

	for ip := 0; memory[ip] != 99; ip += 4 {
		in1 := memory[ip+1]
		in2 := memory[ip+2]
		out := memory[ip+3]

		switch memory[ip] {
		case 1:
			memory[out] = memory[in1] + memory[in2]
		case 2:
			memory[out] = memory[in1] * memory[in2]
		}
	}

	return memory[0]
}
