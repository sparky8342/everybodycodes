package quest11

import (
	"fmt"
	"loader"
	"strings"
)

func parse_data(data []string) [][]int {
	rules := make([][]int, 26)

	for _, line := range data {
		parts := strings.Split(line, ":")
		source := parts[0]
		destination := []int{}
		for _, dest := range strings.Split(parts[1], ",") {
			destination = append(destination, int(dest[0]-65))
		}
		rules[int(source[0]-65)] = destination
	}

	return rules
}

func day(rules [][]int, population []int) []int {
	next := make([]int, 26)
	for i := 0; i < 26; i++ {
		if population[i] > 0 {
			for _, n := range rules[i] {
				next[n] += population[i]
			}
		}
	}
	return next
}

func days(rules [][]int, start int, amount int) int {
	population := make([]int, 26)
	population[start] = 1
	for i := 0; i < amount; i++ {
		population = day(rules, population)
	}
	count := 0
	for i := 0; i < 26; i++ {
		count += population[i]
	}
	return count
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 11, 1

	data := loader.GetStrings()
	rules := parse_data(data)
	part1 := days(rules, 0, 4)

	loader.Part = 2
	data = loader.GetStrings()
	rules = parse_data(data)
	part2 := days(rules, 25, 10)

	part3 := -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
