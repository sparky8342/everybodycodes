package quest4

import (
	"fmt"
	"loader"
)

func turns(gears []int) int {
	return gears[0] * 2025 / gears[len(gears)-1]
}

func turns_needed(gears []int) int {
	return int((10000000000000.0 * float64(gears[len(gears)-1]) / float64(gears[0])) + 0.9)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 4, 1

	gears := loader.GetInts()
	part1 := turns(gears)

	loader.Part = 2
	gears = loader.GetInts()
	part2 := turns_needed(gears)

	fmt.Printf("%d %d %d\n", part1, part2, 0)
}
