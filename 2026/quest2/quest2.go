package quest2

import (
	"fmt"
	"loader"
	"sort"
	"strings"
)

func arrange_alloys(line string) int {
	rules := map[int]bool{}

	for _, rule := range strings.Split(line, " ") {
		a := int(rule[0] - '0')
		b := int(rule[2] - '0')
		var result bool
		if rule[1] == '<' {
			result = false
		} else if rule[1] == '>' {
			result = true
		}

		rules[a*10+b] = result
		rules[b*10+a] = !result
	}

	alloys := []int{1, 2, 3, 4, 5}
	sort.Slice(alloys, func(i, j int) bool {
		return rules[alloys[i]*10+alloys[j]]
	})

	checksum := 0
	for i := 0; i < 5; i++ {
		checksum = checksum*10 + alloys[i]
	}

	return checksum
}

func arrange_groups(data []string) int {
	total_checksum := 0
	for _, line := range data {
		total_checksum += arrange_alloys(line)
	}
	return total_checksum
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 2, 1

	data := loader.GetStrings()
	part1 := arrange_groups(data)

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
