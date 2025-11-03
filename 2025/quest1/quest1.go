package quest1

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

func get_name(data []string, circular bool) string {
	names := strings.Split(data[0], ",")
	instructions := strings.Split(data[2], ",")

	pos := 0
	for _, ins := range instructions {
		n, err := strconv.Atoi(ins[1:])
		if err != nil {
			panic(err)
		}

		if ins[0] == 'L' {
			pos -= n
		} else if ins[0] == 'R' {
			pos += n
		}
		if circular {
			pos = pos % len(names)
			for pos < 0 {
				pos += len(names)
			}
		} else if pos < 0 {
			pos = 0
		} else if pos >= len(names) {
			pos = len(names) - 1
		}
	}

	return names[pos]
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 1, 1

	data := loader.GetStrings()
	part1 := get_name(data, false)

	loader.Part = 2
	data = loader.GetStrings()
	part2 := get_name(data, true)

	fmt.Printf("%s %s %s\n", part1, part2, "")

}
