package quest3

import (
	"fmt"
	"loader"
	"sort"
	"strconv"
	"strings"
)

const (
	TOP     = 1
	LEFT    = 2
	COLOUR1 = 4
	COLOUR2 = 8
)

func parse_data(data []string) (int, int, string, string) {
	width, err := strconv.Atoi(strings.Split(data[0], "=")[1])
	if err != nil {
		panic(err)
	}
	height, err := strconv.Atoi(strings.Split(data[1], "=")[1])
	if err != nil {
		panic(err)
	}
	h_offsets := strings.Split(data[2], "=")[1]
	v_offsets := strings.Split(data[3], "=")[1]
	return width, height, h_offsets, v_offsets
}

func draw_pattern(width int, height int, h_offsets string, v_offsets string) [][]uint8 {
	grid := make([][]uint8, height+1)
	for i := range grid {
		grid[i] = make([]uint8, width+1)
	}

	for y := 0; y <= height; y++ {
		if h_offsets[y%len(h_offsets)] == '1' {
			for x := 1; x <= width; x += 2 {
				grid[y][x] |= TOP
			}
		} else {
			for x := 0; x <= width; x += 2 {
				grid[y][x] |= TOP
			}
		}
	}

	for x := 0; x <= width; x++ {
		if v_offsets[x%len(v_offsets)] == '1' {
			for y := 1; y <= height; y += 2 {
				grid[y][x] |= LEFT
			}
		} else {
			for y := 0; y <= height; y += 2 {
				grid[y][x] |= LEFT
			}
		}
	}

	return grid
}

func count_isolated(grid [][]uint8, width int, height int) int {
	isolated := 0
	colour1 := 0
	colour2 := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x]&(TOP|LEFT) == (TOP | LEFT) {
				if grid[y+1][x]&TOP == TOP && grid[y][x+1]&LEFT == LEFT {
					if grid[y][x]&COLOUR1 == COLOUR1 {
						colour1++
					} else if grid[y][x]&COLOUR2 == COLOUR2 {
						colour2++
					} else {
						isolated++
					}
				}
			}
		}
	}
	counts := []int{isolated, colour1, colour2}
	sort.Ints(counts)
	return counts[2]
}

func fill_row(grid [][]uint8, width int, height int, y int) {
	for x := 1; x <= width; x++ {
		if grid[y][x-1]&COLOUR1 == COLOUR1 {
			if grid[y][x]&LEFT == 0 {
				grid[y][x] |= COLOUR1
			} else {
				grid[y][x] |= COLOUR2
			}
		} else if grid[y][x-1]&COLOUR2 == COLOUR2 {
			if grid[y][x]&LEFT == 0 {
				grid[y][x] |= COLOUR2
			} else {
				grid[y][x] |= COLOUR1
			}
		}
	}
}

func colour_in(grid [][]uint8, width int, height int) {
	grid[0][0] |= COLOUR1
	fill_row(grid, width, height, 0)

	for y := 1; y < height; y++ {
		if grid[y-1][0]&COLOUR1 == COLOUR1 {
			if grid[y][0]&TOP == 0 {
				grid[y][0] |= COLOUR1
			} else {
				grid[y][0] |= COLOUR2
			}
		} else if grid[y-1][0]&COLOUR2 == COLOUR2 {
			if grid[y][0]&TOP == 0 {
				grid[y][0] |= COLOUR2
			} else {
				grid[y][0] |= COLOUR1
			}
		}
		fill_row(grid, width, height, y)
	}
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "4", 3, 1

	data := loader.GetStrings()
	width, height, h_offsets, v_offsets := parse_data(data)
	grid := draw_pattern(width, height, h_offsets, v_offsets)
	part1 := count_isolated(grid, width, height)

	loader.Part = 2
	data = loader.GetStrings()
	width, height, h_offsets, v_offsets = parse_data(data)
	grid = draw_pattern(width, height, h_offsets, v_offsets)
	colour_in(grid, width, height)
	part2 := count_isolated(grid, width, height)

	part3 := -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
