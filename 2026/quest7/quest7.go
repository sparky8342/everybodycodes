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

type Pulse struct {
	energy  int
	targets map[int]int
}

func (c *Cloud) clone() *Cloud {
	return &Cloud{
		id:          c.id,
		energy:      c.energy,
		connections: c.connections, // connections don't change so no need to deepcopy them
	}
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

func charge(clouds map[int]*Cloud, pulse int) int {
	start := 1
	for id, cloud := range clouds {
		if len(cloud.connections) > len(clouds[start].connections) {
			start = id
		}
	}

	targets := map[int]int{start: 1}

	for pulse > 0 {
		new_targets := map[int]int{}
		for target, amount := range targets {
			clouds[target].energy += amount * pulse
			for _, connection := range clouds[target].connections {
				new_targets[connection] += amount
			}
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

func charge_with_explosions(clouds map[int]*Cloud, initial_energy int, start int, explosion_limit int) int {
	pulses := []Pulse{
		Pulse{
			energy:  initial_energy,
			targets: map[int]int{start: 1},
		},
	}

	for len(pulses) > 0 {
		new_pulses := []Pulse{}
		for _, pulse := range pulses {
			if pulse.energy == 0 {
				continue
			}
			new_targets := map[int]int{}
			for target, amount := range pulse.targets {
				if _, ok := clouds[target]; !ok {
					continue
				}
				clouds[target].energy += amount * pulse.energy
				for _, connection := range clouds[target].connections {
					new_targets[connection] += amount
				}
			}
			new_pulses = append(new_pulses, Pulse{
				energy:  pulse.energy - 1,
				targets: new_targets,
			})
		}
		pulses = new_pulses

		to_delete := []int{}
		for id, cloud := range clouds {
			if cloud.energy >= explosion_limit {
				pulse := Pulse{
					energy:  initial_energy,
					targets: map[int]int{},
				}
				for _, connection := range cloud.connections {
					pulse.targets[connection] = 1
				}
				pulses = append(pulses, pulse)
				to_delete = append(to_delete, id)
			}
		}
		for _, id := range to_delete {
			delete(clouds, id)
		}
	}

	total := 0
	for _, cloud := range clouds {
		total += cloud.energy
	}
	return total
}

func find_best_start(clouds map[int]*Cloud, initial_energy int, explosion_limit int) int {
	max := 0
	for start := range clouds {
		clouds_clone := map[int]*Cloud{}
		for id, cloud := range clouds {
			clouds_clone[id] = cloud.clone()
		}
		energy := charge_with_explosions(clouds_clone, initial_energy, start, explosion_limit)
		if energy > max {
			max = energy
		}
	}
	return max
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 7, 1

	data := loader.GetStrings()
	clouds := parse_data(data)
	part1 := charge(clouds, 5)

	loader.Part = 2
	data = loader.GetStrings()
	clouds = parse_data(data)
	part2 := charge(clouds, 10)

	loader.Part = 3
	data = loader.GetStrings()
	clouds = parse_data(data)
	part3 := find_best_start(clouds, 10, 1000000)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
