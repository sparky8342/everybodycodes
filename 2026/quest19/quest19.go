package quest19

import (
	"fmt"
	"loader"
	"sort"
)

func combine_colours(data []string) int {
	height := len(data) / 2
	width := len(data[0])

	colours := map[byte]int{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pair := [2]byte{data[y][x], data[y+height+1][x]}
			if pair[0] == 'K' || pair[1] == 'K' {
				colours['K']++
				continue
			}

			if pair[0] > pair[1] {
				pair[0], pair[1] = pair[1], pair[0]
			}
			switch pair {
			case [2]byte{'C', 'M'}:
				colours['B']++
			case [2]byte{'C', 'Y'}:
				colours['G']++
			case [2]byte{'M', 'Y'}:
				colours['R']++
			}
		}
	}

	values := []int{}
	for _, val := range colours {
		values = append(values, val)
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	return values[0] * values[1] * values[2]
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 19, 1

	data := loader.GetStrings()
	part1 := combine_colours(data)

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
