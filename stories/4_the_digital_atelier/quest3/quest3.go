package quest3

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

const (
	TOP  = 1
	LEFT = 2
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

func draw_pattern(width int, height int, h_offsets string, v_offsets string) int {
	width++
	height++

	grid := make([][]uint8, height)
	for i := range grid {
		grid[i] = make([]uint8, width)
	}

	for y := 0; y < height; y++ {
		if h_offsets[y%len(h_offsets)] == '1' {
			for x := 1; x < width; x += 2 {
				grid[y][x] |= TOP
			}
		} else {
			for x := 0; x < width; x += 2 {
				grid[y][x] |= TOP
			}
		}
	}

	for x := 0; x < width; x++ {
		if v_offsets[x%len(v_offsets)] == '1' {
			for y := 1; y < height; y += 2 {
				grid[y][x] |= LEFT
			}
		} else {
			for y := 0; y < height; y += 2 {
				grid[y][x] |= LEFT
			}
		}
	}

	isolated := 0
	for y := 0; y < height-1; y++ {
		for x := 0; x < width-1; x++ {
			if grid[y][x]&(TOP|LEFT) == (TOP | LEFT) {
				if grid[y+1][x]&TOP == TOP && grid[y][x+1]&LEFT == LEFT {
					isolated++
				}
			}
		}
	}

	return isolated
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "4", 3, 1

	data := loader.GetStrings()
	width, height, h_offsets, v_offsets := parse_data(data)
	part1 := draw_pattern(width, height, h_offsets, v_offsets)

	part2, part3 := -1, -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
