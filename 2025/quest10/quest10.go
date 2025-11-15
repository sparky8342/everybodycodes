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

var height, width int
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
	width = len(board[0])
	height = len(board)

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

func print_board(board [][]byte) {
	for _, row := range board {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func find_max_sheep(start_board []string, turns int) int {
	width = len(start_board[0])
	height = len(start_board)

	board := make([][]byte, len(start_board))
	for i, row := range start_board {
		board[i] = []byte(row)
	}

	start := Pos{x: width / 2, y: height / 2}
	board[start.y][start.x] = '.'

	dd_positions := map[Pos]struct{}{}
	dd_positions[start] = struct{}{}

	sheep := 0
	for i := 0; i < turns; i++ {
		next_positions := map[Pos]struct{}{}
		for pos := range dd_positions {
			for _, dir := range dirs {
				new_pos := Pos{x: pos.x + dir[0], y: pos.y + dir[1]}
				if new_pos.x < 0 || new_pos.x >= width || new_pos.y < 0 || new_pos.y >= height {
					continue
				}
				next_positions[new_pos] = struct{}{}
			}
		}
		dd_positions = next_positions

		for pos := range dd_positions {
			if board[pos.y][pos.x] == 'S' {
				board[pos.y][pos.x] = '.'
				sheep++
			}
		}

		for x := 0; x < width; x++ {
			if board[height-1][x] == 'S' {
				board[height-1][x] = '.'
			} else if board[height-1][x] == '@' {
				board[height-1][x] = '#'
			}
		}

		for y := height - 2; y >= 0; y-- {
			for x := 0; x < width; x++ {
				if board[y][x] == 'S' || board[y][x] == '@' {
					if board[y][x] == 'S' {
						board[y][x] = '.'
					} else if board[y][x] == '@' {
						board[y][x] = '#'
					}

					if board[y+1][x] == '#' {
						board[y+1][x] = '@'
					} else if board[y+1][x] == '.' {
						if _, ok := dd_positions[Pos{x: x, y: y + 1}]; ok {
							sheep++
						} else {
							board[y+1][x] = 'S'
						}
					}
				}
			}
		}
	}

	return sheep
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 10, 1

	board := loader.GetStrings()
	part1 := in_range(board, 4)

	loader.Part = 2
	board = loader.GetStrings()
	part2 := find_max_sheep(board, 20)

	fmt.Printf("%d %d %d\n", part1, part2, 0)
}
