package quest18

import (
	"fmt"
	"loader"
	"math"
)

type Pos struct {
	x int
	y int
}

type State struct {
	pos  Pos
	time int
}

func fill(grid []string) int {
	height := len(grid)
	width := len(grid[0])

	palms := 0
	starts := []Pos{}
	for y := 0; y < height; y++ {
		if grid[y][0] == '.' {
			starts = append(starts, Pos{x: 0, y: y})
		}
		if grid[y][width-1] == '.' {
			starts = append(starts, Pos{x: width - 1, y: y})
		}
		for x := 0; x < width; x++ {
			if grid[y][x] == 'P' {
				palms++
			}
		}
	}

	time, _ := bfs(grid, starts, palms, math.MaxInt32)
	return time
}

func bfs(grid []string, starts []Pos, palms int, max_palms_time int) (int, int) {
	width := len(grid[0])

	queue := []State{}
	visited := map[Pos]struct{}{}
	for _, start := range starts {
		queue = append(queue, State{pos: start})
		visited[start] = struct{}{}
	}
	palms_reached := 0
	palms_time := 0

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		pos := state.pos

		if grid[pos.y][pos.x] == 'P' {
			palms_reached++
			palms_time += state.time
			if palms_reached == palms {
				return state.time, palms_time
			}
			if palms_time >= max_palms_time {
				return -1, math.MaxInt32
			}
		}

		neighbours := []Pos{
			Pos{x: pos.x + 1, y: pos.y},
			Pos{x: pos.x - 1, y: pos.y},
			Pos{x: pos.x, y: pos.y + 1},
			Pos{x: pos.x, y: pos.y - 1},
		}

		for _, neighbour := range neighbours {
			if _, ok := visited[neighbour]; ok {
				continue
			}
			if neighbour.x == -1 || neighbour.x == width || grid[neighbour.y][neighbour.x] == '#' {
				continue
			}
			queue = append(queue, State{pos: neighbour, time: state.time + 1})
			visited[neighbour] = struct{}{}
		}
	}

	return -1, -1
}

func find_start(grid []string) int {
	// just brute force all starting points

	height := len(grid)
	width := len(grid[0])

	palms := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'P' {
				palms++
			}
		}
	}

	min := math.MaxInt32
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '.' {
				start := Pos{x: x, y: y}
				_, time := bfs(grid, []Pos{start}, palms, min)
				if time < min {
					min = time
				}
			}
		}
	}

	return min
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 18, 1

	grid := loader.GetStrings()
	part1 := fill(grid)

	loader.Part = 2
	grid = loader.GetStrings()
	part2 := fill(grid)

	loader.Part = 3
	grid = loader.GetStrings()
	part3 := find_start(grid)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
