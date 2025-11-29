package quest20

import (
	"fmt"
	"loader"
)

type Pos struct {
	x        int
	y        int
	rotation int
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

func rotate_grid(grid []string) []string {
	height := len(grid)
	width := len(grid[0])

	r := make([][]byte, height)
	for y := range r {
		r[y] = make([]byte, width)
		for x := 0; x < width; x++ {
			r[y][x] = '.'
		}
	}

	start_x, start_y := width-1, 0
	for y := 0; y < height; y++ {
		rx, ry := start_x, start_y
		x := 0
		for grid[y][x] == '.' {
			x++
		}
		for ; x < width; x += 2 {
			if grid[y][x] == '.' {
				break
			}
			r[ry][rx] = grid[y][x]
			ry++
			rx--
		}
		start_x--
		rx, ry = start_x, start_y
		x = 0
		for grid[y][x] == '.' {
			x++
		}
		x++
		for ; x < width; x += 2 {
			if grid[y][x] == '.' {
				break
			}
			r[ry][rx] = grid[y][x]
			ry++
			rx--
		}
		start_x--
	}

	rotated := make([]string, height)
	for y := 0; y < height; y++ {
		rotated[y] = string(r[y])
	}

	return rotated
}

func bfs_rotate(rotate0 []string) int {
	height := len(rotate0)
	width := len(rotate0[0])

	rotate1 := rotate_grid(rotate0)
	rotate2 := rotate_grid(rotate1)

	paths := map[Pos][]Pos{}
	rotation := 0

	for _, pair := range [][][]string{[][]string{rotate0, rotate1}, [][]string{rotate1, rotate2}, [][]string{rotate2, rotate0}} {
		from, to := pair[0], pair[1]

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if from[y][x] == 'T' || from[y][x] == 'S' || from[y][x] == 'E' {
					start := Pos{x: x, y: y, rotation: rotation}
					if to[y][x] == 'T' || to[y][x] == 'S' || to[y][x] == 'E' {
						end := Pos{x: x, y: y, rotation: (rotation + 1) % 3}
						if _, ok := paths[start]; !ok {
							paths[start] = []Pos{}
						}
						paths[start] = append(paths[start], end)
					}
					if x < width-1 && (to[y][x+1] == 'T' || to[y][x+1] == 'S' || to[y][x+1] == 'E') {
						end := Pos{x: x + 1, y: y, rotation: (rotation + 1) % 3}
						if _, ok := paths[start]; !ok {
							paths[start] = []Pos{}
						}
						paths[start] = append(paths[start], end)
					}
					if x > 0 && (to[y][x-1] == 'T' || to[y][x-1] == 'S' || to[y][x-1] == 'E') {
						end := Pos{x: x - 1, y: y, rotation: (rotation + 1) % 3}
						if _, ok := paths[start]; !ok {
							paths[start] = []Pos{}
						}
						paths[start] = append(paths[start], end)
					}
					if y < height-1 && (y%2 != x%2) && (to[y+1][x] == 'T' || to[y+1][x] == 'S' || to[y+1][x] == 'E') {
						end := Pos{x: x, y: y + 1, rotation: (rotation + 1) % 3}
						if _, ok := paths[start]; !ok {
							paths[start] = []Pos{}
						}
						paths[start] = append(paths[start], end)
					}
					if y > 0 && (y%2 == x%2) && (to[y-1][x] == 'T' || to[y-1][x] == 'S' || to[y-1][x] == 'E') {
						end := Pos{x: x, y: y - 1, rotation: (rotation + 1) % 3}
						if _, ok := paths[start]; !ok {
							paths[start] = []Pos{}
						}
						paths[start] = append(paths[start], end)
					}

				}
			}
		}

		rotation++

	}

	start, end := Pos{}, Pos{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if rotate0[y][x] == 'S' {
				start.x, start.y = x, y
			} else if rotate0[y][x] == 'E' {
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

		if pos.rotation == 0 && pos.x == end.x && pos.y == end.y {
			return entry.distance - 1
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

	loader.Part = 3
	grid = loader.GetStrings()
	part3 := bfs_rotate(grid)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
