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

var combos [][]int

func get_combos(taken map[int]struct{}, nums []int, combos *[][]int) {
	if len(nums) == 8 {
		*combos = append(*combos, nums)
		return
	}
	if len(nums)%2 == 0 {
		var smallest int
		for i := 0; i < 8; i++ {
			if _, ok := taken[i]; !ok {
				smallest = i
				break
			}
		}
		cpy := make([]int, len(nums))
		copy(cpy, nums)
		cpy = append(cpy, smallest)
		taken[smallest] = struct{}{}
		get_combos(taken, cpy, combos)
		delete(taken, smallest)
	} else {
		for i := 1; i < 8; i++ {
			if _, ok := taken[i]; !ok {
				cpy := make([]int, len(nums))
				copy(cpy, nums)
				cpy = append(cpy, i)
				taken[i] = struct{}{}
				get_combos(taken, cpy, combos)
				delete(taken, i)
			}
		}
	}
}

func init() {
	combos = [][]int{}
	taken := map[int]struct{}{0: struct{}{}}
	get_combos(taken, []int{0}, &combos)
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
			if j == i {
				continue
			} else if j < 1 {
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

	min := basic_strain
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

func resonance_with_crystals(data []string, neighbours int) int {
	portals := parse_data(data, neighbours)
	basic_strain := get_strain(portals)
	add_shortcuts(portals, data)

	min := basic_strain
	interval := len(portals) / 8

	for start := 1; start < 8; start++ {
		new_portals := make([]int, 8)
		for i := 0; i < 8; i++ {
			new_portals[i] = start + interval*i
			if new_portals[i] >= len(portals) {
				new_portals[i] -= len(portals) - 1
			}
		}

		for _, combo := range combos {
			for j := 0; j < 8; j += 2 {
				from := new_portals[combo[j]]
				to := new_portals[combo[j+1]]
				portals[from].connections = append(portals[from].connections, to)
				portals[to].connections = append(portals[to].connections, from)
			}

			strain := get_strain(portals)
			if strain < min {
				min = strain
			}
			for j := 0; j < 8; j++ {
				id := new_portals[j]
				portals[id].connections = portals[id].connections[:len(portals[id].connections)-1]
			}
		}
	}

	return basic_strain - min
}

func get_strain(portals []Portal) int {
	l := len(portals)

	dist := make([][]int, l)
	for i := range dist {
		dist[i] = make([]int, l)
	}

	for i := 1; i < l; i++ {
		for j := 1; j < l; j++ {
			if i == j {
				dist[i][j] = 0
				continue
			}
			dist[i][j] = 99999
			for _, conn := range portals[i].connections {
				if conn == j {
					dist[i][j] = 1
					break
				}
			}
		}
	}

	for k := 1; k < l; k++ {
		for i := 1; i < l; i++ {
			for j := 1; j < l; j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	total := 0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			total += dist[i][j]
		}
	}
	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 11, 1

	data := loader.GetStrings()
	part1 := resonance(data, 1)

	loader.Part = 2
	data = loader.GetStrings()
	part2 := resonance_extra_connection(data, 3)

	loader.Part = 3
	data = loader.GetStrings()
	part3 := resonance_with_crystals(data, 3)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
