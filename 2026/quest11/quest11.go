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

func parse_data(data []string) []Portal {
	max_portal, err := strconv.Atoi(strings.Split(data[0], " ")[1])
	if err != nil {
		panic(err)
	}

	portals := make([]Portal, max_portal+1)

	for i := 1; i <= max_portal; i++ {
		portals[i] = Portal{id: i}
	}
	portals[1].connections = []int{max_portal, 2}
	for i := 2; i < max_portal; i++ {
		portals[i].connections = []int{i - 1, i + 1}
	}
	portals[max_portal].connections = []int{max_portal - 1, 1}

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

	return portals
}

func basic_strain(portals []Portal) int {
	l := len(portals)
	n := l / 2
	strain := n
	if l%2 == 0 {
		strain += n
	}
	for n > 0 {
		n--
		strain += n * 2
	}
	return strain * (l - 1)
}

func resonance(portals []Portal) int {
	basic_strain := basic_strain(portals)

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

	return basic_strain - strain
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 11, 1

	data := loader.GetStrings()
	portals := parse_data(data)
	part1 := resonance(portals)

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
