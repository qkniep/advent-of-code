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

	scanner.Scan() // ignore "your ticket" line
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		for _, numStr := range strings.Split(scanner.Text(), ",") {
			num, _ := strconv.Atoi(numStr)
			myTicket = append(myTicket, num)
		}
	}

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
					fmt.Printf("Offense: %v not allowed for entry %v, found at pos %v\n", value, entry, pos)
					possible[entry][pos] = false
				}
			}
		}
	}

	fmt.Printf("%v\n", possible)

	dirty := true
	for dirty {
		dirty = false
		for pos1 := range validNumsForEntry {
			countF, countB := 0, 0
			valueF, valueB := 0, 0
			//countF := 0
			//valueF := 0
			for pos2 := range validNumsForEntry {
				if possible[pos1][pos2] { // for entry `pos1` position `pos2` on the ticket is possible
					countF++
					valueF = pos2
				}
				if possible[pos2][pos1] { // for entry `pos2` position `pos1` on the ticket is possible
					countB++
					valueB = pos2
				}
			}
			// forward
			if countF == 1 { // for entry `pos1` only one position `valueF` is possible
				for pos2 := range validNumsForEntry {
					if pos2 != pos1 {
						possible[pos2][valueF] = false
					}
				}
			} else {
				dirty = true
			}
			// backward
			if countB == 1 { // for position `pos1` only one entry `valueB` is possible
				for pos2 := range validNumsForEntry {
					if pos2 != pos1 {
						possible[valueB][pos2] = false
					}
				}
			} else {
				dirty = true
			}
		}

	fmt.Printf("%v\n", possible)
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
