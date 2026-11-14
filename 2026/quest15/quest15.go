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
	prev          *QueueEntry
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

func bfs_cycle(grid []string, start Pos) (int, int) {
	queue := []QueueEntry{{pos: start}}

	for len(queue) > 0 {
		qe := queue[0]
		queue = queue[1:]

		pos := qe.pos
		if pos == start && qe.steps > 0 {
			return qe.steps, qe.sulphur_found
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
				new_qe := QueueEntry{pos: new_pos, dir: dir, sulphur_found: qe.sulphur_found, steps: qe.steps + 1}
				if grid[new_pos.y][new_pos.x] == '@' {
					new_qe.sulphur_found++
				}
				queue = append(queue, new_qe)
			}
		}
	}

	return -1, -1
}

func find_shortest_path_with_multiple_sulphur(grid []string, target_sulphur int) int {
	height = len(grid)
	width = len(grid[0])

	start, end, _ := find_poi(grid)

	queue := []QueueEntry{{pos: start}}
	visited := map[[5]int]struct{}{{start.x, start.y, 1, 0, 0}: struct{}{}}

	for len(queue) > 0 {
		qe := queue[0]
		queue = queue[1:]

		pos := qe.pos

		if pos == end && qe.sulphur_found >= target_sulphur {
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
				new_sulphur := qe.sulphur_found
				if grid[new_pos.y][new_pos.x] == '@' {
					new_sulphur++
				}

				if _, ok := visited[[5]int{new_pos.x, new_pos.y, dir.dx, dir.dy, new_sulphur}]; !ok {
					new_qe := QueueEntry{pos: new_pos, dir: dir, sulphur_found: new_sulphur, steps: qe.steps + 1, prev: &qe}
					queue = append(queue, new_qe)
					visited[[5]int{new_pos.x, new_pos.y, dir.dx, dir.dy, new_sulphur}] = struct{}{}
				}
			}
		}
	}

	return -1
}

func find_shortest_path_with_multiple_sulphur3(grid []string, target_sulphur int) int {
	height = len(grid)
	width = len(grid[0])

	start, end, _ := find_poi(grid)

	data_grid := make([][][2]int, height)
	for i := range data_grid {
		data_grid[i] = make([][2]int, width)
		for j := 0; j < width; j++ {
			data_grid[i][j] = [2]int{99999, 0}
		}
	}

	min_steps := 999999999999
	dfs(grid, end, start.x, start.y, 0, 0, 0, target_sulphur, 0, &min_steps, data_grid)
	return min_steps
}

func dfs(grid []string, end Pos, x int, y int, dx int, dy int, sulphur int, target int, steps int, min_steps *int, data_grid [][][2]int) {
	if sulphur >= target && x == end.x && y == end.y {
		if steps < *min_steps {
			*min_steps = steps
			fmt.Printf("%d found in %d steps\n", sulphur, steps)
		}
		return
	}

	if steps >= *min_steps {
		return
	}

	prev_steps, prev_sulphur := data_grid[y][x][0], data_grid[y][x][1]
	if prev_sulphur > sulphur {
		return
	}
	if prev_sulphur == sulphur && steps > prev_steps {
		//fmt.Println(x, y, steps, sulphur, prev_steps, prev_sulphur)
		return
	}

	data_grid[y][x] = [2]int{steps, sulphur}

	for _, dir := range dirs {
		if dir.dx == dx*-1 && dir.dy == dy*-1 {
			continue
		}
		new_x, new_y := x+dir.dx, y+dir.dy
		if new_x < 0 || new_x == width || new_y < 0 || new_y == height {
			continue
		}
		if grid[new_y][new_x] != '#' {
			new_sulphur := sulphur
			if grid[new_y][new_x] == '@' {
				new_sulphur++
			}
			dfs(grid, end, new_x, new_y, dir.dx, dir.dy, new_sulphur, target, steps+1, min_steps, data_grid)
		}
	}

}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 15, 1

	grid := loader.GetStrings()
	part1 := find_shortest_path(grid)

	loader.Part = 2
	grid = loader.GetStrings()
	part2 := find_shortest_path_with_multiple_sulphur(grid, 2026)

	loader.Part = 3
	grid = loader.GetStrings()
	part3 := -1
	//part3 := find_shortest_path_with_multiple_sulphur2(grid, 202620262026)
	// 377292901683 wrong
	// 1215721526613 wrong

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
