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

	// play recursive combat
	_, recScore := recursiveCombat(deck1, deck2, len(deck1), len(deck2))

	// play normal combat
	var winner int
	for len(deck1) > 0 && len(deck2) > 0 {
		card1, card2 := deck1[0], deck2[0]
		deck1, deck2 = deck1[1:], deck2[1:]
		if card1 > card2 { // player 1 wins
			deck1 = append(deck1, card1, card2)
			winner = 1
		} else { // player 2 wins
			deck2 = append(deck2, card2, card1)
			winner = 2
		}
	}

	fmt.Printf("%v\n", determineWinnerScore(winner, deck1, deck2))
	fmt.Printf("%v\n", recScore)
}

func recursiveCombat(deckPlayer1 []int, deckPlayer2 []int, cards1 int, cards2 int) (winner int, score int) {
	var alreadyPlayed = make(map[string]bool, 0)

	// copy decks
	deck1, deck2 := make([]int, cards1), make([]int, cards2)
	copy(deck1, deckPlayer1[:cards1])
	copy(deck2, deckPlayer2[:cards2])

	for len(deck1) > 0 && len(deck2) > 0 {
		winner = 1
		if alreadyPlayed[fmt.Sprintf("%v", deck1)] {
			break
		}
		alreadyPlayed[fmt.Sprintf("%v", deck1)] = true
		card1, card2 := deck1[0], deck2[0]
		deck1, deck2 = deck1[1:], deck2[1:]
		if card1 <= len(deck1) && card2 <= len(deck2) {
			winner, _ = recursiveCombat(deck1, deck2, card1, card2)
		} else if card2 > card1 {
			winner = 2
		}
		// give cards to winner
		if winner == 1 {
			deck1 = append(deck1, card1, card2)
		} else {
			deck2 = append(deck2, card2, card1)
		}
	}

	score = determineWinnerScore(winner, deck1, deck2)
	return
}

func determineWinnerScore(winner int, deck1 []int, deck2 []int) (score int) {
	var winnerDeck = deck1
	if winner == 2 {
		winnerDeck = deck2
	}

	for i, card := range winnerDeck {
		score += (len(winnerDeck) - i) * card
	}
	return
}
