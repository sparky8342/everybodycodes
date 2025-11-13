package quest9

import (
	"fmt"
	"loader"
	"strings"
)

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

func similarity_sum(sequences []string) int {
	sum := 0

	for i := 0; i < len(sequences); i++ {
		child := sequences[i]

		for j := 0; j < len(sequences); j++ {
			for k := j + 1; k < len(sequences); k++ {
				if j == i || k == i {
					continue
				}

				parent1 := sequences[j]
				parent2 := sequences[k]

				sum += similarity(child, parent1, parent2)
			}
		}
	}

	return sum
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 9, 1

	data := loader.GetStrings()
	sequences := parse_data(data)
	part1 := similarity_part1(sequences)

	loader.Part = 2
	data = loader.GetStrings()
	sequences = parse_data(data)
	part2 := similarity_sum(sequences)

	fmt.Printf("%d %d %d\n", part1, part2, 0)
}
