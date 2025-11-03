package quest1

import (
	"fmt"
	"loader"
	"strings"
)

func get_name(data []string) string {
	names := strings.Split(data[0], ",")
	instructions := strings.Split(data[2], ",")

	pos := 0
	for _, ins := range instructions {
		n := int(ins[1] - '0')
		if ins[0] == 'L' {
			pos -= n
			if pos < 0 {
				pos = 0
			}
		} else if ins[0] == 'R' {
			pos += n
			if pos >= len(names) {
				pos = len(names) - 1
			}
		}
	}

	return names[pos]
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 1, 1

	data := loader.GetStrings()
	part1 := get_name(data)

	fmt.Println(part1)
}
