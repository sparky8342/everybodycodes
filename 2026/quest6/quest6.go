package quest6

import (
	"fmt"
	"loader"
	"strings"
)

type Ingredient struct {
	name     string
	parents  []string
	children []string
}

type Ingredients map[string]*Ingredient

func parse_data(data []string) Ingredients {
	ingredients := Ingredients{}
	for _, line := range data {
		names := append([]string{line[0:4]}, strings.Split(line[5:], ",")...)
		for _, name := range names {
			if _, ok := ingredients[name]; !ok {
				ingredients[name] = &Ingredient{
					name:     name,
					parents:  []string{},
					children: []string{},
				}
			}
		}
		for i := 1; i < len(names); i++ {
			ingredients[names[0]].children = append(ingredients[names[0]].children, names[i])
			ingredients[names[i]].parents = append(ingredients[names[i]].parents, names[0])
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
	for _, requirement := range ingredients[current_name].children {
		if requirement == name {
			continue
		}
		if _, ok := requirements_found[requirement]; !ok {
			requirements_found[requirement] = struct{}{}
			dfs(ingredients, name, requirement, requirements_found)
		}
	}
}

func find_no_products(ingredients Ingredients, name string) int {
	products_found := map[string]struct{}{}
	dfs_up(ingredients, name, name, products_found)
	return len(products_found)
}

func dfs_up(ingredients Ingredients, name string, current_name string, products_found map[string]struct{}) {
	for _, product := range ingredients[current_name].parents {
		if product == name {
			continue
		}
		if _, ok := products_found[product]; !ok {
			products_found[product] = struct{}{}
			dfs_up(ingredients, name, product, products_found)
		}
	}
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 6, 1

	data := loader.GetStrings()
	ingredients := parse_data(data)
	part1 := find_no_requirements(ingredients, "NULL")

	loader.Part = 2
	data = loader.GetStrings()
	ingredients = parse_data(data)
	part2 := find_no_products(ingredients, "NULL")

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
