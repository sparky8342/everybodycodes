package quest8

import (
	"fmt"
	"loader"
)

func build(blocks int) int {
	triangle_size := 1
	blocks_used := 1

	for blocks_used < blocks {
		triangle_size += 2
		blocks_used += triangle_size
	}

	return (blocks_used - blocks) * triangle_size
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 8, 1

	blocks := loader.GetOneInt()
	part1 := build(blocks)

	part2 := -1
	part3 := -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
