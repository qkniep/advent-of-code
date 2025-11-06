package main

import "fmt"

const (
	shield = iota
	poison
	recharge
)

type effect int

type spell struct {
	cost        int
	damage      int
	healing     int
	effectTimer int
	effectType  effect
}

type basicCharacter struct {
	hitPoints, damage int
}

type wizard struct {
	hitPoints, mana, armor int
}

type gameState struct {
	activeEffects      [3]int
	player             wizard
	boss               basicCharacter
	playerWon, bossWon bool
	totalManaSpent     int
}

var spells = []spell{
	{53, 4, 0, 0, 0},
	{73, 2, 2, 0, 0},
	{113, 0, 0, 6, shield},
	{173, 0, 0, 6, poison},
	{229, 0, 0, 5, recharge},
}

func main() {
	var state gameState
	state.player = wizard{50, 500, 0}

	// read boss' stats
	fmt.Scanf("Hit Points: %d", &state.boss.hitPoints)
	fmt.Scanf("Damage: %d", &state.boss.damage)

	minManaEasy := findCheapestWin(state, false)
	minManaHard := findCheapestWin(state, true)

	fmt.Printf("Minimum amount of mana to win on easy mode: %d\n", minManaEasy)
	fmt.Printf("Minimum amount of mana to win on hard mode: %d\n", minManaHard)
}

// Use Dijkstra's algorithm to find the cheapest winning strategy, in terms of mana spent.
// Returns the amount of mana needed.
func findCheapestWin(initialState gameState, hardMode bool) int {
	var manaToReachState = map[gameState]int{initialState: 0}
	var visited = make(map[gameState]bool, 0)

	for {
		// find cheapest reachable state
		var currentCheapestMana = 999999
		var currentCheapestState gameState
		for state, mana := range manaToReachState {
			if !visited[state] && mana < currentCheapestMana {
				currentCheapestMana = mana
				currentCheapestState = state
			}
		}

		visited[currentCheapestState] = true
		if currentCheapestState.bossWon {
			continue
		} else if currentCheapestState.playerWon {
			return currentCheapestMana
		}

		for _, spell := range possibleSpells(currentCheapestState) {
			nextState := simulateRound(currentCheapestState, spell, hardMode)
			newTotalMana := currentCheapestState.totalManaSpent + spell.cost
			if newTotalMana < manaToReachState[nextState] || manaToReachState[nextState] == 0 {
				manaToReachState[nextState] = newTotalMana
			}
		}
	}
}

// Simulates one round (one player turn + one enemy turn) of the game.
// Returns the updated game state.
func simulateRound(state gameState, action spell, hardMode bool) gameState {
	if hardMode {
		state.player.hitPoints--
		if state.player.hitPoints <= 0 {
			state.bossWon = true
			return state
		}
	}

	// player's turn
	state.player.mana -= action.cost
	state.totalManaSpent += action.cost
	state.boss.hitPoints -= action.damage
	state.player.hitPoints += action.healing
	state.activeEffects[action.effectType] = action.effectTimer
	if state.boss.hitPoints <= 0 {
		state.playerWon = true
		return state
	}

	// boss' turn
	state = applyEffects(state)
	state.player.hitPoints -= max(1, state.boss.damage-state.player.armor)
	if state.player.hitPoints <= 0 {
		state.bossWon = true
	}

	state = applyEffects(state)
	return state
}

func applyEffects(state gameState) gameState {
	state.player.armor = 0
	if state.activeEffects[shield] > 0 {
		state.player.armor = 7
		state.activeEffects[shield]--
	}
	if state.activeEffects[poison] > 0 {
		state.boss.hitPoints -= 3
		state.activeEffects[poison]--
	}
	if state.activeEffects[recharge] > 0 {
		state.player.mana += 101
		state.activeEffects[recharge]--
	}
	return state
}

// Returns a slice with all spells the player could cast based on the current game state.
// The player can not cast a spell that needs too much mana or if its effect is currently active.
func possibleSpells(state gameState) (possible []spell) {
	for _, spell := range spells {
		if state.player.mana >= spell.cost && state.activeEffects[spell.effectType] == 0 {
			possible = append(possible, spell)
		}
	}
	return possible
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
