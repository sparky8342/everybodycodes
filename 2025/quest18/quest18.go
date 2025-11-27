package quest18

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
	"math/rand/v2"
	"sort"
)

type Plant struct {
	id        int
	thickness int
	children  []*Plant
	branches  []int
}

type TestCase struct {
	settings Bitfield
	energy   int
}

type Bitfield struct {
	field1 uint64
	field2 uint64
}

func (b *Bitfield) get(n int) bool {
	if n > 63 {
		return b.field2&(1<<(n-63)) != 0
	} else {
		return b.field1&(1<<n) != 0
	}
}

func (b *Bitfield) set(n int) {
	if n > 63 {
		b.field2 = b.field2 | (1 << (n - 63))
	} else {
		b.field1 = b.field1 | (1 << n)
	}
}

func (b *Bitfield) clear(n int) {
	if n > 63 {
		b.field2 = b.field2 ^ (1 << (n - 63))
	} else {
		b.field1 = b.field1 ^ (1 << n)
	}
}

func (b *Bitfield) inc() {
	b.field1++
	if b.field1 == 0 {
		b.field2++
	}
}

func parse_data(data []string) (*Plant, []*Plant, []Bitfield) {
	var plant *Plant
	plants := map[int]*Plant{}
	free_plants := []*Plant{}
	test_cases := []Bitfield{}

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
			test_case := Bitfield{}
			for i := 0; i < len(parts); i++ {
				if parts[i][0] == '1' {
					test_case.set(i)
				}
			}
			test_cases = append(test_cases, test_case)
		}
	}

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

func run_test_case(top *Plant, free_plants []*Plant, test_case Bitfield) int {
	for i := 0; i < len(free_plants); i++ {
		if test_case.get(i) {
			free_plants[i].thickness = 1
		} else {
			free_plants[i].thickness = 0
		}
	}
	return top.energy()
}

func run_test_cases(top *Plant, free_plants []*Plant, test_cases []Bitfield) int {
	total := 0
	for _, test_case := range test_cases {
		total += run_test_case(top, free_plants, test_case)
	}
	return total
}

func next_gen(a Bitfield, b Bitfield, l int) Bitfield {
	r := rand.IntN(l)
	child := Bitfield{}
	for i := 0; i < r; i++ {
		if a.get(i) {
			child.set(i)
		}
	}
	for i := r; i < l; i++ {
		if b.get(i) {
			child.set(i)
		}
	}
	return child
}

func mutation(n Bitfield, l int) Bitfield {
	r := rand.IntN(10)
	if r == 0 {
		r = rand.IntN(l)
		if n.get(r) {
			n.clear(r)
		} else {
			n.set(r)
		}
	}
	return n
}

func max_energy(top *Plant, free_plants []*Plant) int {
	// basic genetic algorithm to find the best settings for the free_plants (branches)

	amount := 100

	lookup := []int{0}
	inc := 1
	for len(lookup) <= amount {
		lookup = append(lookup, lookup[len(lookup)-1]+inc)
		inc++
	}
	max_rand := lookup[len(lookup)-1]

	// initial cases
	test_cases := make([]TestCase, amount)

	for i := 0; i < amount; i++ {
		for j := 0; j < len(free_plants); j++ {
			n := rand.IntN(2)
			if n == 1 {
				test_cases[i].settings.set(j)
			}
		}
	}

	for i := 0; i < amount; i++ {
		test_cases[i].energy = run_test_case(top, free_plants, test_cases[i].settings)
	}

	sort.Slice(test_cases, func(i, j int) bool {
		return test_cases[i].energy < test_cases[j].energy
	})

	for i := 0; i < 100000; i++ {
		// create next generation - select 2 with bias on the higher energy cases,
		// create a child and then a random mutation

		// doesn't always find correct answer, maybe better to run until all cases have the same (max) energy

		next_cases := make([]TestCase, amount)
		for i := 0; i < amount; i++ {
			selections := []int{}
			for i := 0; i < 2; i++ {
				r := rand.IntN(max_rand)
				for i, n := range lookup {
					if r < n {
						selections = append(selections, i-1)
						break
					}
				}
			}
			next_case := TestCase{settings: mutation(next_gen(test_cases[selections[0]].settings, test_cases[selections[1]].settings, len(free_plants)), len(free_plants))}
			next_case.energy = run_test_case(top, free_plants, next_case.settings)
			next_cases[i] = next_case
		}

		test_cases = next_cases
		sort.Slice(test_cases, func(i, j int) bool {
			return test_cases[i].energy < test_cases[j].energy
		})

	}

	return test_cases[amount-1].energy
}

func energy_diff(top *Plant, free_plants []*Plant, test_cases []Bitfield) int {
	max := max_energy(top, free_plants)
	total_diff := 0
	for _, test_case := range test_cases {
		for i := 0; i < len(free_plants); i++ {
			if test_case.get(i) {
				free_plants[i].thickness = 1
			} else {
				free_plants[i].thickness = 0
			}
		}
		energy := top.energy()
		if energy > 0 {
			total_diff += max - energy
		}
		//	fmt.Println(test_case, top.energy())
	}
	return total_diff
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

	loader.Part = 3
	data = loader.GetStrings()
	top, free_plants, test_cases = parse_data(data)
	part3 := energy_diff(top, free_plants, test_cases)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
