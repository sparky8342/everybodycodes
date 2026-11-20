package quest19

import (
	"fmt"
	"loader"
	"sort"
)

func parse_data(data []string) ([][]byte, [][]byte) {
	h := len(data) / 2
	grid1 := make([][]byte, h)
	for y := range grid1 {
		grid1[y] = []byte(data[y])
	}
	grid2 := make([][]byte, h)
	for y := range grid2 {
		grid2[y] = []byte(data[y+h+1])
	}
	return grid1, grid2
}

func combine_colours(grid1 [][]byte, grid2 [][]byte) int {
	height := len(grid1)
	width := len(grid1[0])

	colours := map[byte]int{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pair := [2]byte{grid1[y][x], grid2[y][x]}
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

func combinations(grid1 [][]byte, grid2 [][]byte) int {
	combos := [][]byte{
		[]byte("CMYK"),
		[]byte("CMKY"),
		[]byte("CYMK"),
		[]byte("CYKM"),
		[]byte("CKYM"),
		[]byte("CKMY"),
		[]byte("MCYK"),
		[]byte("MCKY"),
		[]byte("MYCK"),
		[]byte("MYKC"),
		[]byte("MKCY"),
		[]byte("MKYC"),
		[]byte("YCMK"),
		[]byte("YCKM"),
		[]byte("YMCK"),
		[]byte("YMKC"),
		[]byte("YKCM"),
		[]byte("YKMC"),
		[]byte("KCMY"),
		[]byte("KCYM"),
		[]byte("KMCY"),
		[]byte("KMYC"),
		[]byte("KYCM"),
		[]byte("KYMC"),
	}

	height := len(grid1)
	width := len(grid1[0])

	max := 0
	for _, combo := range combos {
		colours := make([][]byte, height)
		for y := 0; y < len(grid2); y++ {
			colours[y] = make([]byte, width)
			copy(colours[y], grid2[y])
		}

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				colours[i][j] = combo[int(colours[i][j]-'0')-1]
			}
		}

		crystals := combine_colours(grid1, colours)
		if crystals > max {
			max = crystals
		}
	}

	return max
}

func double_combinations(data []string) int {
	height := len(data)
	width := len(data[0])

	K := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if data[y][x] == 'K' {
				K++
			}
		}
	}

	no_colours := height*width - K

	// 2 colours
	a := no_colours / 2
	b := no_colours - a
	max := a * b * K

	// 3 colours
	a = no_colours / 3
	no_colours -= a
	b = no_colours / 2
	c := no_colours - b
	ans := a * b * c
	if ans > max {
		max = ans
	}

	return max
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 19, 1

	data := loader.GetStrings()
	grid1, grid2 := parse_data(data)
	part1 := combine_colours(grid1, grid2)

	loader.Part = 2
	data = loader.GetStrings()
	grid1, grid2 = parse_data(data)
	part2 := combinations(grid1, grid2)

	loader.Part = 3
	data = loader.GetStrings()
	part3 := double_combinations(data)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
