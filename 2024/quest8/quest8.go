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

func build_part3(blocks int, multiplier int, mod int) int {
	width := 1
	height_to_add := 1
	heights := []int{1}

	for {
		width += 2
		height_to_add = ((height_to_add * multiplier) % mod) + mod
		for i := 0; i < len(heights); i++ {
			heights[i] += height_to_add
		}
		heights = append(heights, height_to_add)

		// central column
		blocks_used := heights[0]
		blocks_used -= (multiplier * width * heights[0]) % mod

		// end columns
		blocks_used += heights[len(heights)-1] * 2

		// all other columns
		for i := 1; i < len(heights)-1; i++ {
			blocks_used += heights[i] * 2
			blocks_used -= ((multiplier * width * heights[i]) % mod) * 2
		}

		if blocks_used > blocks {
			return blocks_used - blocks
		}
	}
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

	loader.Part = 3
	blocks = 202400000
	multiplier = loader.GetOneInt()
	mod = 10
	part3 := build_part3(blocks, multiplier, mod)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
