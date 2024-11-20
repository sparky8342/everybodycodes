package quest3

import (
	"fmt"
	"loader"
)

func dig(data []string) int {
	height := len(data)
	width := len(data[0])

	blocks := 0
	grid := make([][]byte, height)
	for i, line := range data {
		grid_line := []byte(line)
		for j := 0; j < width; j++ {
			if grid_line[j] == '#' {
				grid_line[j] = '1'
				blocks++
			}
		}
		grid[i] = grid_line
	}

	dirs := [][]int{
		[]int{0, 1},
		[]int{0, -1},
		[]int{1, 0},
		[]int{-1, 0},
	}

	var num byte = '1'
	last_blocks := 0
	for blocks > last_blocks {
		last_blocks = blocks
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if grid[y][x] == num {
					neighbours := 0
					for _, dir := range dirs {
						if grid[y+dir[1]][x+dir[0]] >= num {
							neighbours++
						}
					}
					if neighbours == 4 {
						grid[y][x]++
						blocks++
					}
				}
			}
		}
		num++
	}

	return blocks
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 3, 1

	data := loader.GetStrings()
	part1 := dig(data)

	loader.Part = 2
	data = loader.GetStrings()
	part2 := dig(data)

	fmt.Printf("%d %d %d\n", part1, part2, -1)
}
