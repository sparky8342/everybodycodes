package quest20

import (
	"fmt"
	"loader"
)

type Glider struct {
	x        uint8
	y        uint8
	dir      byte
	altitude int
	time     int
	letters  int
}

type State struct {
	x       uint8
	y       uint8
	dir     byte
	letters int
}

var height, width uint8

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

		if new_x == 255 || new_x == width || new_y == 255 || new_y == height {
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
	height = uint8(len(grid))
	width = uint8(len(grid[0]))

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

func state_key(glider Glider) int {
	return ((((((int(glider.x) * 256) + int(glider.y)) * 256) + int(glider.dir-'A')) * 10) + glider.letters)
}

func shortest_path(grid []string) int {
	height = uint8(len(grid))
	width = uint8(len(grid[0]))

	start_x := width / 2

	start_glider := Glider{x: start_x, y: 0, dir: 'D', altitude: 10000, time: 0}
	queue := []Glider{start_glider}

	visited := map[int]int{}

	for len(queue) > 0 {
		glider := queue[0]
		queue = queue[1:]

		if glider.x == start_x && glider.y == 0 && glider.letters == 7 && glider.altitude >= 10000 {
			return glider.time
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

			if new_x == 255 || new_x == width || new_y == 255 || new_y == height {
				continue
			}
			if grid[new_y][new_x] == '#' {
				continue
			}

			altitude := glider.altitude
			letters := glider.letters
			switch grid[new_y][new_x] {
			case '.':
				altitude--
			case '-':
				altitude -= 2
			case '+':
				altitude++
			case 'A':
				altitude--
				letters |= 1
			case 'B':
				if letters&1 == 0 {
					continue
				}
				altitude--
				letters |= 2
			case 'C':
				if letters&1 == 0 || letters&2 == 0 {
					continue
				}
				altitude--
				letters |= 4
			case 'S':
				altitude--
			}

			new_glider := Glider{
				x:        new_x,
				y:        new_y,
				dir:      dir,
				altitude: altitude,
				time:     glider.time + 1,
				letters:  letters,
			}

			key := state_key(new_glider)
			if val, ok := visited[key]; ok {
				if val >= altitude {
					continue
				}
			}

			queue = append(queue, new_glider)
			visited[key] = altitude

		}
	}

	return -1
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 20, 1

	grid := loader.GetStrings()
	part1 := highest_altitude(grid)

	loader.Part = 2
	grid = loader.GetStrings()
	part2 := shortest_path(grid)

	part3 := -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
