package quest1

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

func parse_data(data []string) ([]string, []int) {
	names := strings.Split(data[0], ",")
	moves := []int{}

	for _, ins := range strings.Split(data[2], ",") {
		n, err := strconv.Atoi(ins[1:])
		if err != nil {
			panic(err)
		}

		if ins[0] == 'L' {
			n *= -1
		}

		moves = append(moves, n)
	}

	return names, moves
}

func get_name(names []string, moves []int) string {
	pos := 0
	for _, move := range moves {
		pos += move
		if pos < 0 {
			pos = 0
		} else if pos >= len(names) {
			pos = len(names) - 1
		}
	}
	return names[pos]
}

func get_name_circular(names []string, moves []int) string {
	pos := 0
	for _, move := range moves {
		pos = (pos + move) % len(names)
		for pos < 0 {
			pos += len(names)
		}
	}
	return names[pos]
}

func get_name_with_swaps(names []string, moves []int) string {
	for _, move := range moves {
		move = move % len(names)
		for move < 0 {
			move += len(names)
		}
		names[0], names[move] = names[move], names[0]
	}

	return names[0]
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 1, 1

	data := loader.GetStrings()
	names, moves := parse_data(data)
	part1 := get_name(names, moves)

	loader.Part = 2
	data = loader.GetStrings()
	names, moves = parse_data(data)
	part2 := get_name_circular(names, moves)

	loader.Part = 3
	data = loader.GetStrings()
	names, moves = parse_data(data)
	part3 := get_name_with_swaps(names, moves)

	fmt.Printf("%s %s %s\n", part1, part2, part3)
}
