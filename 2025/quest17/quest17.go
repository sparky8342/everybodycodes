package quest17

import (
	"fmt"
	"loader"
)

func pow2(n int) int {
	return n * n
}

func lava_spread(grid []string, radius int) int {
	size := len(grid)

	cv := size / 2
	square_radius := radius * radius

	total := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if x == cv && y == cv {
				fmt.Print("@")
			} else if pow2(cv-x)+pow2(cv-y) <= square_radius {
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
