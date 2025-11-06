package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var validNums, validNumsForEntry = make(map[int]bool, 0), make([]map[int]bool, 0)
	var validTickets, myTicket = make([][]int, 0), make([]int, 0)
	var min, max int

	// read entry valid intervals
	scanner := bufio.NewScanner(os.Stdin)
	for entry := 0; scanner.Scan(); entry++ {
		if len(scanner.Text()) == 0 {
			break
		}
		validNumsForEntry = append(validNumsForEntry, make(map[int]bool, 0))
		line := scanner.Text()
		fields := strings.Fields(line[strings.Index(line, ": ")+2:])
		for i := 0; i < len(fields); i += 2 {
			fmt.Sscanf(fields[i], "%d-%d", &min, &max)
			for n := min; n <= max; n++ {
				validNums[n] = true
				validNumsForEntry[entry][n] = true
			}
		}
	}

	// read own ticket values
	scanner.Scan() // ignore "your ticket" line
	scanner.Scan()
	for _, numStr := range strings.Split(scanner.Text(), ",") {
		num, _ := strconv.Atoi(numStr)
		myTicket = append(myTicket, num)
	}
	scanner.Scan() // ignore empty line

	// read nearby tickets
	scanner.Scan() // ignore "nearby tickets" line
	errorRate := 0
	for scanner.Scan() {
		ticket := make([]int, 0)
		valid := true
		for _, numStr := range strings.Split(scanner.Text(), ",") {
			num, _ := strconv.Atoi(numStr)
			ticket = append(ticket, num)
			if !validNums[num] {
				errorRate += num
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	// ==== Figure out order of entries ====

	var possible = make([][]bool, len(validNumsForEntry))
	for i := range validNumsForEntry {
		possible[i] = make([]bool, len(validNumsForEntry))
		for j := range validNumsForEntry {
			possible[i][j] = true
		}
	}

	for _, ticket := range validTickets {
		for pos, value := range ticket {
			for entry := range validNumsForEntry {
				if !validNumsForEntry[entry][value] {
					possible[entry][pos] = false
				}
			}
		}
	}

	dirty := true
	for dirty {
		dirty = false
		for pos1 := range validNumsForEntry {
			count := 0
			value := 0
			for pos2 := range validNumsForEntry {
				if possible[pos1][pos2] { // for entry `pos1` position `pos2` on the ticket is possible
					count++
					value = pos2
				}
			}
			if count == 1 { // for entry `pos1` only one position `valueF` is possible
				for pos2 := range validNumsForEntry {
					if pos2 != pos1 {
						possible[pos2][value] = false
					}
				}
			} else {
				dirty = true
			}
		}
	}

	correctPositions := make([]int, len(validNumsForEntry))
	for i := range validNumsForEntry {
		for j := range validNumsForEntry {
			if possible[i][j] {
				correctPositions[i] = j
			}
		}
	}

	productOfDepartureValues := 1
	for i := 0; i < 6; i++ {
		productOfDepartureValues *= myTicket[correctPositions[i]]
	}

	fmt.Printf("Ticket scanning error rate: %v\n", errorRate)
	fmt.Printf("Product of departure values: %v\n", productOfDepartureValues)
}
