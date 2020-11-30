package main

import "fmt"

func main() {
	list := createList(256)
	lengths := []int{227, 169, 3, 166, 246, 201, 0, 47, 1, 255, 2, 254, 96, 3, 97, 144}
	prepareList(list, lengths)
	fmt.Printf("Product of first two values: %v\n", list[0]*list[1])
	//fmt.Printf("Garbaged Characters: %v\n", garbaged)
}

func createList(size int) (list []int) {
	list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}
	return
}

func prepareList(list []int, lengths []int) {
	currentPos := 0
	skipSize := 0
	for _, l := range lengths {
		a, b := currentPos, currentPos+l-1
		for a < b {
			tmp := list[a%len(list)]
			list[a%len(list)] = list[b%len(list)]
			list[b%len(list)] = tmp
			a, b = a+1, b-1
		}
		currentPos += l + skipSize
		skipSize += 1
	}
}
