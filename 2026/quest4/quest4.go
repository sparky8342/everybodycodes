package quest4

import (
	"fmt"
	"loader"
	"sort"
	"strconv"
	"strings"
)

type Warrior struct {
	id   int
	time int
}

func parse_data(data []string) []*Warrior {
	warriors := map[int]*Warrior{}

	for i := 0; i < len(data); i++ {
		parts := strings.Split(data[i], " ")
		id, err := strconv.Atoi(strings.Split(parts[0], "=")[1])
		if err != nil {
			panic(err)
		}
		time, err := strconv.Atoi(strings.Split(parts[3], "=")[1])
		if err != nil {
			panic(err)
		}

		if _, ok := warriors[id]; !ok {
			warriors[id] = &Warrior{id: id}
		}

		warriors[id].time += time
	}

	sl := []*Warrior{}
	for _, warrior := range warriors {
		sl = append(sl, warrior)
	}

	return sl
}

func order_warriors(warriors []*Warrior) int {
	sort.Slice(warriors, func(i, j int) bool {
		return warriors[i].time < warriors[j].time
	})
	checksum := 0
	for i := 0; i < len(warriors); i++ {
		checksum += (i + 1) * warriors[i].time
	}
	return checksum
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 4, 1

	data := loader.GetStrings()
	warriors := parse_data(data)
	part1 := order_warriors(warriors)

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
