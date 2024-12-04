package quest16

import (
	"fmt"
	"loader"
	"os"
	"strconv"
	"strings"
	"utils"
)

type Machine struct {
	wheels    [][]string
	rotations []int
	positions [10]int
}

func parse_data(data []string) Machine {
	machine := Machine{}

	rot := strings.Split(data[0], ",")
	machine.rotations = make([]int, len(rot))
	for i, r := range rot {
		n, err := strconv.Atoi(r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error %v\n", err)
			os.Exit(1)
		}
		machine.rotations[i] = n
	}

	machine.wheels = make([][]string, len(rot))
	for _, line := range data[2:] {
		wheel := 0
		for i := 0; i < len(line); i += 4 {
			cat := line[i : i+3]
			if cat != "   " {
				machine.wheels[wheel] = append(machine.wheels[wheel], cat)
			}
			wheel++
		}
	}

	return machine
}

func (m *Machine) reset() {
	for i := 0; i < len(m.positions); i++ {
		m.positions[i] = 0
	}
}

func (m *Machine) spin(amount int) {
	for i := 0; i < len(m.wheels); i++ {
		m.positions[i] = (m.positions[i] + m.rotations[i]*amount) % len(m.wheels[i])
	}
}

func (m *Machine) score() int {
	symbols := map[byte]int{}
	for i := 0; i < len(m.wheels); i++ {
		symbols[m.wheels[i][m.positions[i]][0]]++
		symbols[m.wheels[i][m.positions[i]][2]]++
	}
	score := 0
	for _, count := range symbols {
		if count >= 3 {
			score += count - 2
		}
	}
	return score
}

func (m *Machine) spin_with_score(amount int) int {
	total := 0
	sequences := map[int][]int{}

	for i := 0; i < amount; i++ {
		for j := 0; j < len(m.wheels); j++ {
			m.positions[j] = (m.positions[j] + m.rotations[j]) % len(m.wheels[j])
		}

		symbols := map[byte]int{}
		for i := 0; i < len(m.wheels); i++ {
			symbols[m.wheels[i][m.positions[i]][0]]++
			symbols[m.wheels[i][m.positions[i]][2]]++
		}
		score := 0
		for _, count := range symbols {
			if count >= 3 {
				score += count - 2
			}
		}
		total += score

		state := 0
		for i := 0; i < len(m.wheels); i++ {
			state = state*100 + m.positions[i]
		}

		if last, ok := sequences[state]; ok {
			diff := total - last[1]

			left := amount - i
			step := i - last[0]
			steps := left / step

			total += diff * steps
			i += steps * step
			sequences = map[int][]int{}
		} else {
			sequences[state] = []int{i, total}
		}

	}

	return total
}

func (m *Machine) display() string {
	cats := make([]string, len(m.wheels))
	for i := 0; i < len(cats); i++ {
		cats[i] = m.wheels[i][m.positions[i]]
	}
	return strings.Join(cats, " ")
}

func (m *Machine) cache_key(pulls int, score int) int {
	n := 0
	for _, p := range m.positions {
		n = n*100 + p
	}
	n = n*1000 + pulls
	n = n*1000 + score
	return n
}

func (m *Machine) left_spin_up() {
	for i := 0; i < len(m.wheels); i++ {
		m.positions[i] = (m.positions[i] - 1) % len(m.wheels[i])
	}
}

func (m *Machine) left_spin_down() {
	for i := 0; i < len(m.wheels); i++ {
		m.positions[i] = (m.positions[i] + 1) % len(m.wheels[i])
	}
}

func (m *Machine) dfs(pulls int, score int, cache map[int]int, min_mode bool) int {
	key := m.cache_key(pulls, score)

	if pulls == 0 {
		return score
	}

	if val, ok := cache[key]; ok {
		return val
	}

	positions := m.positions

	vals := []int{}

	// no left spin
	m.spin(1)
	vals = append(vals, m.dfs(pulls-1, score+m.score(), cache, min_mode))

	// left spin up
	m.positions = positions
	m.left_spin_up()
	m.spin(1)
	vals = append(vals, m.dfs(pulls-1, score+m.score(), cache, min_mode))

	//left spin down
	m.positions = positions
	m.left_spin_down()
	m.spin(1)
	vals = append(vals, m.dfs(pulls-1, score+m.score(), cache, min_mode))

	var val int
	if min_mode {
		val = utils.Min(vals)
	} else {
		val = utils.Max(vals)
	}

	cache[key] = val
	return val
}

func (m *Machine) minmax(pulls int) (int, int) {
	cache := map[int]int{}
	min := m.dfs(pulls, 0, cache, true)

	m.reset()
	cache = map[int]int{}
	max := m.dfs(pulls, 0, cache, false)

	return min, max
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 16, 1

	data := loader.GetStrings()
	machine := parse_data(data)
	machine.spin(100)
	part1 := machine.display()

	loader.Part = 2
	data = loader.GetStrings()
	machine = parse_data(data)
	part2 := machine.spin_with_score(202420242024)

	loader.Part = 3
	data = loader.GetStrings()
	machine = parse_data(data)
	min, max := machine.minmax(256)

	fmt.Printf("%s\n%d\n%d %d\n", part1, part2, max, min)
}
