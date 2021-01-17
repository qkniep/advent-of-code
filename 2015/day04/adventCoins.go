package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

type pos struct {
	x, y int
}

func main() {
	var collision5, collision6 = 0, 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println("Mining AdventCoins...")
	for nonce := 0; collision5 == 0 || collision6 == 0; nonce++ {
		hash := md5.Sum([]byte("bgvyzdsv" + strconv.Itoa(nonce)))
		leadingZeros := 0
		for _, h := range hex.EncodeToString(hash[:]) {
			if h != '0' {
				break
			}
			leadingZeros++
		}
		if leadingZeros >= 5 && collision5 == 0 {
			collision5 = nonce
		}
		if leadingZeros >= 6 && collision6 == 0 {
			collision6 = nonce
		}
	}

	fmt.Printf("First 5 digit collision: %v\n", collision5)
	fmt.Printf("First 6 digit collision: %v\n", collision6)
}

func applyChange(p *pos, dir rune) {
	if dir == '<' {
		p.x--
	} else if dir == '>' {
		p.x++
	} else if dir == '^' {
		p.y++
	} else if dir == 'v' {
		p.y--
	}
}
