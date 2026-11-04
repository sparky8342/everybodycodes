package quest8

import (
	"fmt"
	"loader"
)

var width, height int

func energy_needed(grid []string) int {
	height = len(grid)
	width = len(grid[0])

	total_energy := 0

	x := 0
	for y := height / 2; y >= 0; y-- {
		total_energy += shine(grid, x, y)
		x++
	}

	x++
	for grid[0][x] != '.' {
		total_energy += shine(grid, x, 0)
		x += 2
	}

	return total_energy
}

func shine(grid []string, x int, y int) int {
	energy := 0
	shadow := 0
	for x < width && y < height && grid[y][x] != '.' {
		height := int(grid[y][x] - '0')
		if height > shadow {
			energy++
			shadow = height - 1
		} else {
			shadow--
		}
		x++
		y++
	}
	return energy
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 8, 1

	grid := loader.GetStrings()
	part1 := energy_needed(grid)

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
