package quest16

import (
	"fmt"
	"loader"
)

func sort_tablets(tablets_str string, duplicate int) (int, int) {
	tablets := []byte{}
	for i := 1; i <= duplicate; i++ {
		tablets = append(tablets, []byte(tablets_str)...)
	}

	flips := 0
	tablets_moved := 0

	for len(tablets) > 0 {
		start := 0
		end := len(tablets) - 1
		colour := tablets[0]

		next := []byte{}
		for start <= end {
			for start <= end && tablets[start] == colour {
				start++
			}
			for start <= end && tablets[end] != colour {
				next = append(next, tablets[end])
				end--
			}
			tablets_moved += end - start + 1
			if start <= end {
				flips++
			}

			for start <= end && tablets[end] == colour {
				end--
			}
			for start <= end && tablets[start] != colour {
				next = append(next, tablets[start])
				start++
			}
			tablets_moved += end - start + 1
			if start <= end {
				flips++
			}

		}
		tablets = next
	}

	return flips, tablets_moved
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 16, 1

	tablets := loader.GetStrings()[0]
	part1, _ := sort_tablets(tablets, 1)

	loader.Part = 2
	tablets = loader.GetStrings()[0]
	_, part2 := sort_tablets(tablets, 10)

	loader.Part = 3
	tablets = loader.GetStrings()[0]
	_, part3 := sort_tablets(tablets, 100)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
