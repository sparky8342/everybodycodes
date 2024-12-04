package quest13

import (
	"fmt"
	"loader"
	"math"
)

type Pos struct {
	x int
	y int
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func get_dist(grid []string, a Pos, b Pos) int {
	a_byt := grid[a.y][a.x]
	b_byt := grid[b.y][b.x]
	var a_val, b_val int
	if a_byt == 'S' || a_byt == 'E' {
		a_val = 0
	} else {
		a_val = int(a_byt - '0')
	}
	if b_byt == 'S' || b_byt == 'E' {
		b_val = 0
	} else {
		b_val = int(b_byt - '0')
	}

	dist := abs(a_val - b_val)
	if dist > 5 {
		dist = 10 - dist
	}
	return dist + 1
}

func find_path(grid []string, start_b byte, end_b byte) int {
	height := len(grid)
	width := len(grid[0])

	nodes := map[Pos]int{}
	var start Pos

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			space := grid[y][x]
			if space == start_b {
				start = Pos{x: x, y: y}
				nodes[start] = 0
			} else if space == end_b || (space >= '0' && space <= '9') {
				nodes[Pos{x: x, y: y}] = math.MaxInt32
			}
		}
	}

	for {
		// TODO - use priority queue
		var node Pos
		smallest_distance := math.MaxInt32
		for n, distance := range nodes {
			if distance < smallest_distance {
				smallest_distance = distance
				node = n
			}
		}

		if grid[node.y][node.x] == end_b {
			return smallest_distance
		}

		neighbours := []Pos{
			Pos{x: node.x + 1, y: node.y},
			Pos{x: node.x - 1, y: node.y},
			Pos{x: node.x, y: node.y + 1},
			Pos{x: node.x, y: node.y - 1},
		}

		for _, neighbour := range neighbours {
			if _, ok := nodes[neighbour]; ok {
				dist := get_dist(grid, node, neighbour)
				nodes[neighbour] = min(nodes[neighbour], smallest_distance+dist)
			}
		}

		delete(nodes, node)
	}
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 13, 1

	grid := loader.GetStrings()
	part1 := find_path(grid, 'S', 'E')

	loader.Part = 2
	grid = loader.GetStrings()
	part2 := find_path(grid, 'S', 'E')

	loader.Part = 3
	grid = loader.GetStrings()
	part3 := find_path(grid, 'E', 'S')

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
