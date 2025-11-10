package quest6

import (
	"fmt"
	"loader"
)

func total_A_pairs(data []byte) int {
	pairs := 0
	A := 0
	for _, b := range data {
		if b == 'A' {
			A++
		} else if b == 'a' {
			pairs += A
		}
	}
	return pairs
}

func total_pairs(data []byte) int {
	pairs := 0
	A, B, C := 0, 0, 0
	for _, b := range data {
		switch b {
		case 'A':
			A++
		case 'B':
			B++
		case 'C':
			C++
		case 'a':
			pairs += A
		case 'b':
			pairs += B
		case 'c':
			pairs += C
		}
	}
	return pairs
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 6, 1

	data := loader.GetOneLine()
	part1 := total_A_pairs(data)

	loader.Part = 2
	data = loader.GetOneLine()
	part2 := total_pairs(data)

	fmt.Printf("%d %d %d\n", part1, part2, 0)
}
