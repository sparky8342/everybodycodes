package quest1

import (
	"fmt"
	"loader"
)

func parse_data(data []string) ([]string, []string) {
	var grid, tokens []string

	for i, row := range data {
		if row == "" {
			grid = data[0:i]
			tokens = data[i+1:]
			break
		}
	}

	return grid, tokens
}

func play_tokens(grid []string, tokens []string) int {
	height := len(grid)
	width := len(grid[0])

	total_score := 0

	for i := 0; i < len(tokens); i++ {
		toss_slot := i + 1

		x := i * 2
		y := -1

		token_i := 0

		var final_slot int

	outer:
		for {
			for grid[y+1][x] == '.' {
				y++
				if y == height-1 {
					final_slot = x/2 + 1
					break outer
				}
			}

			if x == 0 {
				x++
			} else if x == width-1 {
				x--
			} else if tokens[i][token_i] == 'R' {
				x++
			} else if tokens[i][token_i] == 'L' {
				x--
			}

			token_i++
		}

		score := final_slot*2 - toss_slot
		if score > 0 {
			total_score += score
		}

	}

	return total_score
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2", 1, 1

	data := loader.GetStrings()
	grid, tokens := parse_data(data)
	part1 := play_tokens(grid, tokens)

	fmt.Printf("%d\n", part1)
}
