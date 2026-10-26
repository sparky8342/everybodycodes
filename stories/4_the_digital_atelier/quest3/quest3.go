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

type Grid struct {
	squares [][]uint8
	height int
	width int
	h_offsets string
	v_offsets string
	rect_width int
	rect_height int
	extra_width int
	extra_height int
}

func parse_data(data []string) *Grid {
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
	return &Grid{
		width : width,
		height : height,
		h_offsets : h_offsets,
		v_offsets : v_offsets,
	}
}

func draw_pattern(grid *Grid) {
	grid.rect_width = len(grid.v_offsets) * 2
	if grid.rect_width > grid.width {
		grid.rect_width /= 2
	}
	grid.rect_height = len(grid.h_offsets) * 2
	if grid.rect_height > grid.height {
		grid.rect_height /= 2
	}
	fmt.Println(grid)

	grid.squares = make([][]uint8, grid.rect_height + 1)
	for i := range grid.squares {
		grid.squares[i] = make([]uint8, grid.rect_width + 1)
	}

	for y := 0; y <= grid.rect_height; y++ {
		if grid.h_offsets[y%len(grid.h_offsets)] == '1' {
			for x := 1; x <= grid.rect_width; x += 2 {
				grid.squares[y][x] |= TOP
			}
		} else {
			for x := 0; x <= grid.rect_width; x += 2 {
				grid.squares[y][x] |= TOP
			}
		}
	}

	for x := 0; x <= grid.rect_width; x++ {
		if grid.v_offsets[x%len(grid.v_offsets)] == '1' {
			for y := 1; y <= grid.rect_height; y += 2 {
				grid.squares[y][x] |= LEFT
			}
		} else {
			for y := 0; y <= grid.rect_height; y += 2 {
				grid.squares[y][x] |= LEFT
			}
		}
	}
}

func count_isolated(grid *Grid) int {
	w_mult := grid.width / grid.rect_width
	h_mult := grid.height / grid.rect_height
	fmt.Println("mults", w_mult, h_mult)
	fmt.Println(grid.width, grid.height, grid.rect_width, grid.rect_height)

	grid.extra_width = grid.width % grid.rect_width
	grid.extra_height = grid.height % grid.rect_height

	fmt.Println("extra", grid.extra_width, grid.extra_height)

	isolated := 0
	colour1 := 0
	colour2 := 0

	isolated_extra := 0
	colour1_extra := 0
	colour2_extra := 0
	for y := 0; y < grid.rect_height; y++ {
		for x := 0; x < grid.rect_width; x++ {
			if grid.squares[y][x]&(TOP|LEFT) == (TOP | LEFT) {
				if grid.squares[y+1][x]&TOP == TOP && grid.squares[y][x+1]&LEFT == LEFT {
					if grid.squares[y][x]&COLOUR1 == COLOUR1 {
						colour1++
					} else if grid.squares[y][x]&COLOUR2 == COLOUR2 {
						colour2++
					} else {
						isolated++
					}

					add := 0
					if y < grid.extra_height && x < grid.extra_width {
						add = 1
					} else if y < grid.extra_height || x < grid.extra_width {
						//fmt.Println("adding at", x, y, "extras: ", grid.extra_height, grid.extra_width)
						add = 1
					}

					/*
					if y < grid.extra_height && x < grid.extra_width {
						add = 1
					} else if y < grid.extra_height {
						add = w_mult
					} else if x < grid.extra_width {
						add = h_mult
					}
					*/
			
					if add > 0 {
						if grid.squares[y][x]&COLOUR1 == COLOUR1 {
							colour1_extra += add
						} else if grid.squares[y][x]&COLOUR2 == COLOUR2 {
							colour2_extra += add
						} else {
							isolated_extra += add
						}
					}
				}
			}
		}
	}

	if grid.extra_height > 0 {
		for x := 0; x < grid.rect_width; x++ {
			if grid.squares[grid.rect_height-1][x]&(TOP|LEFT) == (TOP | LEFT) {
				if grid.squares[0][x]&TOP == TOP && grid.squares[grid.rect_height-1][x+1]&LEFT == LEFT {
					if grid.squares[grid.rect_height-1][x]&COLOUR1 == COLOUR1 {
						colour1_extra++
					} else if grid.squares[grid.rect_height-1][x]&COLOUR2 == COLOUR2 {
						colour2_extra++
					} else {
						isolated_extra++
					}
				}
			}
		}
	}
	if grid.extra_width > 0 {
		for y := 0; y < grid.rect_height; y++ {
			if grid.squares[y][grid.rect_width-1]&(TOP|LEFT) == (TOP | LEFT) {
				if grid.squares[y][grid.rect_width-1]&TOP == TOP && grid.squares[y][0]&LEFT == LEFT {
					if grid.squares[y][grid.rect_width-1]&COLOUR1 == COLOUR1 {
						colour1_extra++
					} else if grid.squares[y][grid.rect_width-1]&COLOUR2 == COLOUR2 {
						colour2_extra++
					} else {
						isolated_extra++
					}
				}
			}
		}
	}



	counts := []int{
		isolated * w_mult * h_mult + isolated_extra,
		colour1 * w_mult * h_mult + colour1_extra,
		colour2 * w_mult * h_mult + colour2_extra,
	}
	sort.Ints(counts)
	fmt.Println(counts, isolated_extra, colour1_extra, colour2_extra)
	return counts[2]
}

func fill_row(grid *Grid, y int) {
	for x := 1; x <= grid.rect_width; x++ {
		if grid.squares[y][x-1]&COLOUR1 == COLOUR1 {
			if grid.squares[y][x]&LEFT == 0 {
				grid.squares[y][x] |= COLOUR1
			} else {
				grid.squares[y][x] |= COLOUR2
			}
		} else if grid.squares[y][x-1]&COLOUR2 == COLOUR2 {
			if grid.squares[y][x]&LEFT == 0 {
				grid.squares[y][x] |= COLOUR2
			} else {
				grid.squares[y][x] |= COLOUR1
			}
		}
	}
}

func colour_in(grid *Grid) {
	grid.squares[0][0] |= COLOUR1
	fill_row(grid, 0)

	for y := 1; y < grid.rect_height; y++ {
		if grid.squares[y-1][0]&COLOUR1 == COLOUR1 {
			if grid.squares[y][0]&TOP == 0 {
				grid.squares[y][0] |= COLOUR1
			} else {
				grid.squares[y][0] |= COLOUR2
			}
		} else if grid.squares[y-1][0]&COLOUR2 == COLOUR2 {
			if grid.squares[y][0]&TOP == 0 {
				grid.squares[y][0] |= COLOUR2
			} else {
				grid.squares[y][0] |= COLOUR1
			}
		}
		fill_row(grid, y)
	}
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "4", 3, 1

	data := loader.GetStrings()
	grid := parse_data(data)
	draw_pattern(grid)
	part1 := count_isolated(grid)

	loader.Part = 2
	data = loader.GetStrings()
	grid = parse_data(data)
	draw_pattern(grid)
	colour_in(grid)
	part2 := count_isolated(grid)

	part3 := -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
