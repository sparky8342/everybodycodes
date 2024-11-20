package quest1

import (
	"fmt"
	"loader"
)

func calculate_potions(data []byte) int {
	potions := 0
	for _, b := range data {
		switch b {
		case 'B':
			potions += 1
		case 'C':
			potions += 3
		case 'D':
			potions += 5
		}
	}
	return potions
}

func calculate_pairs(data []byte) int {
	potions := calculate_potions(data)
	for i := 0; i < len(data); i += 2 {
		if data[i] != 'x' && data[i+1] != 'x' {
			potions += 2
		}
	}
	return potions
}

func calculate_triples(data []byte) int {
	potions := calculate_potions(data)
	for i := 0; i < len(data); i += 3 {
		x := 0
		for j := 0; j < 3; j++ {
			if data[i+j] == 'x' {
				x++
			}
		}
		if x == 0 {
			potions += 6
		} else if x == 1 {
			potions += 2
		}
	}
	return potions
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 1, 1

	data := loader.GetOneLine()
	part1 := calculate_potions(data)

	loader.Part = 2
	data = loader.GetOneLine()
	part2 := calculate_pairs(data)

	loader.Part = 3
	data = loader.GetOneLine()
	part3 := calculate_triples(data)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
