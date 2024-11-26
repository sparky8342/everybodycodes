package quest15

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

type Node struct {
	pos           Pos
	letters_found [26]bool
	found_count   int
	distance      int
}

type State struct {
	pos     Pos
	letters [26]bool
}

func find_path(grid []string) int {
	x := 0
	for grid[0][x] != '.' {
		x++
	}

	letters := make([]bool, 26)
	letter_count := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			space := grid[y][x]
			if space >= 'A' && space <= 'Z' {
				n := space - 'A'
				if letters[n] == false {
					letters[n] = true
					letter_count++
				}
			}
		}
	}

	visited := map[State]struct{}{}
	visited[State{pos: Pos{x: x, y: 0}, letters: [26]bool{}}] = struct{}{}

	start_pos := Pos{x: x, y: 1}
	visited[State{pos: start_pos, letters: [26]bool{}}] = struct{}{}

	start_node := Node{
		pos:           start_pos,
		letters_found: [26]bool{},
		distance:      1,
	}
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
			if letters_found[n] == false {
				letters_found[n] = true
				found_count++
			}
		}
		if found_count == letter_count && pos == start_pos {
			return node.distance + 1
		}

		neighbours := []Pos{
			Pos{x: pos.x + 1, y: pos.y},
			Pos{x: pos.x - 1, y: pos.y},
			Pos{x: pos.x, y: pos.y + 1},
			Pos{x: pos.x, y: pos.y - 1},
		}

		for _, neighbour := range neighbours {
			if neighbour.y == -1 || grid[neighbour.y][neighbour.x] == '#' || grid[neighbour.y][neighbour.x] == '~' {
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

	return -1
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 15, 1

	grid := loader.GetStrings()
	part1 := find_path(grid)

	loader.Part = 2
	grid = loader.GetStrings()
	part2 := find_path(grid)

	part3 := -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
