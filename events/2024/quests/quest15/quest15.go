package quest15

import (
	"fmt"
	"loader"
)

type Pos struct {
	x uint8
	y uint8
}

type Node struct {
	pos           Pos
	letters_found int32
	found_count   uint8
	distance      uint16
}

type State struct {
	pos     Pos
	letters int32
}

func parse_data(data []string) [][]byte {
	grid := make([][]byte, len(data))
	for i, line := range data {
		grid[i] = []byte(line)
	}
	return grid
}

func find_path(grid [][]byte, start_pos Pos) uint16 {
	height := uint8(len(grid))
	width := uint8(len(grid[0]))

	var letters int32 = 0
	var letter_count uint8 = 0
	for y := uint8(0); y < uint8(len(grid)); y++ {
		for x := 0; x < len(grid[0]); x++ {
			space := grid[y][x]
			if space >= 'A' && space <= 'Z' {
				n := space - 'A'
				var bit int32 = 1 << n
				if bit&letters == 0 {
					letters |= bit
					letter_count++
				}
			}
		}
	}

	visited := map[State]struct{}{}
	visited[State{pos: start_pos}] = struct{}{}

	start_node := Node{pos: start_pos}
	queue := []Node{start_node}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		pos := node.pos
		letters_found := node.letters_found
		found_count := node.found_count
		space := grid[pos.y][pos.x]
		if space >= 'A' && space <= 'Z' {
			n := space - 'A'
			var bit int32 = 1 << n
			if bit&letters == bit && bit&letters_found == 0 {
				letters_found |= bit
				found_count++
			}
		}
		if found_count == letter_count && pos == start_pos {
			return node.distance
		}

		neighbours := []Pos{
			Pos{x: pos.x + 1, y: pos.y},
			Pos{x: pos.x - 1, y: pos.y},
			Pos{x: pos.x, y: pos.y + 1},
			Pos{x: pos.x, y: pos.y - 1},
		}

		for _, neighbour := range neighbours {
			if neighbour.x == 255 || neighbour.x == width || neighbour.y == 255 || neighbour.y == height {
				continue
			}
			if grid[neighbour.y][neighbour.x] == '#' || grid[neighbour.y][neighbour.x] == '~' {
				continue
			}
			state := State{pos: neighbour, letters: letters_found}
			if _, ok := visited[state]; ok {
				continue
			}
			queue = append(queue, Node{
				pos:           neighbour,
				letters_found: letters_found,
				found_count:   found_count,
				distance:      node.distance + 1,
			})
			visited[state] = struct{}{}
		}
	}

	return 0
}

func simple_solve(grid [][]byte) uint16 {
	var x uint8 = 0
	for grid[0][x] != '.' {
		x++
	}
	start_pos := Pos{x: x, y: 0}
	return find_path(grid, start_pos)
}

func part3_solve(grid [][]byte) uint16 {
	//input specific solution

	grids := make([][][]byte, 3)
	for i := 0; i < 3; i++ {
		grids[i] = make([][]byte, len(grid))
	}

	for i, row := range grid {
		grids[0][i] = row[0:86]
		grids[1][i] = row[86:169]
		grids[2][i] = row[169:]
	}

	// force both K's to be found
	grids[1][75][0] = 'X'

	dist0 := find_path(grids[0], Pos{x: 85, y: 75})
	dist1 := find_path(grids[1], Pos{x: 41, y: 0})
	dist2 := find_path(grids[2], Pos{x: 0, y: 75})

	return dist0 + dist1 + dist2 + 4
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 15, 1

	data := loader.GetStrings()
	part1 := simple_solve(parse_data(data))

	loader.Part = 2
	data = loader.GetStrings()
	part2 := simple_solve(parse_data(data))

	loader.Part = 3
	data = loader.GetStrings()
	part3 := part3_solve(parse_data(data))
	//part3 := -1
	fmt.Println(part3)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
