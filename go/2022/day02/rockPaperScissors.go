package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var strategy []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strategy = append(strategy, scanner.Text())
	}

	var scoreMove = 0
	var scoreOutcome = 0
	for _, turn := range strategy {
		opponentMove := int(turn[0] - 'A')
		strategyAdvice := int(turn[2] - 'X')
		scoreMove += scoreRound(strategyAdvice, opponentMove)
		ownMove := moveForOutcome(opponentMove, strategyAdvice)
		scoreOutcome += scoreRound(ownMove, opponentMove)
	}

	fmt.Println("Score if interpreting second column as moves:", scoreMove)
	fmt.Println("Score if interpreting second column as outcomes:", scoreOutcome)
}

// Returns the move that you need to play, given the desired outcome and the opponent's move.
func moveForOutcome(opponentMove int, desiredOutcome int) (ownMove int) {
	if desiredOutcome == 0 {
		ownMove = (opponentMove + 2) % 3
	} else if desiredOutcome == 1 {
		ownMove = opponentMove
	} else if desiredOutcome == 2 {
		ownMove = (opponentMove + 1) % 3
	}
	return
}

// Returns the score of the game round, given the moves of the two players.
func scoreRound(ownMove int, opponentMove int) (res int) {
	res = ownMove + 1
	if ownMove == opponentMove {
		res += 3
	} else if ownMove == (opponentMove+1)%3 {
		res += 6
	}
	return
}
