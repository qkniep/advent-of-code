package main

import "fmt"

type character struct {
	hitPoints, damage, armor int
}

type item struct {
	cost, damage, armor int
}

func main() {
	var weapons = []item{
		{ 8, 4, 0},
		{10, 5, 0},
		{25, 6, 0},
		{40, 7, 0},
		{74, 8, 0},
	}
	var armors = []item{
		{  0, 0, 0},
		{ 13, 0, 1},
		{ 31, 0, 2},
		{ 53, 0, 3},
		{ 75, 0, 4},
		{102, 0, 5},
	}
	var rings = []item {
		{  0, 0, 0},
		{ 25, 1, 0},
		{ 50, 2, 0},
		{100, 3, 0},
		{ 20, 0, 1},
		{ 40, 0, 2},
		{ 80, 0, 3},
	}
	var player = character{100, 0, 0}
	var boss character

	// read boss' stats
	fmt.Scanf("Hit Points: %d", &boss.hitPoints)
	fmt.Scanf("Damage: %d", &boss.damage)
	fmt.Scanf("Armor: %d", &boss.armor)

	var minGold = 99999
	var maxGold = -1
	for _, weapon := range weapons {
		for _, armor := range armors {
			for r1, ring1 := range rings {
				for r2, ring2 := range rings {
					if r1 == r2 && r1 != 0 { // cannot purchase same ring twice
						continue
					}
					totalGold := weapon.cost + armor.cost + ring1.cost + ring2.cost
					if simulateFight(player, boss, [4]item{weapon, armor, ring1, ring2}) {
						if totalGold < minGold {
							minGold = totalGold
						}
					} else if totalGold > maxGold {
						maxGold = totalGold
					}
				}
			}
		}
	}

	fmt.Printf("Minimum amount of gold to win: %d\n", minGold)
	fmt.Printf("Maximum amount of gold to still lose: %d\n", maxGold)
}

// Returns true iff the player wins the fight using the given items.
func simulateFight(player character, enemy character, items [4]item) bool {
	var itemDamage = items[0].damage + items[1].damage + items[2].damage + items[3].damage
	var itemArmor = items[0].armor + items[1].armor + items[2].armor + items[3].armor
	// damage per round (DPR)
	var playerDPR = max(1, (player.damage + itemDamage) - enemy.armor)
	var enemyDPR = max(1, enemy.damage - (player.armor + itemArmor))

	var roundsToKillEnemy = (enemy.hitPoints + playerDPR - 1) / playerDPR
	var roundsToKillPlayer = (player.hitPoints + enemyDPR - 1) / enemyDPR

	return roundsToKillEnemy <= roundsToKillPlayer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
