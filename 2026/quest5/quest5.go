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

var stamp [][]byte

var grid [][]byte

func init() {
	stamp = [][]byte{
		[]byte(".+++."),
		[]byte("+++++"),
		[]byte("++X++"),
		[]byte("+++++"),
		[]byte(".+++."),
	}
}

func in_bounds(x int, y int) bool {
	return !(x < 0 || x >= width || y < 0 || y >= height)
}

func on_edge(x int, y int) bool {
	return x == 0 || x == width-1 || y == 0 || y == height-1
}

func get_neighbours(pos Pos) []Pos {
	neighbours := []Pos{}
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			n_x := pos.x + dx
			n_y := pos.y + dy
			if in_bounds(n_x, n_y) {
				neighbours = append(neighbours, Pos{x: n_x, y: n_y})
			}
		}
	}
	return neighbours
}

func surrounding_count(pos Pos) (int, int, int) {
	land := 0
	sand := 0
	trees := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			n_x := pos.x + dx
			n_y := pos.y + dy
			if in_bounds(n_x, n_y) {
				if grid[n_y][n_x] == '1' || grid[n_y][n_x] == '#' {
					land++
				} else if grid[n_y][n_x] == '=' {
					sand++
				} else if grid[n_y][n_x] == '@' {
					trees++
				}
			}
		}
	}
	return land, sand, trees
}

func next_to_water(x int, y int) bool {
	if on_edge(x, y) {
		return true
	}

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			n_x := x + dx
			n_y := y + dy
			if grid[n_y][n_x] != '#' {
				return true
			}
		}
	}
	return false
}

func reconstruct() int {
	height = len(grid)
	width = len(grid[0])

	start := Pos{x: width / 2, y: height / 2}
	grid[start.y][start.x] = '#'
	queue := []Pos{start}

	land := 1
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		neighbours := get_neighbours(pos)
		updated := true
		for updated == true {
			updated = false
			for _, neighbour := range neighbours {
				if grid[neighbour.y][neighbour.x] == '#' {
					continue
				}
				land_count, _, _ := surrounding_count(neighbour)
				if land_count >= 5 {
					grid[neighbour.y][neighbour.x] = '#'
					land++
					queue = append(queue, neighbour)
					updated = true
				}
			}
		}
	}

	return land
}

func initial_sand() []Pos {
	sand := []Pos{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '#' && next_to_water(x, y) {
				sand = append(sand, Pos{x: x, y: y})
			}
		}
	}
	for _, pos := range sand {
		grid[pos.y][pos.x] = '='
	}
	return sand
}

func reconstruct_with_sand() int {
	reconstruct()
	sand := initial_sand()

	done := false
	for !done {
		done = true
		for _, pos := range sand {
			neighbours := get_neighbours(pos)
			for _, neighbour := range neighbours {
				if grid[neighbour.y][neighbour.x] == '#' {
					_, sand_count, _ := surrounding_count(neighbour)
					if sand_count >= 5 {
						grid[neighbour.y][neighbour.x] = '='
						sand = append(sand, neighbour)
						done = false
					}
				}
			}
		}
	}

	return len(sand)
}

func stamp_fits(x int, y int) bool {
	for sy := 0; sy < 5; sy++ {
		for sx := 0; sx < 5; sx++ {
			if stamp[sy][sx] == '.' {
				continue
			} else if stamp[sy][sx] == '+' {
				if !(grid[y+sy][x+sx] == '#' || grid[y+sy][x+sx] == '@') {
					return false
				}
			} else if stamp[sy][sx] == 'X' {
				if grid[y+sy][x+sx] != '#' {
					return false
				}
			}
		}
	}
	return true
}

func reconstruct_with_sand_and_trees() int {
	reconstruct_with_sand()

	trees := map[Pos]struct{}{}
	for y := 0; y < height-5; y++ {
		for x := 0; x < width-5; x++ {
			if stamp_fits(x, y) {
				tree := Pos{x: x + 2, y: y + 2}
				trees[tree] = struct{}{}
				grid[tree.y][tree.x] = '@'
			}
		}
	}

	for i := 0; i <= 1; i++ {
		new_trees := map[Pos]struct{}{}
		for tree := range trees {
			neighbours := get_neighbours(tree)
			for _, neighbour := range neighbours {
				_, sand_count, tree_count := surrounding_count(neighbour)
				if tree_count >= 4 && sand_count == 0 {
					new_trees[neighbour] = struct{}{}
				}
			}
		}
		for tree := range new_trees {
			grid[tree.y][tree.x] = '@'
			trees[tree] = struct{}{}
		}
	}

	return len(trees)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 5, 1

	grid = loader.GetByteRows()
	part1 := reconstruct()

	loader.Part = 2
	grid = loader.GetByteRows()
	part2 := reconstruct_with_sand()

	loader.Part = 3
	grid = loader.GetByteRows()
	part3 := reconstruct_with_sand_and_trees()

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
