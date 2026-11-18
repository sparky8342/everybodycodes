package quest17

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

type Entry struct {
	pos      Pos
	distance int
	prev     *Entry
}

type Dir struct {
	dx int
	dy int
}

type Edge struct {
	a Pos
	b Pos
}

var dirs [4]Dir

func init() {
	dirs = [4]Dir{
		Dir{dx: 1, dy: 0},
		Dir{dx: -1, dy: 0},
		Dir{dx: 0, dy: 1},
		Dir{dx: 0, dy: -1},
	}
}

func shortest_path(grid []string) int {
	height := len(grid)
	width := len(grid[0])

	start := Pos{}
	end := Pos{x: width - 1, y: height - 1}

	queue := []Entry{{pos: start}}
	visited := map[Pos]struct{}{start: struct{}{}}

	for len(queue) > 0 {
		entry := queue[0]
		queue = queue[1:]

		pos := entry.pos
		if pos == end {
			return entry.distance
		}

		for _, dir := range dirs {
			new_pos := Pos{x: pos.x + dir.dx, y: pos.y + dir.dy}
			if new_pos.x < 0 || new_pos.x == width || new_pos.y < 0 || new_pos.y == height {
				continue
			}
			if grid[new_pos.y][new_pos.x] == '#' {
				if _, ok := visited[new_pos]; !ok {
					queue = append(queue, Entry{pos: new_pos, distance: entry.distance + 1})
					visited[new_pos] = struct{}{}
				}
			}
		}
	}

	return -1
}

func find_connected_edges(path map[Edge]struct{}, edges map[Edge]struct{}) bool {
	found := false
	for edge := range path {
		for edge2 := range edges {
			if edge.a == edge2.a || edge.a == edge2.b || edge.b == edge2.a || edge.b == edge2.b {
				path[edge2] = struct{}{}
				delete(edges, edge2)
				found = true
				break
			}
		}
	}
	return found
}

func edge_distance(grid []string) (int, int) {
	height := len(grid)
	width := len(grid[0])

	distance := 0

	edges := map[Edge]struct{}{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '#' {
				if x == 0 {
					distance++
					a := Pos{x: x, y: y}
					b := Pos{x: x, y: y + 1}
					edges[Edge{a: a, b: b}] = struct{}{}
				} else if x == width-1 {
					distance++
					a := Pos{x: x + 1, y: y}
					b := Pos{x: x + 1, y: y + 1}
					edges[Edge{a: a, b: b}] = struct{}{}
				}

				if y == 0 {
					distance++
					a := Pos{x: x, y: y}
					b := Pos{x: x + 1, y: y}
					edges[Edge{a: a, b: b}] = struct{}{}
				} else if y == height-1 {
					distance++
					a := Pos{x: x, y: y + 1}
					b := Pos{x: x + 1, y: y + 1}
					edges[Edge{a: a, b: b}] = struct{}{}
				}
			}
			if x > 0 {
				if grid[y][x] != grid[y][x-1] {
					distance++
					a := Pos{x: x, y: y}
					b := Pos{x: x, y: y + 1}
					edges[Edge{a: a, b: b}] = struct{}{}
				}
			}
			if y > 0 {
				if grid[y][x] != grid[y-1][x] {
					distance++
					a := Pos{x: x, y: y}
					b := Pos{x: x + 1, y: y}
					edges[Edge{a: a, b: b}] = struct{}{}
				}
			}
		}
	}

	top_path := map[Edge]struct{}{}
	left_path := map[Edge]struct{}{}

	top_start := Edge{a: Pos{x: 0, y: 0}, b: Pos{x: 1, y: 0}}
	top_end := Edge{a: Pos{x: width, y: height - 1}, b: Pos{x: width, y: height}}
	left_start := Edge{a: Pos{x: 0, y: 0}, b: Pos{x: 0, y: 1}}
	left_end := Edge{a: Pos{x: width - 1, y: height}, b: Pos{x: width, y: height}}

	top_path[top_start] = struct{}{}
	delete(edges, top_start)
	top_path[top_end] = struct{}{}
	delete(edges, top_end)

	left_path[left_start] = struct{}{}
	delete(edges, left_start)
	left_path[left_end] = struct{}{}
	delete(edges, left_end)

	for {
		found1 := find_connected_edges(top_path, edges)
		found2 := find_connected_edges(left_path, edges)
		if !found1 && !found2 {
			break
		}
	}

	var path_length int
	if len(top_path) < len(left_path) {
		path_length = len(top_path)
	} else {
		path_length = len(left_path)
	}

	return distance, path_length
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 17, 1

	grid := loader.GetStrings()
	part1 := shortest_path(grid)

	loader.Part = 2
	grid = loader.GetStrings()
	part2, _ := edge_distance(grid)

	loader.Part = 3
	grid = loader.GetStrings()
	_, part3 := edge_distance(grid)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
