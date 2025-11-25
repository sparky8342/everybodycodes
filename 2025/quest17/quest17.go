package quest17

import (
	"fmt"
	"loader"
)

func lava_spread(grid []string, radius int) int {
	size := len(grid)

	xv := size / 2
	yv := xv

	total := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if x == xv && y == yv {
				fmt.Print("@")
			} else if (xv-x)*(xv-x)+(yv-y)*(yv-y) <= radius*radius {
				total += int(grid[y][x] - '0')
				fmt.Print(string(grid[y][x]))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 17, 1

	grid := loader.GetStrings()
	part1 := lava_spread(grid, 10)

	fmt.Printf("%d %d %d\n", part1, 0, 0)
}
