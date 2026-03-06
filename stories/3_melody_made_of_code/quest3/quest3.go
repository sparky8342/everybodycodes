package quest3

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Node struct {
	id           int
	plug         string
	left_socket  string
	right_socket string
	plug_conn    *Node
	left_conn    *Node
	right_conn   *Node
}

func parse_value(line string) string {
	parts := strings.Split(line, "=")
	return parts[1]
}

func parse_line(line string) *Node {
	node := &Node{}

	parts := strings.Split(line, ", ")
	id_str := parse_value(parts[0])
	id, err := strconv.Atoi(id_str)
	if err != nil {
		panic(err)
	}
	node.id = id
	node.plug = parse_value(parts[1])
	node.left_socket = parse_value(parts[2])
	node.right_socket = parse_value(parts[3])

	return node
}

func add_node(node *Node, to_add *Node) bool {
	if node.left_conn != nil {
		if add_node(node.left_conn, to_add) {
			return true
		}
	} else if node.left_socket == to_add.plug {
		node.left_conn = to_add
		to_add.plug_conn = node
		return true
	}

	if node.right_conn != nil {
		if add_node(node.right_conn, to_add) {
			return true
		}
	} else if node.right_socket == to_add.plug {
		node.right_conn = to_add
		to_add.plug_conn = node
		return true
	}

	return false
}

func read(node *Node, ids *[]int) {
	if node.left_conn != nil {
		read(node.left_conn, ids)
	}

	*ids = append(*ids, node.id)

	if node.right_conn != nil {
		read(node.right_conn, ids)
	}
}

func read_tree(root *Node) int {
	ids := []int{}
	read(root, &ids)
	checksum := 0
	for i := 0; i < len(ids); i++ {
		checksum += ids[i] * (i + 1)
	}
	return checksum
}

func parse_data(data []string) *Node {
	root := parse_line(data[0])

	for i := 1; i < len(data); i++ {
		node := parse_line(data[i])
		add_node(root, node)
	}

	return root
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "3", 3, 1

	data := loader.GetStrings()
	root := parse_data(data)
	part1 := read_tree(root)

	fmt.Printf("%d %d %d\n", part1, -1, -1)
}
