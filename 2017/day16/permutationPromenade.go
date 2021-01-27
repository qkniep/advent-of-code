package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const numPrograms = 16
const numIterations = 1_000_000_000

func main() {
	var programs string
	var dance []string
	var scanner = bufio.NewScanner(os.Stdin)

	for p := 0; p < numPrograms; p++ {
		programs += string('a' + rune(p))
	}

	scanner.Scan()
	dance = strings.Split(scanner.Text(), ",")

	// find period length
	for period := 1; ; period++ {
		performDance(&programs, dance)

		isOriginal := true
		for i, r := range programs {
			if r != 'a' + rune(i) {
				isOriginal = false
			}
		}

		if isOriginal {
			// perform only the necessary iterations
			for i := 0; i < numIterations % period; i++ {
				performDance(&programs, dance)
			}
			break
		}
	}

	fmt.Printf("Program order after the first dance: %v\n", programs)
	fmt.Printf("Order after 1,000,000,000 dances: %v\n", programs)
}

func performDance(programs *string, dance []string) {
	for _, move := range dance {
		if strings.IndexRune(move, 's') == 0 {
			spin, _ := strconv.Atoi(move[1:])
			*programs = (*programs)[len(*programs)-spin:] + (*programs)[:len(*programs)-spin]
		} else if strings.IndexRune(move, 'x') == 0 {
			is := strings.Split(move[1:], "/")
			i1, _ := strconv.Atoi(is[0])
			i2, _ := strconv.Atoi(is[1])
			ps := []byte(*programs)
			ps[i1], ps[i2] = ps[i2], ps[i1]
			*programs = string(ps)
		} else if strings.IndexRune(move, 'p') == 0 {
			i1 := strings.IndexByte(*programs, move[1])
			i2 := strings.IndexByte(*programs, move[3])
			ps := []byte(*programs)
			ps[i1], ps[i2] = ps[i2], ps[i1]
			*programs = string(ps)
		}
	}
	return
}
