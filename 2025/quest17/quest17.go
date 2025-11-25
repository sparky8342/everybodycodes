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

func max_ring(grid []string) int {
	size := len(grid)

	cv := size / 2

	max := 0
	max_radius := 0

	for radius := 1; radius <= cv; radius++ {
		square_radius := radius * radius
		prev_square_radius := pow2(radius - 1)

		lava := 0
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				if x == cv && y == cv {
					continue
				}

				r := pow2(cv-x) + pow2(cv-y)
				if r <= square_radius && r > prev_square_radius {
					lava += int(grid[y][x] - '0')
				}
			}
		}

		if lava > max {
			max = lava
			max_radius = radius
		}
	}

	return max * max_radius
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 17, 1

	grid := loader.GetStrings()
	part1 := lava_spread(grid, 10)

	loader.Part = 2
	grid = loader.GetStrings()
	part2 := max_ring(grid)

	fmt.Printf("%d %d %d\n", part1, part2, 0)
}
