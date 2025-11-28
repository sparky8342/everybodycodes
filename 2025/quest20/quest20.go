package quest20

import (
	"fmt"
	"loader"
)

func count_pairs(grid []string) int {
	height := len(grid)
	width := len(grid[0])

	pairs := 0
	for y := 0; y < height-1; y++ {
		for x := 0; x < width-1; x++ {
			if grid[y][x] == 'T' {
				if grid[y][x+1] == 'T' {
					pairs++
				}
				if (y%2 != x%2) && grid[y+1][x] == 'T' {
					pairs++
				}
			}
		}
	}

	return pairs
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 20, 1

	grid := loader.GetStrings()
	part1 := count_pairs(grid)

	fmt.Printf("%d %d %d\n", part1, 0, 0)
}
