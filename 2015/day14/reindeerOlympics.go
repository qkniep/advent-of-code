package main

import (
	"bufio"
	"fmt"
	"os"
)

const raceDuration = 2503
const statsFormat = "%s can fly %d km/s for %d seconds, but then must rest for %d seconds."

type reindeer struct {
	name     string
	speed    int
	flyTime  int
	restTime int
}

func main() {
	var reindeers []reindeer
	var winningDistance int
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var r reindeer
		fmt.Sscanf(scanner.Text(), statsFormat, &r.name, &r.speed, &r.flyTime, &r.restTime)
		distance := distanceTraveledDuringRace(r)
		if distance > winningDistance {
			winningDistance = distance
		}
		reindeers = append(reindeers, r)
	}

	fmt.Println("Distance of winning reindeer:", winningDistance)
	fmt.Println("Most time ahead of any reindeer:", winnerTimeAhead(reindeers))
}

func distanceTraveledDuringRace(r reindeer) int {
	lowerBound := raceDuration / (r.flyTime + r.restTime) * r.speed * r.flyTime
	lastSegmentLen := raceDuration % (r.flyTime + r.restTime)
	return lowerBound + r.speed*min(r.flyTime, lastSegmentLen)
}

func winnerTimeAhead(rs []reindeer) int {
	var distances = make([]int, len(rs))
	var scores = make([]int, len(rs))

	for ri := range rs {
		distances[ri] = 0
		scores[ri] = 0
	}

	// simulate race keeping track of scores
	for t := 0; t < raceDuration; t++ {
		for ri, r := range rs {
			if t%(r.flyTime+r.restTime) < r.flyTime {
				distances[ri] += r.speed
			}
		}
		maxDistance, rAhead := 0, []int{0}
		for ri := range rs {
			if distances[ri] > maxDistance {
				maxDistance = distances[ri]
				rAhead = []int{ri}
			} else if distances[ri] == maxDistance {
				rAhead = append(rAhead, ri)
			}
		}
		for _, ri := range rAhead {
			scores[ri]++
		}
	}

	// find max score
	var maxScore int
	for ri := range rs {
		if scores[ri] > maxScore {
			maxScore = scores[ri]
		}
	}

	return maxScore
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
