package quest1

import (
	"fmt"
	"loader"
	"sort"
	"strconv"
	"strings"
)

type Scale struct {
	id    int
	red   int
	green int
	blue  int
	sum   int
	shine int
}

func colour_value(colour string) int {
	value := 0
	for i := 5; i >= 0; i-- {
		if colour[i] == 'R' || colour[i] == 'G' || colour[i] == 'B' || colour[i] == 'S' {
			value += 1 << (5 - i)
		}
	}
	return value
}

func parse_data(data []string) []Scale {
	scales := make([]Scale, len(data))

	for i, line := range data {
		parts := strings.Split(line, ":")

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		scales[i].id = id

		parts2 := strings.Split(parts[1], " ")
		scales[i].red = colour_value(parts2[0])
		scales[i].green = colour_value(parts2[1])
		scales[i].blue = colour_value(parts2[2])
		scales[i].sum = scales[i].red + scales[i].green + scales[i].blue

		if len(parts2) == 4 {
			scales[i].shine = colour_value(parts2[3])
		}
	}

	return scales
}

func green_dominant(scales []Scale) int {
	total := 0
	for _, scale := range scales {
		if scale.green > scale.red && scale.green > scale.blue {
			total += scale.id
		}
	}
	return total
}

func darkest_shiny(scales []Scale) int {
	sort.Slice(scales, func(i, j int) bool {
		if scales[i].shine == scales[j].shine {
			return scales[i].sum < scales[j].sum
		} else {
			return scales[i].shine > scales[j].shine
		}
	})
	return scales[0].id
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "3", 1, 1

	data := loader.GetStrings()
	scales := parse_data(data)
	part1 := green_dominant(scales)

	loader.Part = 2
	data = loader.GetStrings()
	scales = parse_data(data)
	part2 := darkest_shiny(scales)

	fmt.Printf("%d %d %d\n", part1, part2, -1)
}
