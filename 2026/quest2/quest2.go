package quest2

import (
	"fmt"
	"loader"
	"sort"
	"strings"
)

type pair [2]int

func find_missing_rule(rules map[pair]bool, max int) (int, int) {
	for a := 1; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			if _, ok := rules[pair{a, b}]; !ok {
				return a, b
			}
		}
	}
	return -1, -1
}

func arrange_alloys(line string) int {
	rules := map[pair]bool{}

	max := 1
	for _, rule := range strings.Split(line, " ") {
		a := int(rule[0] - '0')
		b := int(rule[2] - '0')
		var result bool
		if rule[1] == '<' {
			result = false
		} else if rule[1] == '>' {
			result = true
		}

		rules[pair{a, b}] = result
		rules[pair{b, a}] = !result

		if a > max {
			max = a
		}
		if b > max {
			max = b
		}
	}

	expected_rules := 0
	for i := 1; i < max; i++ {
		expected_rules += i
	}
	expected_rules *= 2

	if len(rules) < expected_rules {
		a, b := find_missing_rule(rules, max)

		var val int
		var rule pair
		var order bool
		for rule, order = range rules {
			if rule[0] == a {
				val = rule[1]
				break
			}
		}

		order2 := rules[pair{val, b}]

		if order == order2 {
			rules[pair{a, b}] = order
			rules[pair{b, a}] = !order
		} else {
			fmt.Println("NOT MATCHING") 
			// this didn't happen? my thought is if it did,
			// a matching pair must be found instead
		}
	}

	alloys := []int{}
	for i := 1; i <= max; i++ {
		alloys = append(alloys, i)
	}
	sort.Slice(alloys, func(i, j int) bool {
		return rules[pair{alloys[i], alloys[j]}]
	})

	checksum := 0
	for _, alloy := range alloys {
		checksum = checksum*10 + alloy
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
	data = loader.GetStrings()
	part2 := arrange_groups(data)

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
