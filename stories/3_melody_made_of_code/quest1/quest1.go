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
		if colour[i] < 91 { // is upper case
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

		parts = strings.Split(parts[1], " ")
		scales[i].red = colour_value(parts[0])
		scales[i].green = colour_value(parts[1])
		scales[i].blue = colour_value(parts[2])
		scales[i].sum = scales[i].red + scales[i].green + scales[i].blue

		if len(parts) == 4 {
			scales[i].shine = colour_value(parts[3])
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

func largest_group(scales []Scale) int {
	scale_types := map[string][]int{}

	for _, scale := range scales {
		var colour byte
		if scale.red > scale.green && scale.red > scale.blue {
			colour = 'R'
		} else if scale.green > scale.red && scale.green > scale.blue {
			colour = 'G'
		} else if scale.blue > scale.red && scale.blue > scale.green {
			colour = 'B'
		} else {
			continue
		}

		var shine byte
		if scale.shine <= 30 {
			shine = 'M'
		} else if scale.shine >= 33 {
			shine = 'G'
		} else {
			continue
		}

		typ := string([]byte{colour, shine})
		scale_types[typ] = append(scale_types[typ], scale.id)
	}

	best := 0
	best_type := ""
	for key, value := range scale_types {
		if len(value) > best {
			best = len(value)
			best_type = key
		}
	}

	total := 0
	for _, id := range scale_types[best_type] {
		total += id
	}

	return total
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

	loader.Part = 3
	data = loader.GetStrings()
	scales = parse_data(data)
	part3 := largest_group(scales)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
