package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var deck1, deck2 = make([]int, 0), make([]int, 0)

	// read lists of players' cards and create decks
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // ignore 'Player 1' line
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		card, _ := strconv.Atoi(scanner.Text())
		deck1 = append(deck1, card)
	}
	scanner.Scan() // ignore 'Player 2' line
	for scanner.Scan() {
		card, _ := strconv.Atoi(scanner.Text())
		deck2 = append(deck2, card)
	}
	fmt.Printf("%v -- %v\n", deck1, deck2)

	// play recursive combat
	_, recScore := recursiveCombat(deck1, deck2)

	// play normal combat
	for len(deck1) > 0 && len(deck2) > 0 {
		card1, card2 := deck1[0], deck2[0]
		deck1, deck2 = deck1[1:], deck2[1:]
		if card1 > card2 { // player 1 wins
			deck1 = append(deck1, card1, card2)
		} else { // player 2 wins
			deck2 = append(deck2, card2, card1)
		}
	}

	_, score := determineWinnerScore(deck1, deck2)
	fmt.Printf("%v\n", score)
	fmt.Printf("%v\n", recScore)
}

func recursiveCombat(deck1 []int, deck2 []int) (winner int, score int) {
	// copy decks
	for len(deck1) > 0 && len(deck2) > 0 {
		// TODO: player 1 wins if the decks were already played this game
		card1, card2 := deck1[0], deck2[0]
		deck1, deck2 = deck1[1:], deck2[1:]
		roundWinner := 1
		if card1 <= len(deck1) && card2 <= len(deck2) {
			roundWinner, _ = recursiveCombat(deck1, deck2)
		} else if card2 > card1 {
			roundWinner = 2
		}
		if roundWinner == 1 { // player 1 wins
			deck1 = append(deck1, card1, card2)
		} else { // player 2 wins
			deck2 = append(deck2, card2, card1)
		}
	}
	return determineWinnerScore(deck1, deck2)
}

func determineWinnerScore(deck1 []int, deck2 []int) (winner int, score int) {
	var winnerDeck = deck1
	if len(deck1) == 0 {
		winnerDeck = deck2
	}

	for i, card := range winnerDeck {
		score += (len(winnerDeck) - i) * card
	}
	return
}
