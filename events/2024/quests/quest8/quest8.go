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

func build_part2(blocks int, multiplier int, mod int) int {
	width := 1
	height := 1
	blocks_used := 1

	for blocks_used < blocks {
		width += 2
		height = (height * multiplier) % mod
		blocks_used += width * height
	}

	return (blocks_used - blocks) * width
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 8, 1

	blocks := loader.GetOneInt()
	part1 := build(blocks)

	loader.Part = 2
	blocks = 20240000
	multiplier := loader.GetOneInt()
	mod := 1111
	part2 := build_part2(blocks, multiplier, mod)

	part3 := -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
