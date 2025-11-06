package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operator struct {
	prec   int
	rAssoc bool
}

var ops1 = map[string]operator{
	"*": {1, false},
	"+": {1, false},
}

var ops2 = map[string]operator{
	"*": {1, false},
	"+": {2, false},
}

func main() {
	var sum1, sum2 = 0, 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		expr := strings.Replace(scanner.Text(), "(", "( ", -1)
		expr = strings.Replace(expr, ")", " )", -1)
		sum1 += evaluateRPN(shuntingYard(strings.Fields(expr), ops1))
		sum2 += evaluateRPN(shuntingYard(strings.Fields(expr), ops2))
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}

// Turn infix notation into postfix (reverse polish) notation.
// Source: https://rosettacode.org/wiki/Parsing/Shunting-yard_algorithm#Go
func shuntingYard(tokens []string, ops map[string]operator) (output []string) {
	var stack []string
	for _, t := range tokens {
		switch t {
		case "(":
			stack = append(stack, t)
		case ")":
			var op string
			for {
				// pop item ("(" or operator) from stack
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op == "(" {
					break // discard "("
				}
				output = append(output, op)
			}
		default:
			if o1, isOp := ops[t]; isOp { // token is an operator
				for len(stack) > 0 {
					// consider top item on stack
					op := stack[len(stack)-1]
					if o2, isOp := ops[op]; !isOp || o1.prec > o2.prec ||
						o1.prec == o2.prec && o1.rAssoc {
						break
					}
					// top item is an operator that needs to come off
					stack = stack[:len(stack)-1] // pop it
					output = append(output, op)
				}
				// push operator (the new one) to stack
				stack = append(stack, t)
			} else { // token is an operand
				output = append(output, t)
			}
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		output = append(output, stack[i])
	}
	return
}

// Evaluate a calculation
func evaluateRPN(tokens []string) int {
	var stack []int
	for _, t := range tokens {
		l := len(stack)
		switch t {
		case "+":
			stack[l-2] += stack[l-1]
			stack = stack[:l-1]
		case "*":
			stack[l-2] *= stack[l-1]
			stack = stack[:l-1]
		default:
			n, _ := strconv.Atoi(t)
			stack = append(stack, n)
		}
	}
	return stack[0]
}
