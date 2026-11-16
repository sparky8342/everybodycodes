package quest16

import (
	"fmt"
	"loader"
)

func sort_tablets(tablets_str string) int {
	tablets := []byte(tablets_str)

	flips := 0
outer:
	for i := 0; i < len(tablets)-1; i++ {
		colour := tablets[i]
		p1 := i + 1
		for tablets[p1] == colour {
			p1++
			if p1 == len(tablets) {
				continue outer
			}
		}
		p2 := len(tablets) - 1
		for tablets[p2] != colour {
			p2--
		}

		if p1 < p2 {
			for p1 < p2 {
				tablets[p1], tablets[p2] = tablets[p2], tablets[p1]
				p1++
				p2--
			}
			flips++
		}
	}

	return flips
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 16, 1

	tablets := loader.GetStrings()[0]
	part1 := sort_tablets(tablets)

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
