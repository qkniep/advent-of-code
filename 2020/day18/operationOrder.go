package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum = 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		expr := strings.Replace(scanner.Text(), "(", "( ", -1)
		expr = strings.Replace(expr, ")", " )", -1)
		sum += evaluateExpression(strings.Fields(expr))
	}

	fmt.Println(sum)
}

func evaluateExpression(tokens []string) int {
	i := 0
	result := evaluateParens(tokens, &i)
	for i < len(tokens) {
		op := tokens[i]
		i++
		n := evaluateParens(tokens, &i)
		if op == "+" {
			result += n
		} else if op == "*" {
			result *= n
		}
	}
	return result
}

func evaluateParens(tokens []string, pos *int) int {
	start := *pos
	parens := 0
	if tokens[start] != "(" {
		n, _ := strconv.Atoi(tokens[*pos])
		*pos++
		return n
	}
	for {
		*pos++
		if tokens[*pos-1] == "(" {
			parens++
		} else if tokens[*pos-1] == ")" {
			parens--
			if parens == 0 {
				return evaluateExpression(tokens[start+1:*pos-1])
			}
		}
	}
}
