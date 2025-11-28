package quest20

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
}

func count_pairs(grid []string) (int, map[Pos][]Pos) {
	height := len(grid)
	width := len(grid[0])

	paths := map[Pos][]Pos{}

	pairs := 0
	for y := 0; y < height-1; y++ {
		for x := 0; x < width-1; x++ {
			if grid[y][x] == 'T' || grid[y][x] == 'S' || grid[y][x] == 'E' {
				start := Pos{x: x, y: y}
				if grid[y][x+1] == 'T' || grid[y][x+1] == 'S' || grid[y][x+1] == 'E' {
					end := Pos{x: x + 1, y: y}
					if _, ok := paths[start]; !ok {
						paths[start] = []Pos{}
					}
					paths[start] = append(paths[start], end)
					if _, ok := paths[end]; !ok {
						paths[end] = []Pos{}
					}
					paths[end] = append(paths[end], start)
					pairs++
				}
				if (y%2 != x%2) && (grid[y+1][x] == 'T' || grid[y+1][x] == 'S' || grid[y+1][x] == 'E') {
					end := Pos{x: x, y: y + 1}
					if _, ok := paths[start]; !ok {
						paths[start] = []Pos{}
					}
					paths[start] = append(paths[start], end)
					if _, ok := paths[end]; !ok {
						paths[end] = []Pos{}
					}
					paths[end] = append(paths[end], start)
					pairs++
				}
			}
		}
	}

	return pairs, paths
}

func bfs(grid []string, paths map[Pos][]Pos) int {
	height := len(grid)
	width := len(grid[0])

	start, end := Pos{}, Pos{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'S' {
				start.x, start.y = x, y
			} else if grid[y][x] == 'E' {
				end.x, end.y = x, y
			}
		}
	}

	queue := []Entry{Entry{pos: start}}
	visited := map[Pos]struct{}{}

	for len(queue) > 0 {
		entry := queue[0]
		queue = queue[1:]

		pos := entry.pos

		if pos.x == end.x && pos.y == end.y {
			return entry.distance
		}

		for _, neighbour := range paths[pos] {
			if _, ok := visited[neighbour]; !ok {
				queue = append(queue, Entry{pos: neighbour, distance: entry.distance + 1})
				visited[neighbour] = struct{}{}
			}
		}
	}

	return -1
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 20, 1

	grid := loader.GetStrings()
	part1, _ := count_pairs(grid)

	loader.Part = 2
	grid = loader.GetStrings()
	_, paths := count_pairs(grid)
	part2 := bfs(grid, paths)

	fmt.Printf("%d %d %d\n", part1, part2, 0)
}
