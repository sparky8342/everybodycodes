package quest3

import (
	"fmt"
	"loader"
)

func _dig(data []string, diag bool) int {
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
	if diag {
		dirs = append(dirs, [][]int{
			[]int{1, 1},
			[]int{1, -1},
			[]int{-1, 1},
			[]int{-1, -1}}...)
	}

	var num byte = '1'
	last_blocks := 0
	for blocks > last_blocks {
		last_blocks = blocks
		for y := 0; y < height; y++ {
		loop:
			for x := 0; x < width; x++ {
				if grid[y][x] == num {
					for _, dir := range dirs {
						check_y := y + dir[1]
						if check_y < 0 || check_y == height {
							continue loop
						}
						check_x := x + dir[0]
						if check_x < 0 || check_x == width {
							continue loop
						}
						if grid[check_y][check_x] < num {
							continue loop
						}
					}
					grid[y][x]++
					blocks++
				}
			}
		}
		num++
	}

	return blocks
}

func dig(data []string) int {
	return _dig(data, false)
}

func dig_diag(data []string) int {
	return _dig(data, true)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 3, 1

	data := loader.GetStrings()
	part1 := dig(data)

	loader.Part = 2
	data = loader.GetStrings()
	part2 := dig(data)

	loader.Part = 3
	data = loader.GetStrings()
	part3 := dig_diag(data)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
