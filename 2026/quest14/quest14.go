package quest14

import (
	"fmt"
	"loader"
	"math"
)

type Pos struct {
	x      int
	y      int
	border int
}

type Connection struct {
	a Pos
	b Pos
}

func bindings_needed(window []string) int {
	height := len(window)
	width := len(window[0])

	bindings := height*2 + width*2

	for y := 0; y < height; y++ {
		for x := 1; x < width; x++ {
			if window[y][x] != window[y][x-1] {
				bindings++
			}
			if y > 0 && window[y][x] != window[y-1][x] {
				bindings++
			}
		}
	}

	return bindings + 1
}

func minimum_bindings_needed(window []string) int {
	height := len(window)
	width := len(window[0])

	border_shapes := map[byte][]Pos{}
	connections := map[[2]byte]int{}
	shape_connections := map[byte]map[byte]struct{}{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			shape_id := window[y][x]
			if _, ok := shape_connections[shape_id]; !ok {
				shape_connections[shape_id] = map[byte]struct{}{}
			}

			pos := Pos{x: x, y: y}
			if x == 0 || x == width-1 {
				pos.border++
			}
			if y == 0 || y == height-1 {
				pos.border++
			}
			if pos.border > 0 {
				border_shapes[shape_id] = append(border_shapes[shape_id], pos)
			}

			if x > 0 && window[y][x] != window[y][x-1] {
				name := [2]byte{window[y][x], window[y][x-1]}
				if name[0] > name[1] {
					name[0], name[1] = name[1], name[0]
				}
				shape_connections[name[0]][name[1]] = struct{}{}
				shape_connections[name[1]][name[0]] = struct{}{}
				connections[name]++
				connections[[2]byte{name[1], name[0]}]++
			}
			if y > 0 && window[y][x] != window[y-1][x] {
				name := [2]byte{window[y][x], window[y-1][x]}
				if name[0] > name[1] {
					name[0], name[1] = name[1], name[0]
				}
				shape_connections[name[0]][name[1]] = struct{}{}
				shape_connections[name[1]][name[0]] = struct{}{}
				connections[name]++
				connections[[2]byte{name[1], name[0]}]++

			}
		}
	}

	smallest := minimum_spanning_tree(border_shapes, shape_connections, connections)

	return smallest
}

func minimum_spanning_tree(border_shapes map[byte][]Pos, shape_connections map[byte]map[byte]struct{}, connections map[[2]byte]int) int {
	shape_nums := map[byte]int{}
	num_shapes := map[int]byte{}
	// 0 is outside
	num := 1
	for shape_id := range shape_connections {
		shape_nums[shape_id] = num
		num_shapes[num] = shape_id
		num++
	}

	// Prim's algorithm
	vertices := []int{0}

	available := map[int]struct{}{}
	for shape_id := range shape_connections {
		available[shape_nums[shape_id]] = struct{}{}
	}

	total := 0

	for len(available) > 0 {
		min := math.MaxInt32
		closest := 0
		for _, vertex := range vertices {
			for vertex2 := range available {
				if vertex == 0 || vertex2 == 0 {
					var id int
					if vertex == 0 {
						id = vertex2
					} else if vertex2 == 0 {
						id = vertex
					}
					if blocks, ok := border_shapes[num_shapes[id]]; ok {
						dist := 0
						for _, block := range blocks {
							dist += block.border
						}
						if dist < min {
							min = dist
							closest = vertex2
						}
					}
					continue
				}

				name := [2]byte{num_shapes[vertex], num_shapes[vertex2]}
				if name[0] > name[1] {
					name[0], name[1] = name[1], name[0]
				}
				if dist, ok := connections[name]; ok {
					if dist < min {
						min = dist
						closest = vertex2
					}
				}
			}
		}
		delete(available, closest)
		vertices = append(vertices, closest)
		total += min
	}

	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 14, 1

	window := loader.GetStrings()
	part1 := bindings_needed(window)

	loader.Part = 2
	window = loader.GetStrings()
	part2 := minimum_bindings_needed(window)

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
