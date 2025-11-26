package quest18

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Plant struct {
	id        int
	thickness int
	children  []*Plant
	branches  []int
}

func parse_data(data []string) (*Plant, []*Plant, [][]int) {
	var plant *Plant
	plants := map[int]*Plant{}
	free_plants := []*Plant{}
	test_cases := [][]int{}

	for _, row := range data {
		parts := strings.Split(row, " ")

		if parts[0] == "Plant" {
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			thickness, err := strconv.Atoi(parts[4][:len(parts[4])-1])
			if err != nil {
				panic(err)
			}

			plant = &Plant{id: id, thickness: thickness}
			plants[id] = plant
		} else if parts[0] == "-" && parts[1] == "free" {
			thickness, err := strconv.Atoi(parts[5])
			if err != nil {
				panic(err)
			}

			free_plant := &Plant{thickness: 1}
			plant.children = append(plant.children, free_plant)
			plant.branches = append(plant.branches, thickness)
			free_plants = append(free_plants, free_plant)
		} else if parts[0] == "-" && parts[1] == "branch" {
			child, err := strconv.Atoi(parts[4])
			if err != nil {
				panic(err)
			}
			thickness, err := strconv.Atoi(parts[7])
			if err != nil {
				panic(err)
			}

			plant.children = append(plant.children, plants[child])
			plant.branches = append(plant.branches, thickness)
		} else if parts[0] == "0" || parts[0] == "1" {
			test_case := make([]int, len(parts))
			for i, part := range parts {
				test_case[i] = int(part[0] - '0')
			}
			test_cases = append(test_cases, test_case)
		}
	}

	/*
		for _, p := range plants {
			fmt.Println(p.id, p.thickness, p.branches)
			for _, c := range p.children {
				if c == nil {
					fmt.Println("nil")
				} else {
					fmt.Println(c.id)
				}
			}
			fmt.Println()
		}
	*/

	return plant, free_plants, test_cases
}

func (p *Plant) energy() int {
	incoming := 0
	if len(p.children) == 0 {
		incoming = p.thickness
	} else {
		for i := 0; i < len(p.children); i++ {
			incoming += p.branches[i] * p.children[i].energy()
		}
	}
	if incoming < p.thickness {
		return 0
	} else {
		return incoming
	}
}

func run_test_cases(top *Plant, free_plants []*Plant, test_cases [][]int) int {
	total := 0
	for _, test_case := range test_cases {
		for i := 0; i < len(test_case); i++ {
			free_plants[i].thickness = test_case[i]
		}
		total += top.energy()
	}
	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 18, 1

	data := loader.GetStrings()
	top, _, _ := parse_data(data)
	part1 := top.energy()

	loader.Part = 2
	data = loader.GetStrings()
	top, free_plants, test_cases := parse_data(data)
	part2 := run_test_cases(top, free_plants, test_cases)

	fmt.Printf("%d %d %d\n", part1, part2, 0)
}
