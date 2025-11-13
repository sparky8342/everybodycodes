package quest9

import (
	"fmt"
	"loader"
)

func parse_data(data []string) []string {
	for i := range data {
		data[i] = data[i][2:]
	}
	return data
}

func similarity(sequences []string) int {
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

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 9, 1

	data := loader.GetStrings()

	sequences := parse_data(data)
	part1 := similarity(sequences)

	fmt.Printf("%d %d %d\n", part1, 0, 0)
}
