package quest6

import (
	"fmt"
	"loader"
	"strings"
)

type Node struct {
	name     string
	children []*Node
	apple    bool
}

func parse_data(data []string) *Node {
	nodes := map[string]*Node{}

	for _, line := range data {
		parts := strings.Split(line, ":")
		parent := parts[0]

		if parent == "BUG" || parent == "ANT" {
			continue
		}

		if _, ok := nodes[parent]; !ok {
			nodes[parent] = &Node{name: parent}
		}

		children := strings.Split(parts[1], ",")

		for _, name := range children {
			if name == "BUG" || name == "ANT" {
				continue
			}
			if name == "@" {
				nodes[parent].apple = true
				continue
			}
			if _, ok := nodes[name]; !ok {
				nodes[name] = &Node{name: name}
			}
			nodes[parent].children = append(nodes[parent].children, nodes[name])
		}
	}

	return nodes["RR"]
}

func dfs(node *Node, path []string, path_lengths map[int]int, paths map[int]string, compact bool) {
	if node.apple {
		str := strings.Join(path, "")
		path_lengths[len(str)]++
		paths[len(str)] = str
	}
	for _, child_node := range node.children {
		if compact {
			dfs(child_node, append(path, child_node.name[0:1]), path_lengths, paths, compact)
		} else {
			dfs(child_node, append(path, child_node.name), path_lengths, paths, compact)
		}
	}
}

func find_path(root *Node, compact bool) string {
	path_lengths := map[int]int{}
	paths := map[int]string{}
	dfs(root, []string{}, path_lengths, paths, compact)
	for length, amount := range path_lengths {
		if amount == 1 {
			if compact {
				return "R" + paths[length] + "@"
			} else {
				return "RR" + paths[length] + "@"
			}
		}
	}
	return ""
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 6, 1

	data := loader.GetStrings()
	root := parse_data(data)
	part1 := find_path(root, false)

	loader.Part = 2
	data = loader.GetStrings()
	root = parse_data(data)
	part2 := find_path(root, true)

	loader.Part = 3
	data = loader.GetStrings()
	root = parse_data(data)
	part3 := find_path(root, true)

	fmt.Printf("%s %s %s\n", part1, part2, part3)
}
