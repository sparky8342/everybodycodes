package quest2

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Node struct {
	value  int
	symbol byte
	left   *Node
	right  *Node
}

func add_node(node *Node, new_node *Node) {
	if new_node.value < node.value {
		if node.left == nil {
			node.left = new_node
		} else {
			add_node(node.left, new_node)
		}
	} else if new_node.value > node.value {
		if node.right == nil {
			node.right = new_node
		} else {
			add_node(node.right, new_node)
		}
	}
}

func bfs(head *Node) []byte {
	queue := []*Node{head}

	longest_symbols := []byte{}

	for len(queue) > 0 {
		l := len(queue)

		symbols := []byte{}
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			symbols = append(symbols, node.symbol)
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}

		if len(symbols) > len(longest_symbols) {
			longest_symbols = symbols
		}
	}

	return longest_symbols
}

func parse_line(line string) (*Node, *Node) {
	parts := strings.Split(line, " ")
	left := strings.Split(parts[2], "=")[1]
	right := strings.Split(parts[3], "=")[1]

	nodes := make([]Node, 2)

	for i, str := range []string{left, right} {
		pair := strings.Split(str, ",")
		n, err := strconv.Atoi(pair[0][1:])
		if err != nil {
			panic(err)
		}
		nodes[i].value = n
		nodes[i].symbol = pair[1][0]
	}

	return &nodes[0], &nodes[1]
}

func process_data(data []string) string {
	left_head, right_head := parse_line(data[0])

	for i := 1; i < len(data); i++ {
		left_node, right_node := parse_line(data[i])
		add_node(left_head, left_node)
		add_node(right_head, right_node)
	}

	return string(append(bfs(left_head), bfs(right_head)...))
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "1", 2, 1

	data := loader.GetStrings()
	part1 := process_data(data)

	part2, part3 := -1, -1
	fmt.Printf("%s %d %d\n", part1, part2, part3)
}
