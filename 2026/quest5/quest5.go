package quest5

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

var height, width int

func in_bounds(x int, y int) bool {
	return !(x < 0 || x >= width || y < 0 || y >= height)
}

func get_neighbours(grid [][]byte, pos Pos) []Pos {
	neighbours := []Pos{}
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			n_x := pos.x + dx
			n_y := pos.y + dy
			if in_bounds(n_x, n_y) && grid[n_y][n_x] != '#' {
				neighbours = append(neighbours, Pos{x: n_x, y: n_y})
			}
		}
	}
	return neighbours
}

func land_count(grid [][]byte, pos Pos) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			n_x := pos.x + dx
			n_y := pos.y + dy
			if in_bounds(n_x, n_y) && (grid[n_y][n_x] == '1' || grid[n_y][n_x] == '#') {
				count++
			}
		}
	}
	return count
}

func reconstruct(grid [][]byte) int {
	height = len(grid)
	width = len(grid[0])

	start := Pos{x: width / 2, y: height / 2}
	grid[start.y][start.x] = '#'
	queue := []Pos{start}

	land := 1
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		neighbours := get_neighbours(grid, pos)
		updated := true
		for updated == true {
			updated = false
			for _, neighbour := range neighbours {
				if grid[neighbour.y][neighbour.x] == '#' {
					continue
				}
				if land_count(grid, neighbour) >= 5 {
					grid[neighbour.y][neighbour.x] = '#'
					land++
					queue = append(queue, neighbour)
					updated = true
				}
			}
		}
	}

	/*
		for _, r := range grid {
			fmt.Println(string(r))
		}
	*/

	return land
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 5, 1

	grid := loader.GetByteRows()
	part1 := reconstruct(grid)

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
