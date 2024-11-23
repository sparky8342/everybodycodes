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

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 10, 1

	data := loader.GetStrings()
	part1 := solve_grid(data)

	part2, part3 := -1, -1
	fmt.Printf("%s %d %d\n", part1, part2, part3)
}
