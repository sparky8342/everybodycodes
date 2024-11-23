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

	part3 := -1
	fmt.Printf("%s %d %d\n", part1, part2, part3)
}
