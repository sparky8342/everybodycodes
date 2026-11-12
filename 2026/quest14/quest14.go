package quest14

import (
	"fmt"
	"loader"
)

func bindings_needed(window []string) int {
	height := len(window)
	width := len(window[0])

	bindings := height*2 + width*2

	for y := 0; y < height; y++ {
		for x := 1; x < width; x++ {
			if window[y][x] != window[y][x-1] {
				bindings++
			}
			if y > 0 && window[y][x] != window[y-1][x] {
				bindings++
			}
		}
	}

	return bindings + 1
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 14, 1

	window := loader.GetStrings()
	part1 := bindings_needed(window)

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
