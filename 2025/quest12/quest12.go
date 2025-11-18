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

func shoot_barrels(grid []string) int {
	height := len(grid)
	width := len(grid[0])

	start := Pos{x: 0, y: 0}
	visited := map[Pos]struct{}{}
	visited[start] = struct{}{}

	queue := []Pos{start}

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
	part1 := shoot_barrels(grid)

	fmt.Printf("%d %d %d\n", part1, 0, 0)
}
