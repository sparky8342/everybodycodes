package quest18

import (
	"fmt"
	"loader"
	"slices"
	"strconv"
	"strings"
)

type Stack struct {
	x int
	y int
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func print_stacks(stacks []Stack, max_y int) {
	grid := make([][]byte, max_y+1)
	for i := range grid {
		grid[i] = make([]byte, 100)
		for j := 0; j < 100; j++ {
			grid[i][j] = ' '
		}
	}
	for i := 0; i < len(stacks); i++ {
		for y := max_y; y > max_y-stacks[i].y; y-- {
			grid[y][stacks[i].x] = '#'
		}
	}
	for _, r := range grid {
		fmt.Println(string(r))
	}
	fmt.Println()
}

func material_needed(data []string) int {
	stacks := []Stack{}
	max_y := 0
	for _, line := range data {
		parts := strings.Split(line, " ")
		x, err := strconv.Atoi(strings.Split(parts[0], "=")[1])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(strings.Split(parts[1], "=")[1])
		if err != nil {
			panic(err)
		}
		if y > max_y {
			max_y = y
		}

		width := 1
		if len(parts) == 3 {
			width, err = strconv.Atoi(strings.Split(parts[2], "=")[1])
			if err != nil {
				panic(err)
			}
		}

		stacks = append(stacks, Stack{x: x, y: y})
		if width > 1 {
			stacks = append(stacks, Stack{x: x + width - 1, y: y})
		}
	}

	material := 0

	// remove in betweens
	for i := 0; i < len(stacks)-2; i++ {
		if stacks[i].y >= stacks[i+1].y && stacks[i+1].y <= stacks[i+2].y {
			stacks = slices.Delete(stacks, i+1, i+2)
			i = -1
		}
	}

	for i := 0; i < len(stacks)-1; i++ {
		material += stacks[i].y
		material += min(stacks[i].y, stacks[i+1].y) * (stacks[i+1].x - stacks[i].x - 1)
	}
	material += stacks[len(stacks)-1].y

	return material
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 18, 1

	data := loader.GetStrings()
	part1 := material_needed(data)

	loader.Part = 2
	data = loader.GetStrings()
	part2 := material_needed(data)

	loader.Part = 3
	data = loader.GetStrings()
	part3 := material_needed(data)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
