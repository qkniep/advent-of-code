package main

import (
	"fmt"
	"math/big"
)

const inFormat = "To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d."

func main() {
	var row, column int64

	fmt.Scanf(inFormat, &row, &column)

	cn := codeNumber(row, column)
	res := fastModPow(252533, cn-1, 33554393)
	code := (20151125 * res) % 33554393

	fmt.Printf("The code at position (row=%d, col=%d) in the manual: %d\n", row, column, code)
}

func codeNumber(row, column int64) int64 {
	firstInCol := column * (column+1) / 2
	rowPart := (column+row-2) * (column+row-1) / 2 - column * (column-1) / 2
	return firstInCol + rowPart
}

func fastModPow(base, exp, mod int64) int64 {
	b := big.NewInt(base)
	e := big.NewInt(exp)
	m := big.NewInt(mod)
	return b.Exp(b, e, m).Int64()
}
