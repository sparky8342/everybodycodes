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

func in_bounds(x int, y int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func get_edges(grid []string) [][]Pos {
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

func all_energy_needed(grid []string) int {
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
		total_energy += energy_needed(grid, points, dir)
	}

	return total_energy
}

func energy_needed_nw(grid []string) int {
	height = len(grid)
	width = len(grid[0])
	edges := get_edges(grid)
	points := append(edges[0], edges[1][1:]...)
	return energy_needed(grid, points, Dir{dx: 1, dy: 1})
}

func energy_needed(grid []string, points []Pos, dir Dir) int {
	total_energy := 0
	for _, point := range points {
		total_energy += shine(grid, point, dir)
	}
	return total_energy
}

func shine(grid []string, pos Pos, dir Dir) int {
	energy := 0
	shadow := 0
	x, y := pos.x, pos.y
	for in_bounds(x, y) && grid[y][x] != '.' {
		height := int(grid[y][x] - '0')
		if height > shadow {
			energy++
			shadow = height - 1
		} else {
			shadow--
		}
		x += dir.dx
		y += dir.dy
	}
	return energy
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 8, 1

	grid := loader.GetStrings()
	part1 := energy_needed_nw(grid)

	loader.Part = 2
	grid = loader.GetStrings()
	part2 := all_energy_needed(grid)

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
