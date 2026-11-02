package quest6

import (
	"fmt"
	"loader"
	"strings"
)

type Ingredients map[string][]string

func parse_data(data []string) Ingredients {
	ingredients := Ingredients{}
	for _, line := range data {
		ing := line[0:4]
		ingredients[ing] = []string{}
		for _, requirement := range strings.Split(line[5:], ",") {
			ingredients[ing] = append(ingredients[ing], requirement)
		}
	}
	return ingredients
}

func find_no_requirements(ingredients Ingredients, name string) int {
	requirements_found := map[string]struct{}{}
	dfs(ingredients, name, name, requirements_found)
	return len(requirements_found)
}

func dfs(ingredients Ingredients, name string, current_name string, requirements_found map[string]struct{}) {
	for _, requirement := range ingredients[current_name] {
		if requirement == name {
			continue
		}
		if _, ok := requirements_found[requirement]; !ok {
			requirements_found[requirement] = struct{}{}
			dfs(ingredients, name, requirement, requirements_found)
		}
	}
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 6, 1

	data := loader.GetStrings()
	ingredients := parse_data(data)
	part1 := find_no_requirements(ingredients, "NULL")

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
