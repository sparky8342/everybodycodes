package quest12

import (
	"fmt"
	"hash/fnv"
	"loader"
	"sort"
	"strconv"
	"utils"
)

type Pos struct {
	x int
	y int
}

type Dir struct {
	dx int
	dy int
}

var height, width int

var dirs [4]Dir

func init() {
	dirs = [4]Dir{
		Dir{dx: +1, dy: +0},
		Dir{dx: -1, dy: +0},
		Dir{dx: +0, dy: +1},
		Dir{dx: +0, dy: -1},
	}
}

func copy_shape(shape map[Pos]struct{}) map[Pos]struct{} {
	cpy := map[Pos]struct{}{}
	for pos := range shape {
		cpy[pos] = struct{}{}
	}
	return cpy
}

func hash_shape(shape map[Pos]struct{}) uint64 {
	positions := []Pos{}
	for pos := range shape {
		positions = append(positions, pos)
	}
	sort.Slice(positions, func(i, j int) bool {
		if positions[i].x == positions[j].x {
			return positions[i].y < positions[j].y
		} else {
			return positions[i].x < positions[j].x
		}
	})
	b := []byte{}
	for _, pos := range positions {
		b = append(b, []byte(strconv.Itoa(pos.x))...)
		b = append(b, '_')
		b = append(b, []byte(strconv.Itoa(pos.y))...)
	}
	h := fnv.New64a()
	h.Write(b)
	return h.Sum64()
}

func find_shape(grid [][]byte, x int, y int) map[Pos]struct{} {
	start := Pos{x: x, y: y}
	queue := []Pos{start}
	visited := map[Pos]struct{}{start: struct{}{}}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			new_pos := Pos{pos.x + dir.dx, pos.y + dir.dy}
			if new_pos.x < 0 || new_pos.x == width || new_pos.y < 0 || new_pos.y == height {
				continue
			}
			if grid[new_pos.y][new_pos.x] == '#' {
				if _, ok := visited[new_pos]; !ok {
					queue = append(queue, new_pos)
					visited[new_pos] = struct{}{}
				}
			}
		}
	}

	return visited
}

func dfs_remove_shape(shape map[Pos]struct{}, x_min int, x_max int, y_min int, y_max int, cache map[uint64]int) int {
	hash := hash_shape(shape)
	if val, ok := cache[hash]; ok {
		return val
	}

	spell_count := 0

	for len(shape) > 0 {
		longest_lines := [][]Pos{}
		longest := 0

		for y := y_min; y <= y_max; y++ {
			for x := x_min; x <= x_max; x++ {
				pos := Pos{x: x, y: y}
				if _, ok := shape[pos]; ok {
					line := []Pos{pos}
					for dx := x + 1; dx <= x_max; dx++ {
						next_pos := Pos{x: dx, y: y}
						if _, ok := shape[next_pos]; ok {
							line = append(line, next_pos)
						} else {
							break
						}
					}
					if len(line) > longest {
						longest_lines = [][]Pos{line}
						longest = len(line)
					} else if len(line) > 1 && len(line) == longest {
						longest_lines = append(longest_lines, line)
					}

					line = []Pos{pos}
					for dy := y + 1; dy <= y_max; dy++ {
						next_pos := Pos{x: x, y: dy}
						if _, ok := shape[next_pos]; ok {
							line = append(line, next_pos)
						} else {
							break
						}
					}
					if len(line) > longest {
						longest_lines = [][]Pos{line}
						longest = len(line)
					} else if len(line) > 1 && len(line) == longest {
						longest_lines = append(longest_lines, line)
					}
				}
			}
		}

		if len(longest_lines) > 1 {
			min := 99999
			for _, line := range longest_lines {
				shape_cpy := copy_shape(shape)
				for _, pos := range line {
					delete(shape_cpy, pos)
				}
				count := dfs_remove_shape(shape_cpy, x_min, x_max, y_min, y_max, cache)
				if count < min {
					min = count
				}
			}
			spell_count += min + 1
			cache[hash] = spell_count
			return spell_count
		} else if len(longest_lines) == 1 {
			spell_count++
			for _, pos := range longest_lines[0] {
				delete(shape, pos)
			}
		}
	}

	cache[hash] = spell_count
	return spell_count
}

func remove_shape(shape map[Pos]struct{}) int {
	x_min, x_max, y_min, y_max := width, 0, height, 0
	for pos := range shape {
		x_min = utils.Min([]int{x_min, pos.x})
		x_max = utils.Max([]int{x_max, pos.x})
		y_min = utils.Min([]int{y_min, pos.y})
		y_max = utils.Max([]int{y_max, pos.y})
	}

	cache := map[uint64]int{}
	return dfs_remove_shape(shape, x_min, x_max, y_min, y_max, cache)
}

func spell_count(grid [][]byte) int {
	height = len(grid)
	width = len(grid[0])

	shapes := []map[Pos]struct{}{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '#' {
				shape := find_shape(grid, x, y)
				shapes = append(shapes, shape)
				for pos := range shape {
					grid[pos.y][pos.x] = '.'
				}
			}
		}
	}

	count := 0
	for _, shape := range shapes {
		count += remove_shape(shape)
	}

	return count
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 12, 1

	grid := loader.GetByteRows()
	part1 := spell_count(grid)

	loader.Part = 2
	grid = loader.GetByteRows()
	part2 := spell_count(grid)

	loader.Part = 3
	grid = loader.GetByteRows()
	part3 := spell_count(grid)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
