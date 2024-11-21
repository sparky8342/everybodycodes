package quest6

import (
	"fmt"
	"loader"
	"strings"
)

type Node struct {
	name     string
	parent   *Node
	children []*Node
	apple    bool
}

func parse_data(data []string) *Node {
	nodes := map[string]*Node{}

	for _, line := range data {
		parts := strings.Split(line, ":")
		parent := parts[0]
		children := strings.Split(parts[1], ",")

		for _, name := range append([]string{parent}, children...) {
			if name == "@" {
				continue
			}
			if _, ok := nodes[name]; !ok {
				nodes[name] = &Node{name: name}
			}
		}

		for _, child := range children {
			if child == "@" {
				nodes[parent].apple = true
				continue
			}
			nodes[child].parent = nodes[parent]
			nodes[parent].children = append(nodes[parent].children, nodes[child])
		}
	}

	return nodes["RR"]
}

func dfs(node *Node, path []string, path_lengths map[int]int, paths map[int]string) {
	if node.apple {
		str := strings.Join(path, "")
		path_lengths[len(str)]++
		paths[len(str)] = str
	}
	for _, child_node := range node.children {
		dfs(child_node, append(path, child_node.name), path_lengths, paths)
	}
}

func find_path(root *Node) string {
	path_lengths := map[int]int{}
	paths := map[int]string{}
	dfs(root, []string{}, path_lengths, paths)
	for length, amount := range path_lengths {
		if amount == 1 {
			return "RR" + paths[length] + "@"
		}
	}
	return ""
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 6, 1

	data := loader.GetStrings()
	root := parse_data(data)
	part1 := find_path(root)

	part2, part3 := -1, -1
	fmt.Printf("%s %d %d\n", part1, part2, part3)
}
