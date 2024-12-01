package quest19

import (
	"fmt"
	"loader"
)

func parse_data(data []string) (string, [][]byte) {
	key := data[0]
	data = data[2:]
	grid := make([][]byte, len(data))
	for i, line := range data {
		grid[i] = []byte(line)
	}
	return key, grid
}

func rotate_left(grid [][]byte, x int, y int) {
	tmp := grid[y-1][x-1]
	grid[y-1][x-1] = grid[y-1][x]
	grid[y-1][x] = grid[y-1][x+1]
	grid[y-1][x+1] = grid[y][x+1]
	grid[y][x+1] = grid[y+1][x+1]
	grid[y+1][x+1] = grid[y+1][x]
	grid[y+1][x] = grid[y+1][x-1]
	grid[y+1][x-1] = grid[y][x-1]
	grid[y][x-1] = tmp
}

func rotate_right(grid [][]byte, x int, y int) {
	tmp := grid[y-1][x-1]
	grid[y-1][x-1] = grid[y][x-1]
	grid[y][x-1] = grid[y+1][x-1]
	grid[y+1][x-1] = grid[y+1][x]
	grid[y+1][x] = grid[y+1][x+1]
	grid[y+1][x+1] = grid[y][x+1]
	grid[y][x+1] = grid[y-1][x+1]
	grid[y-1][x+1] = grid[y-1][x]
	grid[y-1][x] = tmp
}

func decode(data []string, rounds int) string {
	key, grid := parse_data(data)

	height := len(grid)
	width := len(grid[0])

	for i := 0; i < rounds; i++ {
		key_pos := 0
		for y := 1; y < height-1; y++ {
			for x := 1; x < width-1; x++ {
				if key[key_pos] == 'L' {
					rotate_left(grid, x, y)
				} else if key[key_pos] == 'R' {
					rotate_right(grid, x, y)
				}
				key_pos++
				if key_pos == len(key) {
					key_pos = 0
				}
			}
		}
	}

	message := []byte{}
	in_message := false
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '>' {
				in_message = true
			} else if grid[y][x] == '<' {
				return string(message)
			} else if in_message {
				message = append(message, grid[y][x])
			}
		}
	}

	return ""
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 19, 1

	data := loader.GetStrings()
	part1 := decode(data, 1)

	loader.Part = 2
	data = loader.GetStrings()
	part2 := decode(data, 100)

	part3 := ""
	fmt.Printf("%s %s %s\n", part1, part2, part3)
}
