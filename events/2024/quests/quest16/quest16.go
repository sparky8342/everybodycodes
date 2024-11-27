package quest16

import (
	"fmt"
	"loader"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	wheels    [][]string
	rotations []int
	positions []int
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

	machine.positions = make([]int, len(rot))

	return machine
}

func (m *Machine) reset() {
	for i := 0; i < len(m.positions); i++ {
		m.positions[i] = 0
	}
}

func (m *Machine) spin(amount int) {
	for i := 0; i < len(m.positions); i++ {
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
	sequences := map[string][]int{}

	for i := 0; i < amount; i++ {
		for j := 0; j < len(m.positions); j++ {
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

		pos := make([]string, len(m.positions))
		for i, position := range m.positions {
			pos[i] = strconv.Itoa(position)
		}
		state := strings.Join(pos, "")

		if last, ok := sequences[state]; ok {
			diff := total - last[1]

			left := amount - i
			step := i - last[0]
			steps := left / step

			total += diff * steps
			i += steps * step
			sequences = map[string][]int{}
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

	part3 := -1
	fmt.Printf("%s %d %d\n", part1, part2, part3)
}
