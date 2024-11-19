package day1

import (
	"fmt"
	"loader"
)

func calculate_potions(data []byte) int {
	potions := 0
	for _, b := range data {
		if b == 'B' {
			potions++
		} else if b == 'C' {
			potions += 3
		}
	}
	return potions
}

func calculate_group(data []byte) int {
	potions := 0
	x := 0
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case 'B':
			potions += 1
		case 'C':
			potions += 3
		case 'D':
			potions += 5
		case 'x':
			x++
		}
	}
	if len(data) == 2 && x == 0 {
		potions += 2
	}
	if len(data) == 3 {
		if x == 0 {
			potions += 6
		} else if x == 1 {
			potions += 2
		}
	}
	return potions
}

func calculate_potions_pairs(data []byte) int {
	potions := 0
	for i := 0; i < len(data); i += 2 {
		potions += calculate_group(data[i : i+2])
	}
	return potions
}

func calculate_potions_triples(data []byte) int {
	potions := 0
	for i := 0; i < len(data); i += 3 {
		potions += calculate_group(data[i : i+3])
	}

	return potions
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 1, 1

	data := loader.GetOneLine()
	fmt.Println(calculate_potions(data))

	loader.Part = 2
	data = loader.GetOneLine()
	fmt.Println(calculate_potions_pairs(data))

	loader.Part = 3
	data = loader.GetOneLine()
	fmt.Println(calculate_potions_triples(data))
}
