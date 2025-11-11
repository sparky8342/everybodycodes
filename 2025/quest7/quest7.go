package quest7

import (
	"fmt"
	"loader"
	"strings"
)

type Rules map[byte]map[byte]struct{}

func parse_data(data []string) ([]string, Rules) {
	names := strings.Split(data[0], ",")

	rules := Rules{}

	for i := 2; i < len(data); i++ {
		parts := strings.Split(data[i], " > ")
		pre := parts[0][0]
		rules[pre] = map[byte]struct{}{}
		for _, letter := range strings.Split(parts[1], ",") {
			rules[pre][letter[0]] = struct{}{}
		}
	}

	return names, rules
}

func valid_name(name string, rules Rules) bool {
	for i := 0; i < len(name)-1; i++ {
		if _, ok := rules[name[i]][name[i+1]]; !ok {
			return false
		}
	}
	return true
}

func get_valid_name(names []string, rules Rules) string {
	for _, name := range names {
		if valid_name(name, rules) {
			return name
		}
	}
	return ""
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 7, 1

	data := loader.GetStrings()
	names, rules := parse_data(data)
	part1 := get_valid_name(names, rules)

	fmt.Printf("%s %s %s\n", part1, "", "")
}
