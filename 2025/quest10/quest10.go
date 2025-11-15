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

var cache map[int]int

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

func next_dragons(dragons map[Pos]struct{}) map[Pos]struct{} {
	next := map[Pos]struct{}{}
	for dragon := range dragons {
		for _, dir := range dirs {
			new_dragon := Pos{x: dragon.x + dir[0], y: dragon.y + dir[1]}
			if new_dragon.x < 0 || new_dragon.x == width || new_dragon.y < 0 || new_dragon.y == height {
				continue
			}
			next[new_dragon] = struct{}{}
		}
	}
	return next
}

func in_range(board []string, moves int) int {
	width = len(board[0])
	height = len(board)

	start := Pos{x: width / 2, y: height / 2}

	dragons := map[Pos]struct{}{}
	dragons[start] = struct{}{}

	for i := 0; i < moves; i++ {
		next := next_dragons(dragons)
		for dragon := range next {
			dragons[dragon] = struct{}{}
		}
	}

	sheep := 0
	for dragon := range dragons {
		if board[dragon.y][dragon.x] == 'S' {
			sheep++
		}
	}

	return sheep
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
		dd_positions = next_dragons(dd_positions)

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

func state_key(dragonduck Pos, sheep []Pos, sheep_move bool) int {
	key := (dragonduck.x+1)*8 + (dragonduck.y + 1)
	for _, s := range sheep {
		key = key*8 + (s.x + 1)
		key = key*8 + (s.y + 1)
	}
	if sheep_move {
		key *= -1
	}
	return key
}

func sheep_moves(board []string, sheep []Pos, dragonduck Pos) int {
	key := state_key(dragonduck, sheep, true)
	if val, ok := cache[key]; ok {
		return val
	}

	seq := 0

	for i, s := range sheep {
		if s.y == height-1 {
			continue
		}

		if dragonduck.x == s.x && dragonduck.y == s.y+1 && board[s.y+1][s.x] != '#' {
			if len(sheep) == 1 {
				seq += dragonduck_moves(board, sheep, dragonduck)
			}
			continue
		}

		sheep[i].y++
		seq += dragonduck_moves(board, sheep, dragonduck)
		sheep[i].y--
	}

	cache[key] = seq
	return seq
}

func dragonduck_moves(board []string, sheep []Pos, dragonduck Pos) int {
	key := state_key(dragonduck, sheep, false)
	if val, ok := cache[key]; ok {
		return val
	}

	seq := 0

outer:
	for _, dir := range dirs {
		new_pos := Pos{x: dragonduck.x + dir[0], y: dragonduck.y + dir[1]}
		if new_pos.x < 0 || new_pos.x >= width || new_pos.y < 0 || new_pos.y >= height {
			continue
		}

		for i, s := range sheep {
			if s.x == new_pos.x && s.y == new_pos.y && board[s.y][s.x] != '#' {
				if len(sheep) == 1 {
					seq++
					continue outer
				}

				cpy := make([]Pos, len(sheep))
				copy(cpy, sheep)
				cpy = append(cpy[:i], cpy[i+1:]...)

				seq += sheep_moves(board, cpy, new_pos)
				continue outer
			}
		}

		seq += sheep_moves(board, sheep, new_pos)

	}

	cache[key] = seq
	return seq
}

func find_sequences(board []string) int {
	cache = map[int]int{}

	width = len(board[0])
	height = len(board)

	dragonduck := Pos{x: width / 2, y: height - 1}

	// find sheep
	sheep := []Pos{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if board[y][x] == 'S' {
				sheep = append(sheep, Pos{x: x, y: y})
			}
		}
	}

	return sheep_moves(board, sheep, dragonduck)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 10, 1

	board := loader.GetStrings()
	part1 := in_range(board, 4)

	loader.Part = 2
	board = loader.GetStrings()
	part2 := find_max_sheep(board, 20)

	loader.Part = 3
	board = loader.GetStrings()
	part3 := find_sequences(board)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
