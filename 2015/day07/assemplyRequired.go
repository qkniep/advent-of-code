package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var signals = make(map[string]uint16, 0)
	var gates = make(map[string][]string, 0)
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// read one instruction
		parsed := strings.Fields(scanner.Text())

		if len(parsed) == 3 {
			gates[parsed[2]] = parsed[:1]
		} else if len(parsed) == 4 { // NOT
			gates[parsed[3]] = parsed[:2]
		} else {
			gates[parsed[4]] = parsed[:3]
		}
	}

	var a1 = signal("a", gates, signals)
	signals = make(map[string]uint16, 0)
	signals["b"] = a1
	var a2 = signal("a", gates, signals)

	fmt.Println("Signal on wire 'a':", a1)
	fmt.Println("After changing 'b':", a2)
}

func signal(wire string, gates map[string][]string, signals map[string]uint16) uint16 {
	if s, err := strconv.Atoi(wire); err == nil {
		return uint16(s)
	} else if s, ok := signals[wire]; ok {
		return s
	}

	var gate = gates[wire]
	if len(gate) == 1 {
		signals[wire] = signal(gate[0], gates, signals)
	} else if len(gate) == 2 {
		signals[wire] = ^signal(gate[1], gates, signals)
	} else {
		a := signal(gate[0], gates, signals)
		switch gate[1] {
		case "AND":
			signals[wire] = a & signal(gate[2], gates, signals)
		case "OR":
			signals[wire] = a | signal(gate[2], gates, signals)
		case "LSHIFT":
			steps, _ := strconv.Atoi(gate[2])
			signals[wire] = a << steps
		case "RSHIFT":
			steps, _ := strconv.Atoi(gate[2])
			signals[wire] = a >> steps
		}
	}
	return signals[wire]
}
