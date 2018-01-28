package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	memory := readInput()
	toLoop, inLoop := findLoop(memory)
	fmt.Printf("Cycles To Loop: %v\n", toLoop)
	fmt.Printf("Cycles In Loop: %v\n", inLoop)
}

func readInput() (memory []int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		fmt.Println(scanner.Text())
		for _, s := range strings.Fields(scanner.Text()) {
			blocks, _ := strconv.Atoi(s)
			memory = append(memory, blocks)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func findLoop(memory []int) (cyclesToLoop int, cyclesInLoop int) {
	states := make(map[string]int)
	key := generateKey(memory)

	for states[key] == 0 {
		states[key] = cyclesToLoop + 1

		// find position of maximum
		maxPos := 0
		for i, v := range memory {
			if v > memory[maxPos] {
				maxPos = i
			}
		}

		// redistribute
		blocks := memory[maxPos]
		memory[maxPos] = 0
		i := (maxPos + 1) % len(memory)
		for ; blocks > 0; blocks-- {
			memory[i] += 1
			i = (i + 1) % len(memory)
		}

		key = generateKey(memory)
		cyclesToLoop += 1
	}

	cyclesInLoop = cyclesToLoop + 1 - states[key]
	return
}

func generateKey(mem []int) string {
	return fmt.Sprint(mem)
}
