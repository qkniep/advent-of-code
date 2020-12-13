package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	var busID, minDeparture = -1, 9999999
	var schedule = make([]int64, 0)

	// read input as 2 lines
	reader := bufio.NewReader(os.Stdin)
	line1, _ := reader.ReadString('\n')
	earliestPossibleDeparture, _ := strconv.Atoi(strings.TrimSpace(line1))
	line2, _ := reader.ReadString('\n')

	for _, busStr := range strings.Split(strings.TrimSpace(line2), ",") {
		if busStr == "x" {
			schedule = append(schedule, -1)
			continue
		}
		// build slice for part 2
		bus, _ := strconv.Atoi(busStr)
		schedule = append(schedule, int64(bus))
		// solve part 1 while reading input
		nextDeparture := findNextDeparture(earliestPossibleDeparture, bus)
		if nextDeparture < minDeparture {
			busID = bus
			minDeparture = nextDeparture
		}
	}

	fmt.Printf("Next bus ID * wait offset: %v\n", busID * (minDeparture-earliestPossibleDeparture))
	fmt.Printf("First timestamp with ordered offsets: %v\n", findFirstInOrderTime(schedule))
}

// Returns the next departure of bus after earliest timestamp.
func findNextDeparture(earliest int, bus int) int {
	return earliest + bus - (earliest % bus)
}

// Returns the first timestamp t for which each bus schedule[i] leaves at time t+i.
func findFirstInOrderTime(schedule []int64) int64 {
	var offsets, busses = make([]*big.Int, 0), make([]*big.Int, 0)
	for offset, bus := range schedule {
		if bus < 0 {
			continue
		}
		offsets = append(offsets, big.NewInt(bus-int64(offset)))
		busses = append(busses, big.NewInt(bus))
	}
	return crt(offsets, busses).Int64()
}

// Chinese Remainder Theorem
// Returns the smallest number x which fulfills equations x=a[i] (mod n[i]) for all i.
func crt(a, n []*big.Int) *big.Int {
	var p = new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p)
}
