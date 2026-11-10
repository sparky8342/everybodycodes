package quest11

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Portal struct {
	id          int
	connections []int
}

type QueueEntry struct {
	id    int
	steps int
}

func parse_data(data []string, neighbours int) []Portal {
	max_portal, err := strconv.Atoi(strings.Split(data[0], " ")[1])
	if err != nil {
		panic(err)
	}

	portals := make([]Portal, max_portal+1)

	for i := 1; i <= max_portal; i++ {
		portals[i] = Portal{id: i}
	}

	for i := 1; i <= max_portal; i++ {
		for j := i - neighbours; j <= i+neighbours; j++ {
			if j < 1 {
				portals[i].connections = append(portals[i].connections, j+max_portal)
			} else if j > max_portal {
				portals[i].connections = append(portals[i].connections, j-max_portal)
			} else {
				portals[i].connections = append(portals[i].connections, j)
			}
		}
	}

	return portals
}

func add_shortcuts(portals []Portal, data []string) {
	for _, line := range data[2:] {
		parts := strings.Split(line, " ")
		from, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		portals[from].connections = append(portals[from].connections, to)
		portals[to].connections = append(portals[to].connections, from)
	}
}

func resonance(data []string, neighbours int) int {
	portals := parse_data(data, neighbours)
	basic_strain := get_strain(portals)
	add_shortcuts(portals, data)
	strain := get_strain(portals)
	return basic_strain - strain
}

func resonance_extra_connection(data []string, neighbours int) int {
	portals := parse_data(data, neighbours)
	basic_strain := get_strain(portals)
	add_shortcuts(portals, data)

	min := 999999
	for i := 1; i < len(portals)-1; i++ {
		for j := i + 1; j < len(portals); j++ {
			portals[i].connections = append(portals[i].connections, j)
			portals[j].connections = append(portals[j].connections, i)
			strain := get_strain(portals)
			if strain < min {
				min = strain
			}
			portals[i].connections = portals[i].connections[:len(portals[i].connections)-1]
			portals[j].connections = portals[j].connections[:len(portals[j].connections)-1]
		}
	}

	return basic_strain - min
}

func get_strain(portals []Portal) int {
	strain := 0
	for i := 1; i < len(portals); i++ {
		queue := []QueueEntry{QueueEntry{id: i, steps: 0}}
		visited := map[int]struct{}{i: struct{}{}}
		for len(queue) > 0 {
			qe := queue[0]
			queue = queue[1:]

			strain += qe.steps

			for _, connection := range portals[qe.id].connections {
				if _, ok := visited[connection]; !ok {
					queue = append(queue, QueueEntry{id: connection, steps: qe.steps + 1})
					visited[connection] = struct{}{}
				}
			}
		}
	}
	return strain
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 11, 1

	data := loader.GetStrings()
	part1 := resonance(data, 1)

	loader.Part = 2
	data = loader.GetStrings()
	part2 := resonance_extra_connection(data, 3)

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
