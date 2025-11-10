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

func total_pairs_in_range(data []byte, repeat int, dist int) int {
	if repeat > 1 {
		r := []byte{}
		for i := 0; i < repeat; i++ {
			r = append(r, data...)
		}
		data = r
	}

	pairs := 0

	A, B, C := 0, 0, 0
	for i := 0; i <= dist; i++ {
		switch data[i] {
		case 'A':
			A++
		case 'B':
			B++
		case 'C':
			C++
		}
	}

	p1, p2 := 0, dist

	for _, b := range data {
		switch b {
		case 'a':
			pairs += A
		case 'b':
			pairs += B
		case 'c':
			pairs += C
		}

		if p2 < len(data)-1 {
			p2++
			switch data[p2] {
			case 'A':
				A++
			case 'B':
				B++
			case 'C':
				C++
			}
		}

		if p2 > dist*2 {
			switch data[p1] {
			case 'A':
				A--
			case 'B':
				B--
			case 'C':
				C--
			}
			p1++
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

	loader.Part = 3
	data = loader.GetOneLine()
	part3 := total_pairs_in_range(data, 1000, 1000)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
