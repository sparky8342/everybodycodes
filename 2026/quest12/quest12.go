package quest12

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

type Dir struct {
	dx int
	dy int
}

var height, width int

var dirs []Dir

func init() {
	dirs = []Dir{
		Dir{dx: +1, dy: +0},
		Dir{dx: -1, dy: +0},
		Dir{dx: +0, dy: +1},
		Dir{dx: +0, dy: -1},
	}
}

func find_rectangle(grid [][]byte, x int, y int) int {
	start := Pos{x: x, y: y}
	queue := []Pos{start}
	visited := map[Pos]struct{}{start: struct{}{}}

	x_min, x_max, y_min, y_max := x, x, y, y

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			new_pos := Pos{x: pos.x + dir.dx, y: pos.y + dir.dy}

			if new_pos.y < 0 || new_pos.y == height || new_pos.x < 0 || new_pos.x == width {
				continue
			}

			if grid[new_pos.y][new_pos.x] == '#' {
				if _, ok := visited[new_pos]; !ok {
					queue = append(queue, new_pos)
					visited[new_pos] = struct{}{}

					if new_pos.x < x_min {
						x_min = new_pos.x
					} else if new_pos.x > x_max {
						x_max = new_pos.x
					}

					if new_pos.y < y_min {
						y_min = new_pos.y
					} else if new_pos.y > y_max {
						y_max = new_pos.y
					}
				}
			}
		}
	}

	for pos := range visited {
		grid[pos.y][pos.x] = '.'
	}

	w := x_max - x_min
	h := y_max - y_min
	if w < h {
		return w + 1
	} else {
		return h + 1
	}
}

func spell_count(grid [][]byte) int {
	height = len(grid)
	width = len(grid[0])
	spells := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '#' {
				spells += find_rectangle(grid, x, y)
			}
		}
	}
	return spells
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 12, 1

	grid := loader.GetByteRows()
	part1 := spell_count(grid)

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
