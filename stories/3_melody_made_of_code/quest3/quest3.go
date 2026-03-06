package quest3

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Node struct {
	id                  int
	plug_colour         string
	plug_shape          string
	left_socket_colour  string
	left_socket_shape   string
	right_socket_colour string
	right_socket_shape  string
	plug_conn           *Node
	left_conn           *Node
	right_conn          *Node
}

func parse_value(line string) (string, string) {
	parts := strings.Split(line, "=")
	parts = strings.Split(parts[1], " ")
	return parts[0], parts[1]
}

func parse_line(line string) *Node {
	node := &Node{}

	parts := strings.Split(line, ", ")
	id_str := strings.Split(parts[0], "=")[1]
	id, err := strconv.Atoi(id_str)
	if err != nil {
		panic(err)
	}
	node.id = id
	node.plug_colour, node.plug_shape = parse_value(parts[1])
	node.left_socket_colour, node.left_socket_shape = parse_value(parts[2])
	node.right_socket_colour, node.right_socket_shape = parse_value(parts[3])

	return node
}

func add_node(node *Node, to_add *Node, strong bool) bool {
	if node.left_conn != nil {
		if add_node(node.left_conn, to_add, strong) {
			return true
		}
	} else {
		colour := node.left_socket_colour == to_add.plug_colour
		shape := node.left_socket_shape == to_add.plug_shape
		if (strong && colour && shape) || (!strong && (colour || shape)) {
			node.left_conn = to_add
			to_add.plug_conn = node
			return true
		}
	}

	if node.right_conn != nil {
		if add_node(node.right_conn, to_add, strong) {
			return true
		}
	} else {
		colour := node.right_socket_colour == to_add.plug_colour
		shape := node.right_socket_shape == to_add.plug_shape
		if (strong && colour && shape) || (!strong && (colour || shape)) {
			node.right_conn = to_add
			to_add.plug_conn = node
			return true
		}
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

func parse_data(data []string, strong bool) *Node {
	root := parse_line(data[0])

	for i := 1; i < len(data); i++ {
		node := parse_line(data[i])
		add_node(root, node, strong)
	}

	return root
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "3", 3, 1

	data := loader.GetStrings()
	root := parse_data(data, true)
	part1 := read_tree(root)

	loader.Part = 2
	data = loader.GetStrings()
	root = parse_data(data, false)
	part2 := read_tree(root)

	fmt.Printf("%d %d %d\n", part1, part2, -1)
}
