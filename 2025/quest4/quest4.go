package quest4

import (
	"fmt"
	"loader"
)

func turns(gears []int) int {
	return gears[0] * 2025 / gears[len(gears)-1]
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 4, 1

	gears := loader.GetInts()
	part1 := turns(gears)

	fmt.Printf("%d %d %d\n", part1, 0, 0)
}
