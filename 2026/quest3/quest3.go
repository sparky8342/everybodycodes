package quest3

import (
	"fmt"
	"loader"
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
		p := 3
		for i := 4; i <= 6; i++ {
			total += pow(int(line[i]-'0'), p) * int(line[i+5]-'0')
			p--
		}
		total += int(line[12] - '0')
	}
	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 3, 1

	data := loader.GetStrings()
	part1 := score(data)

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
