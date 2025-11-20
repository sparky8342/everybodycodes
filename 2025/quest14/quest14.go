package quest14

import (
	"fmt"
	"loader"
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

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 14, 1

	grid := loader.GetStrings()
	part1 := steps(grid, 10)

	loader.Part = 2
	grid = loader.GetStrings()
	part2 := steps(grid, 2025)

	fmt.Printf("%d %d %d\n", part1, part2, 0)
}
