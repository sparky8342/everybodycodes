package quest11

import (
	"fmt"
	"loader"
	"math"
	"strings"
)

func parse_data(data []string) map[string][]string {
	rules := map[string][]string{}
	for _, line := range data {
		parts := strings.Split(line, ":")
		source := parts[0]
		rules[source] = strings.Split(parts[1], ",")
	}
	return rules
}

func day(rules map[string][]string, population map[string]uint) map[string]uint {
	next := map[string]uint{}
	for typ, amount := range population {
		for _, dest := range rules[typ] {
			next[dest] += amount
		}
	}
	return next
}

func days(rules map[string][]string, start string, amount int) uint {
	population := map[string]uint{}
	population[start] = 1
	for i := 0; i < amount; i++ {
		population = day(rules, population)
	}
	var count uint = 0
	for _, n := range population {
		count += n
	}
	return count
}

func all_starts(rules map[string][]string) uint {
	var min uint = math.MaxInt64
	var max uint = 0
	for source := range rules {
		count := days(rules, source, 20)
		if count < min {
			min = count
		} else if count > max {
			max = count
		}
	}
	return max - min
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 11, 1

	data := loader.GetStrings()
	rules := parse_data(data)
	part1 := days(rules, "A", 4)

	loader.Part = 2
	data = loader.GetStrings()
	rules = parse_data(data)
	part2 := days(rules, "Z", 10)

	loader.Part = 3
	data = loader.GetStrings()
	rules = parse_data(data)
	part3 := all_starts(rules)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
