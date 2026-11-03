package quest7

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Cloud struct {
	id          int
	energy      int
	connections []int
}

func parse_data(data []string) map[int]*Cloud {
	clouds := map[int]*Cloud{}
	for _, line := range data {
		parts := strings.Split(line, ", ")
		id, err := strconv.Atoi(strings.Split(parts[0], "=")[1])
		if err != nil {
			panic(err)
		}
		energy, err := strconv.Atoi(strings.Split(parts[1], "=")[1])
		if err != nil {
			panic(err)
		}
		connections := []int{}
		for _, str := range strings.Split(strings.Split(parts[2], "=")[1], ",") {
			connection, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			connections = append(connections, connection)
		}
		clouds[id] = &Cloud{
			id:          id,
			energy:      energy,
			connections: connections,
		}
	}
	return clouds
}

func charge(clouds map[int]*Cloud) int {
	start := 1
	for id, cloud := range clouds {
		if len(cloud.connections) > len(clouds[start].connections) {
			start = id
		}
	}

	pulse := 5
	targets := []int{start}

	for pulse > 0 {
		new_targets := []int{}
		for _, target := range targets {
			clouds[target].energy += pulse
			new_targets = append(new_targets, clouds[target].connections...)
		}
		targets = new_targets
		pulse--
	}

	total := 0
	for _, cloud := range clouds {
		total += cloud.energy
	}
	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 7, 1

	data := loader.GetStrings()
	clouds := parse_data(data)
	part1 := charge(clouds)

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
