package quest12

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

var dirs [4][2]int

func init() {
	dirs = [4][2]int{
		[2]int{0, 1},
		[2]int{0, -1},
		[2]int{1, 0},
		[2]int{-1, 0},
	}
}

func shoot_barrels(grid []string, shot_mode int) int {
	height := len(grid)
	width := len(grid[0])

	visited := map[Pos]struct{}{}
	queue := []Pos{}

	if shot_mode == 1 {
		start := Pos{x: 0, y: 0}
		queue = []Pos{start}
		visited[start] = struct{}{}
	} else if shot_mode == 2 {
		start1 := Pos{x: 0, y: 0}
		start2 := Pos{x: width - 1, y: height - 1}
		queue = []Pos{start1, start2}
		visited[start1] = struct{}{}
		visited[start2] = struct{}{}
	}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			next_pos := Pos{x: pos.x + dir[0], y: pos.y + dir[1]}
			if next_pos.x < 0 || next_pos.x == width || next_pos.y < 0 || next_pos.y == height {
				continue
			}
			if grid[next_pos.y][next_pos.x] > grid[pos.y][pos.x] {
				continue
			}
			if _, ok := visited[next_pos]; !ok {
				queue = append(queue, next_pos)
				visited[next_pos] = struct{}{}
			}
		}
	}

	return len(visited)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 12, 1

	grid := loader.GetStrings()
	part1 := shoot_barrels(grid, 1)

	loader.Part = 2
	grid = loader.GetStrings()
	part2 := shoot_barrels(grid, 2)

	fmt.Printf("%d %d %d\n", part1, part2, 0)
}
