package quest10

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

type Entry struct {
	pos   Pos
	moves int
}

var dirs [8][2]int

func init() {
	dirs = [8][2]int{
		[2]int{1, 2},
		[2]int{1, -2},
		[2]int{-1, 2},
		[2]int{-1, -2},
		[2]int{2, 1},
		[2]int{2, -1},
		[2]int{-2, 1},
		[2]int{-2, -1},
	}
}

func in_range(board []string, moves int) int {
	width := len(board[0])
	height := len(board)

	start := Pos{x: width / 2, y: height / 2}

	queue := []Entry{Entry{pos: start}}

	sheep := map[Pos]struct{}{}

	for len(queue) > 0 {
		entry := queue[0]
		queue = queue[1:]

		if entry.moves > moves {
			continue
		}

		pos := entry.pos

		if board[pos.y][pos.x] == 'S' {
			sheep[pos] = struct{}{}
		}

		for _, dir := range dirs {
			new_pos := Pos{x: pos.x + dir[0], y: pos.y + dir[1]}
			if new_pos.x < 0 || new_pos.x == width || new_pos.y < 0 || new_pos.y == height {
				continue
			}
			queue = append(queue, Entry{pos: new_pos, moves: entry.moves + 1})
		}
	}

	return len(sheep)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 10, 1

	board := loader.GetStrings()
	part1 := in_range(board, 4)

	fmt.Printf("%d %d %d\n", part1, 0, 0)
}
