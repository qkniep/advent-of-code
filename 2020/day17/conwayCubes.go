package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var cubes = make([][][][]bool, 32)
	var active = 0

	for w := 0; w < 32; w++ {
		cubes[w] = make([][][]bool, 32)
		for z := 0; z < 32; z++ {
			cubes[w][z] = make([][]bool, 32)
			for y := 0; y < 32; y++ {
				cubes[w][z][y] = make([]bool, 32)
			}
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for y := 0; scanner.Scan(); y++ {
		for x, char := range scanner.Text() {
			if char == '#' {
				cubes[16][16][16+y][16+x] = true
				active++
			}
		}
	}

	// backup map
	newCubes := make([][][][]bool, 32)
	deepCopy(newCubes, cubes)

	for round := 0; round < 6; round++ {
		for w := 0; w < len(cubes); w++ {
			for z := 0; z < len(cubes[w]); z++ {
				for y := 0; y < len(cubes[w][z]); y++ {
					for x := 0; x < len(cubes[w][z][y]); x++ {
						activeNgb := countActiveNeighbors(cubes, x, y, z, w)
						isActive := cubes[w][z][y][x]
						if isActive && (activeNgb < 2 || activeNgb > 3) {
							newCubes[w][z][y][x] = false
							active--
						} else if !isActive && activeNgb == 3 {
							newCubes[w][z][y][x] = true
							active++
						}
					}
				}
			}
		}
		deepCopy(cubes, newCubes)
	}

	fmt.Println(active)
}

func countActiveNeighbors(cubes [][][][]bool, xx int, yy int, zz int, ww int) (activeNgb int) {
	for w := -1; w <= 1; w++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if xx+x < 0 || yy+y < 0 || zz+z < 0 || ww+w < 0 || xx+x >= 32 || yy+y >= 32 || zz+z >= 32 || ww+w >= 32 {
						continue
					}
					if !(x == 0 && y == 0 && z == 0 && w == 0) && cubes[ww+w][zz+z][yy+y][xx+x] {
						activeNgb++
					}
				}
			}
		}
	}
	return
}

func deepCopy(dst [][][][]bool, src [][][][]bool) {
	for w := range src {
		dst[w] = make([][][]bool, len(src[w]))
		for z := range src {
			dst[w][z] = make([][]bool, len(src[w][z]))
			for y := range src {
				dst[w][z][y] = make([]bool, len(src[w][z][y]))
				copy(dst[w][z][y], src[w][z][y])
			}
		}
	}
}
