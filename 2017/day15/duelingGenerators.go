package main

import "fmt"

const judgeNumPairs = 40_000_000
const pickyJudgeNumPairs = 5_000_000
const judgePrecision = 65536

type generator struct {
	state, factor, mod, pickiness int64
}

func (g *generator) next() int64 {
	(*g).state = (g.state * g.factor) % g.mod
	for g.state % g.pickiness != 0 {
		(*g).state = (g.state * g.factor) % g.mod
	}
	return g.state
}

func main() {
	var matches, pickyMatches, startA, startB int64

	fmt.Scanf("Generator A starts with %d", &startA)
	fmt.Scanf("Generator B starts with %d", &startB)

	generatorA := generator{startA, 16807, 2147483647, 1}
	generatorB := generator{startB, 48271, 2147483647, 1}
	matches = judgeGenerators(generatorA, generatorB, judgePrecision, judgeNumPairs)

	pickyGeneratorA := generator{startA, 16807, 2147483647, 4}
	pickyGeneratorB := generator{startB, 48271, 2147483647, 8}
	pickyMatches = judgeGenerators(pickyGeneratorA, pickyGeneratorB, judgePrecision, pickyJudgeNumPairs)

	fmt.Printf("Number of times the lower 16 bits matched: %v\n", matches)
	fmt.Printf("Number of matches for the picky generators: %v\n", pickyMatches)
}

func judgeGenerators(g1, g2 generator, precision, pairs int64) (matches int64) {
	for p := int64(0); p < pairs; p++ {
		if g1.next() % precision == g2.next() % precision {
			matches++
		}
	}
	return
}
