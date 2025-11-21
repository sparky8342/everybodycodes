package quest14

import (
	"fmt"
	"loader"
	"strings"
)

var width, height int

var dirs [4][2]int

func init() {
	dirs = [4][2]int{
		[2]int{1, 1},
		[2]int{1, -1},
		[2]int{-1, 1},
		[2]int{-1, -1},
	}
}

func active_neighbours(grid [][]byte, x int, y int) int {
	count := 0
	for _, dir := range dirs {
		nx := x + dir[0]
		ny := y + dir[1]
		if nx < 0 || nx == width || ny < 0 || ny == height {
			continue
		}
		if grid[ny][nx] == '#' {
			count++
		}
	}
	return count
}

func step(grid [][]byte) [][]byte {
	next := make([][]byte, height)

	for y := 0; y < height; y++ {
		next[y] = make([]byte, width)

		for x := 0; x < width; x++ {
			n := active_neighbours(grid, x, y)
			if grid[y][x] == '#' {
				if n%2 == 1 {
					next[y][x] = '#'
				} else {
					next[y][x] = '.'
				}
			} else if grid[y][x] == '.' {
				if n%2 == 0 {
					next[y][x] = '#'
				} else {
					next[y][x] = '.'
				}
			}
		}
	}

	return next
}

func count_active(grid [][]byte) int {
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '#' {
				count++
			}
		}
	}
	return count
}

func steps(grid_str []string, amount int) int {
	height = len(grid_str)
	width = len(grid_str[0])

	grid := make([][]byte, height)
	for i := 0; i < height; i++ {
		grid[i] = []byte(grid_str[i])
	}

	total := 0
	for i := 0; i < amount; i++ {
		grid = step(grid)
		total += count_active(grid)
	}

	return total
}

func matches(grid [][]byte, match []string) bool {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if grid[y+13][x+13] != match[y][x] {
				return false
			}
		}
	}
	return true
}

func cache_key(grid [][]byte) string {
	strs := make([]string, height)
	for y := 0; y < height; y++ {
		strs[y] = string(grid[y])
	}
	return strings.Join(strs, "")
}

func steps_matching(match []string, amount int) int {
	height = 34
	width = 34

	grid := make([][]byte, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]byte, width)
		for x := 0; x < width; x++ {
			grid[y][x] = '.'
		}
	}

	first_match := ""
	sequence_start := 0
	sequence_end := 0

	total := 0
	for i := 0; i < amount; i++ {
		grid = step(grid)
		if matches(grid, match) {
			key := cache_key(grid)
			if first_match == "" {
				first_match = key
				sequence_start = i
			} else if first_match == key {
				sequence_end = i
				break
			}
			total += count_active(grid)
		}
	}

	sequence := sequence_end - sequence_start
	left := amount - sequence_end
	repeats := left / sequence
	total *= repeats + 1
	from := sequence_end + repeats*sequence

	total += count_active(grid)
	for i := from; i < amount; i++ {
		grid = step(grid)
		if matches(grid, match) {
			total += count_active(grid)
		}
	}

	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 14, 1

	grid := loader.GetStrings()
	part1 := steps(grid, 10)

	loader.Part = 2
	grid = loader.GetStrings()
	part2 := steps(grid, 2025)

	loader.Part = 3
	grid = loader.GetStrings()
	part3 := steps_matching(grid, 1000000000)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
