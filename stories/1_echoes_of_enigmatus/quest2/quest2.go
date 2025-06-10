package quest2

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Node struct {
	id     int
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

func parse_add(line string) (*Node, *Node) {
	parts := strings.Split(line, " ")
	id, err := strconv.Atoi(strings.Split(parts[1], "=")[1])
	if err != nil {
		panic(err)
	}
	left := strings.Split(parts[2], "=")[1]
	right := strings.Split(parts[3], "=")[1]

	nodes := make([]Node, 2)

	for i, str := range []string{left, right} {
		pair := strings.Split(str, ",")
		n, err := strconv.Atoi(pair[0][1:])
		if err != nil {
			panic(err)
		}
		nodes[i].id = id
		nodes[i].value = n
		nodes[i].symbol = pair[1][0]
	}

	return &nodes[0], &nodes[1]
}

func process_data(data []string, swapmode int) string {
	left_nodes := map[int]*Node{}
	right_nodes := map[int]*Node{}

	left_head, right_head := parse_add(data[0])
	left_nodes[left_head.id] = left_head
	right_nodes[right_head.id] = right_head

	for i := 1; i < len(data); i++ {
		if data[i][0:3] == "ADD" {
			left_node, right_node := parse_add(data[i])
			add_node(left_head, left_node)
			add_node(right_head, right_node)
			left_nodes[left_node.id] = left_node
			right_nodes[right_node.id] = right_node
		} else if data[i][0:4] == "SWAP" {
			id, err := strconv.Atoi(data[i][5:])
			if err != nil {
				panic(err)
			}
			if swapmode == 1 {
				left_nodes[id].value, right_nodes[id].value = right_nodes[id].value, left_nodes[id].value
				left_nodes[id].symbol, right_nodes[id].symbol = right_nodes[id].symbol, left_nodes[id].symbol
			} else if swapmode == 2 {
				*left_nodes[id], *right_nodes[id] = *right_nodes[id], *left_nodes[id]
			}
		}
	}

	return string(append(bfs(left_head), bfs(right_head)...))
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "1", 2, 1

	data := loader.GetStrings()
	part1 := process_data(data, 1)

	loader.Part = 2
	data = loader.GetStrings()
	part2 := process_data(data, 1)

	loader.Part = 3
	data = loader.GetStrings()
	part3 := process_data(data, 2)

	fmt.Printf("%s %s %s\n", part1, part2, part3)
}
