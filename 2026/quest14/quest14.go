package quest14

import (
	"fmt"
	"loader"
	"math"
)

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

	num := 1
	shape_ids := map[byte]int{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			id := window[y][x]
			if _, ok := shape_ids[id]; !ok {
				shape_ids[id] = num
				num++
			}
		}
	}
	l := len(shape_ids) + 1

	connections := make([][]int, l)
	for i := 0; i < l; i++ {
		connections[i] = make([]int, l)
		for j := 0; j < l; j++ {
			connections[i][j] = 0
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			shape_id := window[y][x]

			border := 0
			if x == 0 || x == width-1 {
				border++
			}
			if y == 0 || y == height-1 {
				border++
			}
			if border > 0 {
				connections[0][shape_ids[shape_id]] += border
				connections[shape_ids[shape_id]][0] += border
			}

			if x > 0 && window[y][x] != window[y][x-1] {
				a := shape_ids[window[y][x]]
				b := shape_ids[window[y][x-1]]
				connections[a][b]++
				connections[b][a]++
			}
			if y > 0 && window[y][x] != window[y-1][x] {
				a := shape_ids[window[y][x]]
				b := shape_ids[window[y-1][x]]
				connections[a][b]++
				connections[b][a]++
			}
		}
	}

	return minimum_spanning_tree(connections)
}

func minimum_spanning_tree(connections [][]int) int {
	vertices := []int{0}

	available := map[int]struct{}{}
	for i := 1; i < len(connections); i++ {
		available[i] = struct{}{}
	}

	total := 0

	for len(available) > 0 {
		min := math.MaxInt32
		closest := 0
		for _, vertex := range vertices {
			for vertex2 := range available {
				if connections[vertex][vertex2] == 0 {
					continue
				}
				if connections[vertex][vertex2] < min {
					min = connections[vertex][vertex2]
					closest = vertex2
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
