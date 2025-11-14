package quest9

import (
	"fmt"
	"loader"
	"strings"
)

type DragonDuck struct {
	id       int
	parent1  *DragonDuck
	parent2  *DragonDuck
	children []*DragonDuck
}

func parse_data(data []string) []string {
	for i := range data {
		parts := strings.Split(data[i], ":")
		data[i] = parts[1]
	}
	return data
}

func similarity_part1(sequences []string) int {
	parent1 := sequences[0]
	parent2 := sequences[1]
	child := sequences[2]

	p1_matches := 0
	p2_matches := 0
	for i := 0; i < len(child); i++ {
		if parent1[i] == child[i] {
			p1_matches++
		}
		if parent2[i] == child[i] {
			p2_matches++
		}
	}

	return p1_matches * p2_matches
}

func similarity(child string, parent1 string, parent2 string) int {
	p1_matches := 0
	p2_matches := 0

	for i := 0; i < len(child); i++ {
		if child[i] != parent1[i] && child[i] != parent2[i] {
			return 0
		}
		if child[i] == parent1[i] {
			p1_matches++
		}
		if child[i] == parent2[i] {
			p2_matches++
		}
	}

	return p1_matches * p2_matches
}

func similarity_sum(sequences []string) (int, int) {
	sum := 0

	dragonducks := make([]*DragonDuck, len(sequences))
	for i := range dragonducks {
		dragonducks[i] = &DragonDuck{id: i}
	}

	for i := 0; i < len(sequences); i++ {
		child := sequences[i]

		for j := 0; j < len(sequences); j++ {
			for k := j + 1; k < len(sequences); k++ {
				if j == i || k == i {
					continue
				}

				parent1 := sequences[j]
				parent2 := sequences[k]

				score := similarity(child, parent1, parent2)
				if score > 0 {
					sum += score
					dragonducks[i].parent1 = dragonducks[j]
					dragonducks[i].parent2 = dragonducks[k]
					dragonducks[j].children = append(dragonducks[j].children, dragonducks[i])
					dragonducks[k].children = append(dragonducks[k].children, dragonducks[i])
				}
			}
		}
	}

	max_size := 0
	scale_sum := 0

	to_check := map[int]struct{}{}
	for i := 0; i < len(dragonducks); i++ {
		to_check[i] = struct{}{}
	}

	for len(to_check) > 0 {

		var check int
		for i := range to_check {
			check = i
			delete(to_check, i)
			break
		}

		queue := []*DragonDuck{dragonducks[check]}
		visited := map[int]struct{}{}
		visited[dragonducks[check].id] = struct{}{}

		for len(queue) > 0 {
			dd := queue[0]
			queue = queue[1:]

			if dd.parent1 != nil {
				if _, ok := visited[dd.parent1.id]; !ok {
					queue = append(queue, dd.parent1)
					visited[dd.parent1.id] = struct{}{}
					delete(to_check, dd.parent1.id)
				}
			}

			if dd.parent2 != nil {
				if _, ok := visited[dd.parent2.id]; !ok {
					queue = append(queue, dd.parent2)
					visited[dd.parent2.id] = struct{}{}
					delete(to_check, dd.parent2.id)
				}
			}

			for _, child := range dd.children {
				if _, ok := visited[child.id]; !ok {
					queue = append(queue, child)
					visited[child.id] = struct{}{}
					delete(to_check, child.id)
				}
			}
		}

		if len(visited) > max_size {
			max_size = len(visited)
			scale_sum = 0
			for id := range visited {
				scale_sum += id + 1
			}
		}
	}

	return sum, scale_sum
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 9, 1

	data := loader.GetStrings()
	sequences := parse_data(data)
	part1 := similarity_part1(sequences)

	loader.Part = 2
	data = loader.GetStrings()
	sequences = parse_data(data)
	part2, _ := similarity_sum(sequences)

	loader.Part = 3
	data = loader.GetStrings()
	sequences = parse_data(data)
	_, part3 := similarity_sum(sequences)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
