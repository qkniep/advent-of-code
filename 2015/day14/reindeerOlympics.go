package main

import (
	"bufio"
	"fmt"
	"os"
)

const raceDuration = 2503
const statsFormat = "%s can fly %d km/s for %d seconds, but then must rest for %d seconds."

/*type reindeer struct {
	speed, flyTime, restTime int
}*/

func main() {
	//var reindeers []reindeer
	var winningDistance int
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var _name string
		var speed, flyTime, restTime int
		fmt.Sscanf(scanner.Text(), statsFormat, &_name, &speed, &flyTime, &restTime)
		distance := distanceTraveledDuringRace(speed, flyTime, restTime)
		if distance > winningDistance {
			winningDistance = distance
		}
	}

	fmt.Println("Distance of winning reindeer:", winningDistance)
	//fmt.Println("Total brightness:", totalBrightness)
}

func distanceTraveledDuringRace(speed, flyTime, restTime int) int {
	lowerBound := raceDuration / (flyTime + restTime) * speed * flyTime
	lastSegmentLen := raceDuration % (flyTime + restTime)
	return lowerBound + speed * min(flyTime, lastSegmentLen)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
