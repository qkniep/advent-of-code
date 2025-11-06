package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const size = 5

var n2r = make(map[int]int, 0)

func main() {
	var nums []int
	var boards [][][]int

	// read numbers that are drawn
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numsStr := strings.Split(scanner.Text(), ",")
	for round, ns := range numsStr {
		n, _ := strconv.Atoi(ns)
		nums = append(nums, n)
		n2r[n] = round
	}

	// read bingo boards
	i := 0
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		if i%5 == 0 {
			boards = append(boards, make([][]int, size))
		}
		boards[i/5][i%5] = make([]int, size)
		row := strings.Fields(scanner.Text())
		for c, ns := range row {
			n, _ := strconv.Atoi(ns)
			boards[i/5][i%5][c] = n
		}
		i++
	}

	var firstWinBoard, firstWinRound = -1, 99999
	var lastWinBoard, lastWinRound = -1, -1
	for i, board := range boards {
		round := winningRound(board)
		if round < firstWinRound {
			firstWinBoard, firstWinRound = i, round
		}
		if round > lastWinRound {
			lastWinBoard, lastWinRound = i, round
		}
	}

	firstScore := nums[firstWinRound] * sumUndrawn(boards[firstWinBoard], firstWinRound)
	lastScore := nums[lastWinRound] * sumUndrawn(boards[lastWinBoard], lastWinRound)

	fmt.Println("First winning board's score:", firstScore)
	fmt.Println("Last winning board's score:", lastScore)
}

func winningRound(board [][]int) int {
	return min(
		// cols
		max(n2r[board[0][0]], n2r[board[1][0]], n2r[board[2][0]], n2r[board[3][0]], n2r[board[4][0]]),
		max(n2r[board[0][1]], n2r[board[1][1]], n2r[board[2][1]], n2r[board[3][1]], n2r[board[4][1]]),
		max(n2r[board[0][2]], n2r[board[1][2]], n2r[board[2][2]], n2r[board[3][2]], n2r[board[4][2]]),
		max(n2r[board[0][3]], n2r[board[1][3]], n2r[board[2][3]], n2r[board[3][3]], n2r[board[4][3]]),
		max(n2r[board[0][4]], n2r[board[1][4]], n2r[board[2][4]], n2r[board[3][4]], n2r[board[4][4]]),
		// rows
		max(n2r[board[0][0]], n2r[board[0][1]], n2r[board[0][2]], n2r[board[0][3]], n2r[board[0][4]]),
		max(n2r[board[1][0]], n2r[board[1][1]], n2r[board[1][2]], n2r[board[1][3]], n2r[board[1][4]]),
		max(n2r[board[2][0]], n2r[board[2][1]], n2r[board[2][2]], n2r[board[2][3]], n2r[board[2][4]]),
		max(n2r[board[3][0]], n2r[board[3][1]], n2r[board[3][2]], n2r[board[3][3]], n2r[board[3][4]]),
		max(n2r[board[4][0]], n2r[board[4][1]], n2r[board[4][2]], n2r[board[4][3]], n2r[board[4][4]]),
	)
}

// Sums numbers not yet drawn
func sumUndrawn(board [][]int, round int) (sum int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if n2r[board[y][x]] > round {
				sum += board[y][x]
			}
		}
	}
	return
}

func min(is ...int) int {
	min := is[0]
	for _, i := range is[1:] {
		if i < min {
			min = i
		}
	}
	return min
}

func max(is ...int) int {
	max := is[0]
	for _, i := range is[1:] {
		if i > max {
			max = i
		}
	}
	return max
}
