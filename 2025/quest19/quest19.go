package quest19

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Wall struct {
	x          int
	gap_bottom int
	gap_top    int
}

func parse_data(data []string) []Wall {
	walls := []Wall{}

	for _, line := range data {
		parts := strings.Split(line, ",")
		nums := [3]int{}
		for i := 0; i < 3; i++ {
			n, err := strconv.Atoi(parts[i])
			if err != nil {
				panic(err)
			}
			nums[i] = n
		}
		walls = append(walls, Wall{
			x:          nums[0],
			gap_bottom: nums[1],
			gap_top:    nums[1] + nums[2] - 1,
		})
	}

	return walls
}

func flap(walls []Wall) int {
	x, y := 0, 0

	flaps := 0
	for _, wall := range walls {
		for ; x <= wall.x; x++ {
			if y < wall.gap_bottom {
				flaps++
				y++
			} else {
				y--
			}
		}
	}

	return flaps
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 19, 1

	data := loader.GetStrings()
	walls := parse_data(data)
	part1 := flap(walls)

	fmt.Printf("%d %d %d\n", part1, 0, 0)
}
