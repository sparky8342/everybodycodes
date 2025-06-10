package quest3

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Snail struct {
	x int
	y int
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
		snails[i] = Snail{x: x, y: y}
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

func Run() {
	loader.Event, loader.Quest, loader.Part = "1", 3, 1

	data := loader.GetStrings()
	snails := parse_data(data)
	move_snails(snails, 100)
	part1 := position_sum(snails)

	part2, part3 := "", ""

	fmt.Printf("%d %s %s\n", part1, part2, part3)
}
