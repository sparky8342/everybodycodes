package quest1

import (
	"fmt"
	"loader"
	"math"
)

func next_square(n int) int {
	sqrt := math.Sqrt(float64(n))
	floor := float64(int(sqrt))
	if sqrt == floor {
		return n
	}
	floor += 1
	return int(floor * floor)
}

func pines_needed(pines []int) int {
	needed := 0
	for _, n := range pines {
		needed += next_square(n) - n
	}
	return needed
}

func next_frame(n int) int {
	if n <= 8 {
		return 8
	}
	frame_no := (float64(n) - 4) / 4
	floor := float64(int(frame_no))
	if frame_no == floor {
		return n
	}
	floor += 1
	return int(floor*4 + 4)
}

func pines_needed_frame(pines []int) int {
	needed := 0
	for _, n := range pines {
		needed += next_frame(n) - n
	}
	return needed
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 1, 1

	pines := loader.GetCSVInts()[0]
	part1 := pines_needed(pines)

	loader.Part = 2
	pines = loader.GetCSVInts()[0]
	part2 := pines_needed_frame(pines)

	part3 := -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
