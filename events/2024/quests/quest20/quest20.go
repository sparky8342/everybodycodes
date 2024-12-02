package quest20

import (
	"fmt"
	"loader"
)

type Glider struct {
	x        int
	y        int
	dir      byte
	altitude int
	time     int
}

var height, width int

func dfs(grid []string, glider Glider, visited map[Glider]struct{}, max *int) {
	if glider.time == 100 {
		if glider.altitude > *max {
			*max = glider.altitude
		}
		return
	}

	var directions []byte
	switch glider.dir {
	case 'U':
		directions = []byte{'U', 'L', 'R'}
	case 'D':
		directions = []byte{'D', 'L', 'R'}
	case 'L':
		directions = []byte{'L', 'U', 'D'}
	case 'R':
		directions = []byte{'R', 'U', 'D'}
	}

	for _, dir := range directions {
		new_x := glider.x
		new_y := glider.y

		switch dir {
		case 'U':
			new_y--
		case 'D':
			new_y++
		case 'L':
			new_x--
		case 'R':
			new_x++
		}

		if new_x < 0 || new_x == width || new_y < 0 || new_y == height {
			continue
		}
		if grid[new_y][new_x] == '#' {
			continue
		}

		altitude := glider.altitude
		switch grid[new_y][new_x] {
		case '.':
			altitude--
		case '-':
			altitude -= 2
		case '+':
			altitude++
		}

		new_glider := Glider{
			x:        new_x,
			y:        new_y,
			dir:      dir,
			altitude: altitude,
			time:     glider.time + 1,
		}

		if _, ok := visited[new_glider]; !ok {
			visited[new_glider] = struct{}{}
			dfs(grid, new_glider, visited, max)
		}
	}
}

func highest_altitude(grid []string) int {
	height = len(grid)
	width = len(grid[0])

	start_x := width / 2

	starts := []Glider{
		Glider{x: start_x, y: 0, dir: 'L', altitude: 1000, time: 0},
		Glider{x: start_x, y: 0, dir: 'R', altitude: 1000, time: 0},
		Glider{x: start_x, y: 0, dir: 'D', altitude: 1000, time: 0},
	}

	visited := map[Glider]struct{}{}
	max := 0

	for _, glider := range starts {
		dfs(grid, glider, visited, &max)
	}

	return max

}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 20, 1

	grid := loader.GetStrings()
	part1 := highest_altitude(grid)

	part2, part3 := -1, -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
