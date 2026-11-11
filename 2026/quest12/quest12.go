package quest12

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

var height, width int

func spell_count(grid [][]byte) int {
	height = len(grid)
	width = len(grid[0])

	spells := 0

	for {
		longest_line := []Pos{}
		line := []Pos{}

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if grid[y][x] == '#' {
					line = append(line, Pos{x: x, y: y})
				} else if grid[y][x] == '.' && len(line) > 0 {
					if len(line) > len(longest_line) {
						longest_line = line
					}
					line = []Pos{}
				}
			}
			if len(line) > len(longest_line) {
				longest_line = line
			}
			line = []Pos{}
		}

		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				if grid[y][x] == '#' {
					line = append(line, Pos{x: x, y: y})
				} else if grid[y][x] == '.' && len(line) > 0 {
					if len(line) > len(longest_line) {
						longest_line = line
					}
					line = []Pos{}
				}
			}
			if len(line) > len(longest_line) {
				longest_line = line
			}
			line = []Pos{}
		}

		if len(longest_line) > 0 {
			spells++
			for _, pos := range longest_line {
				grid[pos.y][pos.x] = '.'
			}
		} else {
			return spells
		}
	}
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 12, 1

	grid := loader.GetByteRows()
	part1 := spell_count(grid)

	loader.Part = 2
	grid = loader.GetByteRows()
	part2 := spell_count(grid)

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
