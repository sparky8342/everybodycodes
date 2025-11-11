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

func get_valid_names(names []string, rules Rules) ([]string, int) {
	valid_names := []string{}
	index_sum := 0
	for i, name := range names {
		if valid_name(name, rules) {
			valid_names = append(valid_names, name)
			index_sum += i + 1
		}
	}
	return valid_names, index_sum
}

func dfs(name []byte, rules Rules, names map[string]struct{}) {
	l := len(name)
	if l >= 7 {
		str := string(name)
		if _, ok := names[str]; ok {
			return
		}
		names[str] = struct{}{}
		if l == 11 {
			return
		}
	}
	for letter := range rules[name[l-1]] {
		dfs(append(name, letter), rules, names)
	}
}

func possible_names(prefixes []string, rules Rules) int {
	valid_prefixes, _ := get_valid_names(prefixes, rules)
	names := map[string]struct{}{}
	for _, prefix := range valid_prefixes {
		dfs([]byte(prefix), rules, names)
	}
	return len(names)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 7, 1

	data := loader.GetStrings()
	names, rules := parse_data(data)
	valid_names, _ := get_valid_names(names, rules)
	part1 := valid_names[0]

	loader.Part = 2
	data = loader.GetStrings()
	names, rules = parse_data(data)
	_, part2 := get_valid_names(names, rules)

	loader.Part = 3
	data = loader.GetStrings()
	prefixes, rules := parse_data(data)
	part3 := possible_names(prefixes, rules)

	fmt.Printf("%s %d %d\n", part1, part2, part3)
}
