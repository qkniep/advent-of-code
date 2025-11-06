package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type movement struct {
	count int
	from  int
	to    int
}

func main() {
	var stacks [][]rune
	var moves []movement

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		if row[1] == '1' {
			break
		}
		for i := 0; i < len(row); i += 4 {
			if len(stacks) <= i/4 {
				stacks = append(stacks, []rune{})
			}
			if row[i+1] != ' ' {
				stacks[i/4] = append([]rune{rune(row[i+1])}, stacks[i/4]...)
			}
		}
	}
	scanner.Scan() // skip empty line

	// parse and permform movements
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		count, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])

		from = from - 1
		to = to - 1

		moves = append(moves, movement{count, from, to})
	}

	stacksCopy := copyStacks(stacks)
	performIndividualMoves(stacksCopy, moves)
	topIndividual := getTopElements(stacksCopy)

	stacksCopy = copyStacks(stacks)
	performBatchMoves(stacksCopy, moves)
	topBatch := getTopElements(stacksCopy)

	fmt.Println("Top crates after individual moves:", topIndividual)
	fmt.Println("Top crates after batched moves:", topBatch)
}

// Moves crates between the stacks one-by-one.
// This modifies the stacks slice.
func performIndividualMoves(stacks [][]rune, moves []movement) {
	for _, move := range moves {
		for i := 0; i < move.count; i++ {
			fromTop := len(stacks[move.from]) - 1
			stacks[move.to] = append(stacks[move.to], stacks[move.from][fromTop])
			stacks[move.from] = stacks[move.from][:fromTop]
		}
	}
}

// Moves crates between the stacks in bacthes, retaining the original order.
// This modifies the stacks slice.
func performBatchMoves(stacks [][]rune, moves []movement) {
	for _, move := range moves {
		fromTop := len(stacks[move.from]) - move.count
		stacks[move.to] = append(stacks[move.to], stacks[move.from][fromTop:]...)
		stacks[move.from] = stacks[move.from][:fromTop]
	}
}

// Combine the top crates into a single string.
func getTopElements(stacks [][]rune) (topElements string) {
	for _, stack := range stacks {
		topElements += string(stack[len(stack)-1])
	}
	return
}

func copyStacks(stacks [][]rune) (copied [][]rune) {
	for _, stack := range stacks {
		copied = append(copied, append([]rune{}, stack...))
	}
	return
}
