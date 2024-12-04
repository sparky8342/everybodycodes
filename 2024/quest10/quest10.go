package quest10

import (
	"fmt"
	"loader"
)

func solve_grid(grid []string) string {
	positions := [4]int{0, 1, 6, 7}
	word := []byte{}
	for y := 2; y <= 5; y++ {
		for x := 2; x <= 5; x++ {
			col_letters := map[byte]struct{}{}
			for _, sy := range positions {
				col_letters[grid[sy][x]] = struct{}{}
			}
			for _, sx := range positions {
				letter := grid[y][sx]
				if _, ok := col_letters[letter]; ok {
					word = append(word, letter)
					break
				}
			}
		}
	}
	return string(word)
}

func print_grid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func solve_and_update_grid(grid [][]byte) ([][]byte, bool) {
	positions := [4]int{0, 1, 6, 7}
	updated := false

	// letters in common
	for y := 2; y <= 5; y++ {
		for x := 2; x <= 5; x++ {
			if grid[y][x] == '.' {
				col_letters := map[byte]struct{}{}
				for _, sy := range positions {
					col_letters[grid[sy][x]] = struct{}{}
				}
				for _, sx := range positions {
					letter := grid[y][sx]
					if _, ok := col_letters[letter]; ok {
						grid[y][x] = letter
						updated = true
						break
					}
				}
			}
		}
	}

	// single dots
	for y := 2; y <= 5; y++ {
		for x := 2; x <= 5; x++ {
			if grid[y][x] == '.' {
				letters := map[byte]int{}
				for sy := 2; sy <= 5; sy++ {
					letters[grid[sy][x]]++
				}
				for sx := 2; sx <= 5; sx++ {
					letters[grid[y][sx]]++
				}

				if letters['.'] > 2 {
					continue
				}
				delete(letters, '.')

				clues := map[byte]int{}
				question_x, question_y := -1, -1
				for _, sy := range positions {
					if grid[sy][x] == '?' {
						question_x, question_y = x, sy
					}
					clues[grid[sy][x]]++
				}
				for _, sx := range positions {
					if grid[y][sx] == '?' {
						question_x, question_y = sx, y
					}
					clues[grid[y][sx]]++
				}

				if val, ok := clues['?']; ok {
					if val > 1 {
						continue
					}
					delete(clues, '?')
				}

				for letter := range letters {
					clues[letter]--
					if clues[letter] == 0 {
						delete(clues, letter)
					}
				}

				if len(clues) == 1 {
					for letter := range clues {
						grid[y][x] = letter
						if question_x != -1 {
							grid[question_y][question_x] = letter
						}
						updated = true
					}
				}
			}
		}
	}

	return grid, updated
}

func solve_multi_grid(grid_strs []string) int {
	height := len(grid_strs)
	width := len(grid_strs[0])

	grid := make([][]byte, height)
	for i := 0; i < height; i++ {
		grid[i] = []byte(grid_strs[i])
	}

	done := false
	for !done {
		done = true

		for start_y := 0; start_y < height-6; start_y += 6 {
			for start_x := 0; start_x < width-6; start_x += 6 {

				small_grid := make([][]byte, 8)
				for y := 0; y < 8; y++ {
					small_grid[y] = grid[start_y+y][start_x : start_x+8]
				}
				small_grid, updated := solve_and_update_grid(small_grid)

				if updated {
					for y := 0; y < 8; y++ {
						for x := 0; x < 8; x++ {
							grid[start_y+y][start_x+x] = small_grid[y][x]
						}
					}
					done = false
				}
			}
		}
	}

	p := 0
	for start_y := 0; start_y < height-6; start_y += 6 {
	outer:
		for start_x := 0; start_x < width-6; start_x += 6 {
			word := []byte{}
			for y := start_y + 2; y <= start_y+5; y++ {
				for x := start_x + 2; x <= start_x+5; x++ {
					if grid[y][x] == '.' {
						continue outer
					}
					word = append(word, grid[y][x])
				}
			}
			p += power(string(word))
		}
	}
	return p
}

func parse_data(data []string) [][]string {
	height := len(data)
	width := len(data[0])

	grids := [][]string{}
	for start_y := 0; start_y < height; start_y += 9 {
		for start_x := 0; start_x < width; start_x += 9 {
			grid := make([]string, 8)
			for i := 0; i < 8; i++ {
				grid[i] = data[start_y+i][start_x : start_x+8]
			}
			grids = append(grids, grid)
		}
	}

	return grids
}

func power(word string) int {
	p := 0
	for i := 0; i < len(word); i++ {
		p = p + (i+1)*int(word[i]-64)
	}
	return p
}

func total_power(grids [][]string) int {
	p := 0
	for _, grid := range grids {
		p += power(solve_grid(grid))
	}
	return p
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 10, 1

	data := loader.GetStrings()
	part1 := solve_grid(data)

	loader.Part = 2
	data = loader.GetStrings()
	grids := parse_data(data)
	part2 := total_power(grids)

	loader.Part = 3
	data = loader.GetStrings()
	part3 := solve_multi_grid(data)

	fmt.Printf("%s %d %d\n", part1, part2, part3)
}
