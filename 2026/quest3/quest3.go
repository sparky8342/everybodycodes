package quest3

import (
	"fmt"
	"loader"
	"strings"
)

func pow(n, p int) int {
	ans := n
	for i := 1; i < p; i++ {
		ans *= n
	}
	return ans
}

func score(data []string) int {
	total := 0
	for _, line := range data {
		parts := strings.Split(line, " ")
		bases := parts[0][4:]
		number := parts[1]

		total += int(number[len(number)-1] - '0')

		p := 1
		var base_pos int
		if len(bases) == 1 {
			base_pos = 0
		} else {
			base_pos = len(bases) - 2
		}

		for i := len(number) - 2; i >= 0; i-- {
			total += pow(int(bases[base_pos]-'0'), p) * int(number[i]-'0')
			p++
			base_pos--
			if base_pos == -1 {
				base_pos = len(bases) - 1
			}
		}
	}
	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 3, 1

	data := loader.GetStrings()
	part1 := score(data)

	loader.Part = 2
	data = loader.GetStrings()
	part2 := score(data)

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
