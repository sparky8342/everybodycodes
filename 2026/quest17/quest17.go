package quest17

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

type Entry struct {
	pos      Pos
	distance int
	prev     *Entry
}

type Dir struct {
	dx int
	dy int
}

var height, width int

var dirs [4]Dir

func init() {
	dirs = [4]Dir{
		Dir{dx: 1, dy: 0},
		Dir{dx: -1, dy: 0},
		Dir{dx: 0, dy: 1},
		Dir{dx: 0, dy: -1},
	}
}

func edge_distance(grid []string) int {
	height = len(grid)
	width = len(grid[0])

	distance := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '#' {
				if x == 0 || x == width-1 {
					distance++
				}
				if y == 0 || y == height-1 {
					distance++
				}
			}
			if x > 0 {
				if grid[y][x] != grid[y][x-1] {
					distance++
				}
			}
			if y > 0 {
				if grid[y][x] != grid[y-1][x] {
					distance++
				}
			}
		}
	}

	return distance
}

func shortest_path(grid []string) int {
	height = len(grid)
	width = len(grid[0])

	start := Pos{}
	end := Pos{x: width - 1, y: height - 1}

	queue := []Entry{{pos: start}}
	visited := map[Pos]struct{}{start: struct{}{}}

	for len(queue) > 0 {
		entry := queue[0]
		queue = queue[1:]

		pos := entry.pos
		if pos == end {
			return entry.distance
		}

		for _, dir := range dirs {
			new_pos := Pos{x: pos.x + dir.dx, y: pos.y + dir.dy}
			if new_pos.x < 0 || new_pos.x == width || new_pos.y < 0 || new_pos.y == height {
				continue
			}
			if grid[new_pos.y][new_pos.x] == '#' {
				if _, ok := visited[new_pos]; !ok {
					queue = append(queue, Entry{pos: new_pos, distance: entry.distance + 1})
					visited[new_pos] = struct{}{}
				}
			}
		}
	}

	return -1
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 17, 1

	grid := loader.GetStrings()
	part1 := shortest_path(grid)

	loader.Part = 2
	grid = loader.GetStrings()
	part2 := edge_distance(grid)

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
