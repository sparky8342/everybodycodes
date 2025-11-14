package quest9

import (
	"fmt"
	"loader"
	"sort"
	"strings"
)

type intset map[int]struct{}

func (s intset) add(value int) {
	s[value] = struct{}{}
}

func parse_data(data []string) []string {
	for i := range data {
		parts := strings.Split(data[i], ":")
		data[i] = parts[1]
	}
	return data
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

func union(a intset, b intset) intset {
	combined := intset{}
	for id := range a {
		combined.add(id)
	}
	for id := range b {
		combined.add(id)
	}
	return combined
}

func intersect(a intset, b intset) bool {
	for id := range a {
		if _, ok := b[id]; ok {
			return true
		}
	}
	return false
}

func similarity_sum(sequences []string) (int, int) {
	sum := 0

	groups := []intset{}

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
					new_group := intset{}
					new_group.add(i)
					new_group.add(j)
					new_group.add(k)
					groups = append(groups, new_group)
				}
			}
		}
	}

outer:
	for {
		for i := 0; i < len(groups); i++ {
			for j := i + 1; j < len(groups); j++ {
				if intersect(groups[i], groups[j]) {
					new_group := union(groups[i], groups[j])
					groups = append(groups[0:j], groups[j+1:]...)
					groups = append(groups[0:i], groups[i+1:]...)
					groups = append(groups, new_group)
					continue outer
				}
			}
		}

		break
	}

	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})

	scale_sum := 0
	for dragonduck := range groups[0] {
		scale_sum += dragonduck + 1
	}

	return sum, scale_sum
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 9, 1

	data := loader.GetStrings()
	sequences := parse_data(data)
	part1 := similarity(sequences[2], sequences[0], sequences[1])

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
