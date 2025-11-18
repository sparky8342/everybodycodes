package quest12

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

var width, height int
var dirs [4][2]int

func init() {
	dirs = [4][2]int{
		[2]int{0, 1},
		[2]int{0, -1},
		[2]int{1, 0},
		[2]int{-1, 0},
	}
}

func bfs(grid []string, queue []Pos, visited map[Pos]struct{}) {
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
}

func copy_visited(v map[Pos]struct{}) map[Pos]struct{} {
	c := map[Pos]struct{}{}
	for k := range v {
		c[k] = struct{}{}
	}
	return c
}

func find_best_shot(grid []string, available map[Pos]struct{}, current_visited map[Pos]struct{}) map[Pos]struct{} {
	max := 0
	best_visited := map[Pos]struct{}{}

	for start := range available {
		queue := []Pos{start}
		visited := copy_visited(current_visited)
		visited[start] = struct{}{}
		bfs(grid, queue, visited)
		if len(visited) > max {
			max = len(visited)
			best_visited = visited
		}
	}

	for pos := range best_visited {
		delete(available, pos)
	}

	return best_visited
}

func shoot_barrels(grid []string, shot_mode int) int {
	height = len(grid)
	width = len(grid[0])

	switch shot_mode {
	case 1:
		start := Pos{x: 0, y: 0}
		queue := []Pos{start}
		visited := map[Pos]struct{}{}
		visited[start] = struct{}{}
		bfs(grid, queue, visited)
		return len(visited)
	case 2:
		start1 := Pos{x: 0, y: 0}
		start2 := Pos{x: width - 1, y: height - 1}
		queue := []Pos{start1, start2}
		visited := map[Pos]struct{}{}
		visited[start1] = struct{}{}
		visited[start2] = struct{}{}
		bfs(grid, queue, visited)
		return len(visited)
	case 3:
		available := map[Pos]struct{}{}
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				available[Pos{x: x, y: y}] = struct{}{}
			}
		}

		visited := map[Pos]struct{}{}
		for i := 0; i < 3; i++ {
			visited = find_best_shot(grid, available, visited)
		}

		return len(visited)
	}

	return -1
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 12, 1

	grid := loader.GetStrings()
	part1 := shoot_barrels(grid, 1)

	loader.Part = 2
	grid = loader.GetStrings()
	part2 := shoot_barrels(grid, 2)

	loader.Part = 3
	grid = loader.GetStrings()
	part3 := shoot_barrels(grid, 3)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
