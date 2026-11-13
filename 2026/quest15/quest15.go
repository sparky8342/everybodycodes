package quest15

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

type QueueEntry struct {
	pos           Pos
	dir           Dir
	sulphur_found int
	steps         int
}

var width, height int

var dirs [4]Dir

func init() {
	dirs = [4]Dir{
		Dir{dx: +1, dy: +0},
		Dir{dx: -1, dy: +0},
		Dir{dx: +0, dy: +1},
		Dir{dx: +0, dy: -1},
	}
}

func find_poi(grid []string) (Pos, Pos, map[Pos]int) {
	start, end := Pos{}, Pos{}
	sulphur_id := 0
	sulphur := map[Pos]int{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'S' {
				start.x, start.y = x, y
			} else if grid[y][x] == 'E' {
				end.x, end.y = x, y
			} else if grid[y][x] == '@' {
				sulphur[Pos{x: x, y: y}] = sulphur_id
				sulphur_id++
			}
		}
	}
	return start, end, sulphur
}

func find_shortest_path(grid []string) int {
	height = len(grid)
	width = len(grid[0])

	start, end, sulphur := find_poi(grid)

	end_sulphur := 0
	for i := 0; i < len(sulphur); i++ {
		end_sulphur |= (1 << i)
	}

	queue := []QueueEntry{{pos: start}}
	visited := map[[3]int]struct{}{}

	for len(queue) > 0 {
		qe := queue[0]
		queue = queue[1:]

		pos := qe.pos
		if pos == end && qe.sulphur_found == end_sulphur {
			return qe.steps
		}

		for _, dir := range dirs {
			if dir.dx == qe.dir.dx*-1 && dir.dy == qe.dir.dy*-1 {
				continue
			}
			new_pos := Pos{pos.x + dir.dx, pos.y + dir.dy}
			if new_pos.x < 0 || new_pos.x == width || new_pos.y < 0 || new_pos.y == height {
				continue
			}
			if grid[new_pos.y][new_pos.x] != '#' {
				if _, ok := visited[[3]int{new_pos.x, new_pos.y, qe.sulphur_found}]; !ok {
					new_qe := QueueEntry{pos: new_pos, dir: dir, sulphur_found: qe.sulphur_found, steps: qe.steps + 1}
					if grid[new_pos.y][new_pos.x] == '@' {
						new_qe.sulphur_found |= (1 << sulphur[new_pos])
					}
					queue = append(queue, new_qe)
					visited[[3]int{new_pos.x, new_pos.y, new_qe.sulphur_found}] = struct{}{}
				}
			}
		}
	}

	return -1
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 15, 1

	grid := loader.GetStrings()
	part1 := find_shortest_path(grid)

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
