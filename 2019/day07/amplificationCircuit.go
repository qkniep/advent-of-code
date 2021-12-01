package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	positionMode  = 0
	immediateMode = 1

	addOpCode = 1
	mulOpCode = 2
	inOpCode  = 3
	outOpCode = 4
)

func main() {
	var intcode []int
	var scanner = bufio.NewScanner(os.Stdin)
	var testCase = 0

	for scanner.Scan() {
		if testCase > 0 {
			fmt.Println()
		}

		line := scanner.Text()
		code := strings.Split(line, ",")

		// convert to integers
		intcode = make([]int, len(code))
		for i, s := range code {
			intcode[i], _ = strconv.Atoi(s)
		}

		fmt.Println("Max thruster signal:", findMaxThrust(intcode))
		//fmt.Println("Input 5 leads to diagnostic code:", runTestProgram(intcode, 5))
		testCase++
	}
}

// Tries all permutations of the phase settings and return the highest thruster output achieved.
func findMaxThrust(intcode []int) (max int) {
	for _, phaseSettings := range permutations([]int{0, 1, 2, 3, 4}) {
		thrust := runChain(intcode, phaseSettings)
		if thrust > max {
			max = thrust
		}
	}
	return
}

// Runs the whole chain of amplifiers and returns the output signal going to the thrusters.
func runChain(intcode []int, phaseSettings []int) (output int) {
	for _, phaseSetting := range phaseSettings {
		output = runAmplifierProgram(intcode, phaseSetting, output)
	}
	return
}

// Loads a copy of the program into memory (copies slice).
func runAmplifierProgram(intcode []int, phaseSetting, prevOutput int) int {
	memory := make(map[int]int, len(intcode))
	for i, instruction := range intcode {
		memory[i] = instruction
	}

	inputs := []int{phaseSetting, prevOutput}
	output := -1

	for ip := 0; memory[ip] != 99; {
		jumped := false

		if output > 0 {
			fmt.Println("[ERROR] Output of test was:", output)
		}

		// parse op code and parameters
		op, modes := parseOpCode(memory[ip])
		params := make([]int, len(modes))
		for i, mode := range modes {
			params[i] = parseParam(memory, ip+i+1, mode)
		}

		// perform the operation
		switch op {
		case 1:
			memory[memory[ip+3]] = params[0] + params[1]
		case 2:
			memory[memory[ip+3]] = params[0] * params[1]
		case 3:
			memory[memory[ip+1]] = inputs[0]
			inputs = inputs[1:]
		case 4:
			output = params[0]
		case 5:
			if params[0] != 0 {
				ip = params[1]
				jumped = true
			}
		case 6:
			if params[0] == 0 {
				ip = params[1]
				jumped = true
			}
		case 7:
			if params[0] < params[1] {
				memory[memory[ip+3]] = 1
			} else {
				memory[memory[ip+3]] = 0
			}
		case 8:
			if params[0] == params[1] {
				memory[memory[ip+3]] = 1
			} else {
				memory[memory[ip+3]] = 0
			}
		default:
			fmt.Println("[ERROR] Unsupported op code:", op)
		}

		if !jumped {
			ip += 1 + len(params)
		}
	}

	return output
}

// Parses the full instruction with up to 5 digits into op code and params.
// ABCDE = A (3rd param) + B (2nd param) + C (1st param) + DE (op code)
// The three parameters are optional and are zero if ommitted.
func parseOpCode(intcode int) (opcode int, paramModes []int) {
	opcode = intcode % 100

	params := 3
	if opcode == 3 || opcode == 4 {
		params = 1
	} else if opcode == 5 || opcode == 6 {
		params = 2
	}

	paramModes = make([]int, params)
	val := intcode / 100
	for p := 0; p < params; p++ {
		paramModes[p] = val % 10
		val /= 10
	}

	return
}

// Parses the parameter based on its mode (0: position, 1: immediate).
func parseParam(memory map[int]int, memPos, mode int) int {
	switch mode {
	case 1:
		return memory[memPos]
	default:
		return memory[memory[memPos]]
	}
}

// source: https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
