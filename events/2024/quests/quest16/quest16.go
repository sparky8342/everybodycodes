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

func (m *Machine) spin(amount int) {
	for i := 0; i < len(m.positions); i++ {
		m.positions[i] = (m.positions[i] + m.rotations[i]*amount) % len(m.wheels[i])
	}
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

	part2, part3 := -1, -1
	fmt.Printf("%s %d %d\n", part1, part2, part3)
}
