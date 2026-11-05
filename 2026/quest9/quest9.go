package quest9

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Plant struct {
	id     int
	growth int
	limit  int
}

func parse_data(data []string) []Plant {
	plants := []Plant{}
	for _, line := range data {
		parts := strings.Split(line, ", ")
		nums := [3]int{}
		for i := 0; i < 3; i++ {
			num, err := strconv.Atoi(strings.Split(parts[i], "=")[1])
			if err != nil {
				panic(err)
			}
			nums[i] = num
		}
		plants = append(plants, Plant{
			id:     nums[0],
			growth: nums[1],
			limit:  nums[2],
		})
	}
	return plants
}

func growth(plant Plant) int {
	time := 1
	branches := []int{1}
	size := 1

	for size < plant.limit {
		new_branches := []int{}
		for i := 0; i < len(branches); i++ {
			branches[i]++
			size++
			if branches[i] == plant.growth {
				branches[i] /= 2
				new_branches = append(new_branches, branches[i], 1)
				size++
			}
		}
		branches = append(branches, new_branches...)
		time++
	}

	return time
}

func growth_time(plants []Plant) int {
	total := 0
	for _, plant := range plants {
		total += growth(plant)
	}
	return total
}

func growth_leaves_limit(plant Plant) int {
	time := 1
	branches := []int{1}
	leaves := 1

	for leaves < plant.limit {
		new_branches := []int{}
		for i := 0; i < len(branches); i++ {
			branches[i]++
			if branches[i] == plant.growth {
				branches[i] /= 2
				new_branches = append(new_branches, branches[i], 1)
				leaves++
			}
		}
		branches = append(branches, new_branches...)
		time++
	}

	return time
}

func growth_time_leaves_limit(plants []Plant) int {
	total := 0
	for _, plant := range plants {
		total += growth_leaves_limit(plant)
	}
	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 9, 1

	data := loader.GetStrings()
	plants := parse_data(data)
	part1 := growth_time(plants)

	loader.Part = 2
	data = loader.GetStrings()
	plants = parse_data(data)
	part2 := growth_time_leaves_limit(plants)

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
