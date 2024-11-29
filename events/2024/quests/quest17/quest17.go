package quest17

import (
	"fmt"
	"loader"
	"math"
)

type Star struct {
	x int
	y int
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func parse_data(data []string) []Star {
	height := len(data)
	width := len(data[0])

	stars := []Star{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if data[y][x] == '*' {
				star := Star{x: x, y: y}
				stars = append(stars, star)
			}
		}
	}

	return stars
}

func distance(a Star, b Star) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func compute_distances(stars []Star) [][]int {
	distances := make([][]int, len(stars))
	for i := range stars {
		distances[i] = make([]int, len(stars))
	}

	for i := 0; i < len(stars); i++ {
		for j := i + 1; j < len(stars); j++ {
			distances[i][j] = distance(stars[i], stars[j])
			distances[j][i] = distances[i][j]
		}
	}

	return distances
}

func minimum_spanning_tree(stars []Star) int {
	distances := compute_distances(stars)

	// Prim's algorithm
	vertices := []int{0}

	available := map[int]struct{}{}
	for i := 1; i < len(stars); i++ {
		available[i] = struct{}{}
	}

	total := 0

	for len(available) > 0 {
		min := math.MaxInt32
		closest := 0
		for _, vertex := range vertices {
			for vertex2 := range available {
				dist := distances[vertex][vertex2]
				if dist < min {
					min = dist
					closest = vertex2
				}
			}
		}
		delete(available, closest)
		vertices = append(vertices, closest)
		total += min
	}

	return total + len(stars)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 17, 1

	data := loader.GetStrings()
	stars := parse_data(data)
	part1 := minimum_spanning_tree(stars)

	loader.Part = 2
	data = loader.GetStrings()
	stars = parse_data(data)
	part2 := minimum_spanning_tree(stars)

	part3 := -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
