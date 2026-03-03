package quest1

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

func colour_value(colour string) int {
	value := 0
	for i := 5; i >= 0; i-- {
		if colour[i] == 'R' || colour[i] == 'G' || colour[i] == 'B' {
			value += 1 << (5 - i)
		}
	}
	return value
}

func green_dominant(data []string) int {
	total := 0

	for _, line := range data {
		parts := strings.Split(line, ":")
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		parts2 := strings.Split(parts[1], " ")
		red := colour_value(parts2[0])
		green := colour_value(parts2[1])
		blue := colour_value(parts2[2])
		if green > red && green > blue {
			total += id
		}
	}

	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "3", 1, 1

	data := loader.GetStrings()
	part1 := green_dominant(data)

	fmt.Printf("%d %d %d\n", part1, -1, -1)
}
