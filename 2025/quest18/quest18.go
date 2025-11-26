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

func parse_data(data []string) *Plant {
	var plant *Plant
	plants := map[int]*Plant{}

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

			plant.children = append(plant.children, nil)
			plant.branches = append(plant.branches, thickness)
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

	return plant
}

func (p *Plant) energy() int {
	incoming := 0
	for i := 0; i < len(p.children); i++ {
		if p.children[i] == nil {
			incoming += p.branches[i]
		} else {
			incoming += p.branches[i] * p.children[i].energy()
		}
	}
	if incoming < p.thickness {
		return 0
	} else {
		return incoming
	}
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 18, 1

	data := loader.GetStrings()
	top := parse_data(data)
	part1 := top.energy()

	fmt.Printf("%d %d %d\n", part1, 0, 0)
}
