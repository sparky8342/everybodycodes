package quest14

import (
	"fmt"
	"loader"
	"math"
	"strconv"
	"strings"
)

type Segment struct {
	x int
	y int
	z int
}

type Node struct {
	segment  Segment
	distance int
}

func max_height(data []byte) int {
	height := 0
	max := 0
	steps := strings.Split(string(data), ",")
	for _, step := range steps {
		if step[0] == 'U' || step[0] == 'D' {
			n, err := strconv.Atoi(step[1:])
			if err != nil {
				panic(err)
			}
			if step[0] == 'U' {
				height += n
				if height > max {
					max = height
				}
			} else if step[0] == 'D' {
				height -= n
			}
		}
	}
	return max
}

func unique_segments(data []string) int {
	segments := map[Segment]struct{}{}

	for _, line := range data {
		segment := Segment{}

		for _, step := range strings.Split(line, ",") {
			dir := step[0]
			n, err := strconv.Atoi(step[1:])
			if err != nil {
				panic(err)
			}

			for i := 0; i < n; i++ {
				switch dir {
				case 'U':
					segment.y++
				case 'D':
					segment.y--
				case 'R':
					segment.x++
				case 'L':
					segment.x--
				case 'F':
					segment.z++
				case 'B':
					segment.z--
				}
				segments[segment] = struct{}{}
			}
		}
	}

	return len(segments)
}

func path(segments map[Segment]struct{}, start Segment, leaves map[Segment]struct{}) int {
	start_node := Node{segment: start}
	queue := []Node{start_node}
	visited := map[Segment]struct{}{}
	visited[start] = struct{}{}

	distance := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		segment := node.segment
		if _, ok := leaves[segment]; ok {
			distance += node.distance
		}

		neighbours := []Segment{
			Segment{x: segment.x + 1, y: segment.y, z: segment.z},
			Segment{x: segment.x - 1, y: segment.y, z: segment.z},
			Segment{x: segment.x, y: segment.y + 1, z: segment.z},
			Segment{x: segment.x, y: segment.y - 1, z: segment.z},
			Segment{x: segment.x, y: segment.y, z: segment.z + 1},
			Segment{x: segment.x, y: segment.y, z: segment.z - 1},
		}

		for _, neighbour := range neighbours {
			if _, ok := segments[neighbour]; !ok {
				continue
			}
			if _, ok := visited[neighbour]; ok {
				continue
			}
			queue = append(queue, Node{segment: neighbour, distance: node.distance + 1})
			visited[neighbour] = struct{}{}
		}
	}

	return distance
}

func murkiness(data []string) int {
	leaves := map[Segment]struct{}{}
	segments := map[Segment]struct{}{}
	height := 0

	for _, line := range data {
		segment := Segment{}

		for _, step := range strings.Split(line, ",") {
			dir := step[0]
			n, err := strconv.Atoi(step[1:])
			if err != nil {
				panic(err)
			}

			for i := 0; i < n; i++ {
				switch dir {
				case 'U':
					segment.y++
				case 'D':
					segment.y--
				case 'R':
					segment.x++
				case 'L':
					segment.x--
				case 'F':
					segment.z++
				case 'B':
					segment.z--
				}

				segments[segment] = struct{}{}

				if segment.x == 0 && segment.z == 0 {
					if segment.y > height {
						height = segment.y
					}
				}
			}

		}
		leaves[segment] = struct{}{}
	}

	min_dist := math.MaxInt32
	trunk := Segment{}
	for trunk.y = 0; trunk.y <= height; trunk.y++ {
		dist := path(segments, trunk, leaves)
		if dist < min_dist {
			min_dist = dist
		}
	}

	return min_dist
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 14, 1

	data := loader.GetOneLine()
	part1 := max_height(data)

	loader.Part = 2
	data2 := loader.GetStrings()
	part2 := unique_segments(data2)

	loader.Part = 3
	data3 := loader.GetStrings()
	part3 := murkiness(data3)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
