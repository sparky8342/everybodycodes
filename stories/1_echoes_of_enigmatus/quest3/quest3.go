package quest3

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Snail struct {
	x         int
	y         int
	first_top int
	period    int
}

func parse_data(data []string) []Snail {
	snails := make([]Snail, len(data))
	for i, line := range data {
		parts := strings.Split(line, " ")
		x, err := strconv.Atoi(strings.Split(parts[0], "=")[1])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(strings.Split(parts[1], "=")[1])
		if err != nil {
			panic(err)
		}
		snails[i] = Snail{x: x, y: y, first_top: y - 1, period: y - 1 + x}
	}
	return snails
}

func move_snails(snails []Snail, days int) {
	for i := 0; i < days; i++ {
		for j := 0; j < len(snails); j++ {
			if snails[j].y == 1 {
				snails[j].x, snails[j].y = 1, snails[j].x
			} else {
				snails[j].x++
				snails[j].y--
			}
		}
	}
}

func position_sum(snails []Snail) int {
	total := 0
	for _, snail := range snails {
		total += snail.x + snail.y*100
	}
	return total
}

func all_top(snails []Snail) int {
	for len(snails) > 1 {
		a_day := snails[0].first_top
		b_day := snails[1].first_top

		period_start := -1

		for {
			if a_day < b_day {
				a_day += snails[0].period
			} else if b_day < a_day {
				b_day += snails[1].period
			} else {
				if period_start == -1 {
					period_start = a_day
					a_day += snails[0].period
					b_day += snails[1].period
				} else {
					period := a_day - period_start
					snails = snails[2:]
					snails = append(snails, Snail{first_top: period_start, period: period})
					break
				}
			}
		}
	}

	return snails[0].first_top
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "1", 3, 1

	data := loader.GetStrings()
	snails := parse_data(data)
	move_snails(snails, 100)
	part1 := position_sum(snails)

	loader.Part = 2
	data = loader.GetStrings()
	snails = parse_data(data)
	part2 := all_top(snails)

	loader.Part = 3
	data = loader.GetStrings()
	snails = parse_data(data)
	part3 := all_top(snails)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
