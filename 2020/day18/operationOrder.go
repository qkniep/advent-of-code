package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum1, sum2 = 0, 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		expr := strings.Replace(scanner.Text(), "(", "( ", -1)
		expr = strings.Replace(expr, ")", " )", -1)
		sum1 += evaluateExpression(strings.Fields(expr))
		sum2 += evaluateExpression2(strings.Fields(expr))
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
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

func evaluateExpression2(tokens []string) int {
	i := 0
	result := evaluateParensOrAdd(tokens, &i)
	for i < len(tokens) {
		op := tokens[i]
		i++
		n := evaluateParensOrAdd(tokens, &i)
		if op == "+" {
			result += n
		} else if op == "*" {
			result *= n
		}
	}
	fmt.Printf("%v -> %v\n", tokens, result)
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

func evaluateParensOrAdd(tokens []string, pos *int) int {
	start := *pos
	parens := 0
	if tokens[start] != "(" {
		n, _ := strconv.Atoi(tokens[*pos])
		if start+1 < len(tokens) && tokens[start+1] == "+" {
			*pos += 2
			return n + evaluateParensOrAdd(tokens, pos)
		}
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
				return evaluateExpression2(tokens[start+1:*pos-1])
			}
		}
	}
}
