package quest8

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

type Dir struct {
	dx int
	dy int
}

var width, height int

func parse_grid(data []string) [][]byte {
	grid := [][]byte{}
	for _, row := range data {
		grid = append(grid, []byte(row))
	}
	return grid
}

func in_bounds(x int, y int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func get_edges(grid [][]byte) [][]Pos {
	edge_len := (height + 1) / 2
	edges := [][]Pos{}

	dirs := []Dir{
		Dir{dx: +1, dy: -1},
		Dir{dx: +2, dy: +0},
		Dir{dx: +1, dy: +1},
		Dir{dx: -1, dy: +1},
		Dir{dx: -2, dy: +0},
		Dir{dx: -1, dy: -1},
	}

	x := 0
	y := height / 2

	for _, dir := range dirs {
		edge := []Pos{Pos{x: x, y: y}}

		for i := 0; i < edge_len-1; i++ {
			x += dir.dx
			y += dir.dy
			edge = append(edge, Pos{x: x, y: y})
		}
		edges = append(edges, edge)
	}

	return edges
}

func all_energy_needed(grid [][]byte) int {
	height = len(grid)
	width = len(grid[0])

	edges := get_edges(grid)

	dirs := []Dir{
		Dir{dx: +1, dy: +1},
		Dir{dx: -1, dy: +1},
		Dir{dx: -2, dy: +0},
		Dir{dx: -1, dy: -1},
		Dir{dx: +1, dy: -1},
		Dir{dx: +2, dy: +0},
	}

	total_energy := 0

	for i := 0; i < 6; i++ {
		points := append(edges[i], edges[(i+1)%6][1:]...)
		dir := dirs[i]
		energy, _ := energy_needed(grid, points, dir, false)
		total_energy += energy
	}

	return total_energy
}

func energy_needed_nw(grid [][]byte) int {
	height = len(grid)
	width = len(grid[0])
	edges := get_edges(grid)
	points := append(edges[0], edges[1][1:]...)
	energy, _ := energy_needed(grid, points, Dir{dx: 1, dy: 1}, false)
	return energy
}

func energy_needed(grid [][]byte, points []Pos, dir Dir, grow bool) (int, int) {
	total_energy := 0
	total_ents := 0
	for _, point := range points {
		energy, new_ents := shine(grid, point, dir, grow)
		total_energy += energy
		total_ents += new_ents
	}
	return total_energy, total_ents
}

func shine(grid [][]byte, pos Pos, dir Dir, grow bool) (int, int) {
	energy := 0
	shadow := 0
	new_ents := 0
	x, y := pos.x, pos.y
	for in_bounds(x, y) && grid[y][x] != '.' {
		if grid[y][x] >= 'a' && grid[y][x] <= 'c' {
			grid[y][x]++
			if grid[y][x] == 'd' {
				grid[y][x] = '1'
			}
			shadow--
		} else {
			height := int(grid[y][x] - '0')
			if height > shadow {
				energy++
				shadow = height - 1
				if grow {
					if height == 9 {
						grid[y][x] = 'a' // seed goes a, b, c, 1
						new_ents++
					} else {
						grid[y][x]++
					}
				}
			} else {
				shadow--
			}
		}
		x += dir.dx
		y += dir.dy
	}
	return energy, new_ents
}

func grow_ents(grid [][]byte) int {
	height = len(grid)
	width = len(grid[0])

	edges := get_edges(grid)

	dirs := []Dir{
		Dir{dx: +1, dy: +1},
		Dir{dx: -1, dy: +1},
		Dir{dx: -2, dy: +0},
		Dir{dx: -1, dy: -1},
		Dir{dx: +1, dy: -1},
		Dir{dx: +2, dy: +0},
	}

	total_energy := 0
	total_ents := 0

	for {
		for i := 0; i < 6; i++ {
			points := append(edges[i], edges[(i+1)%6][1:]...)
			dir := dirs[i]
			energy, new_ents := energy_needed(grid, points, dir, true)

			total_energy += energy
			total_ents += new_ents
			if total_ents >= 100000 {
				return total_energy
			}
		}
	}

	return -1
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 8, 1

	data := loader.GetStrings()
	grid := parse_grid(data)
	part1 := energy_needed_nw(grid)

	loader.Part = 2
	data = loader.GetStrings()
	grid = parse_grid(data)
	part2 := all_energy_needed(grid)

	loader.Part = 3
	data = loader.GetStrings()
	grid = parse_grid(data)
	part3 := grow_ents(grid)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
