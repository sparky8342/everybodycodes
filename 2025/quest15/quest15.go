package quest15

import (
	"fmt"
	"loader"
	"strings"
)

type Move struct {
	dir    byte
	amount int
}

type Pos struct {
	x int
	y int
}

type Entry struct {
	pos      Pos
	distance int
}

var cardinal_dirs [4][2]int

func init() {
	cardinal_dirs = [4][2]int{
		[2]int{1, 0},
		[2]int{-1, 0},
		[2]int{0, 1},
		[2]int{0, -1},
	}
}

func parse_data(data []byte) []Move {
	parts := strings.Split(string(data), ",")
	moves := make([]Move, len(parts))
	for i, part := range parts {
		moves[i] = Move{
			dir:    part[0],
			amount: int(part[1] - '0'),
		}
	}
	return moves
}

func find_exit(moves []Move) int {
	dirs := [4]byte{'U', 'R', 'D', 'L'}
	current_dir := 0

	walls := map[Pos]struct{}{}

	start := Pos{}
	pos := Pos{}

	for _, move := range moves {
		if move.dir == 'L' {
			if current_dir == 0 {
				current_dir = 3
			} else {
				current_dir--
			}
		} else if move.dir == 'R' {
			if current_dir == 3 {
				current_dir = 0
			} else {
				current_dir++
			}
		}

		for i := 0; i < move.amount; i++ {
			switch dirs[current_dir] {
			case 'U':
				pos.y--
			case 'R':
				pos.x++
			case 'D':
				pos.y++
			case 'L':
				pos.x--
			}
			walls[pos] = struct{}{}
		}
	}

	end := pos
	delete(walls, end)

	queue := []Entry{Entry{pos: start}}
	visited := map[Pos]struct{}{}
	visited[start] = struct{}{}

	for len(queue) > 0 {
		entry := queue[0]
		queue = queue[1:]

		pos := entry.pos

		if pos.x == end.x && pos.y == end.y {
			return entry.distance
		}

		for _, dir := range cardinal_dirs {
			new_pos := Pos{x: pos.x + dir[0], y: pos.y + dir[1]}
			if _, ok := walls[new_pos]; ok {
				continue
			}
			if _, ok := visited[new_pos]; ok {
				continue
			}
			queue = append(queue, Entry{pos: new_pos, distance: entry.distance + 1})
			visited[new_pos] = struct{}{}
		}
	}

	return -1
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 15, 1

	data := loader.GetOneLine()
	moves := parse_data(data)

	part1 := find_exit(moves)

	fmt.Printf("%d %d %d\n", part1, 0, 0)
}
