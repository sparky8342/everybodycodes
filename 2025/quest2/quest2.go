package quest2

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type pair [2]int

func parse_data(data []byte) pair {
	parts := strings.Split(string(data), "=")
	num_parts := strings.Split(parts[1], ",")

	n1, err := strconv.Atoi(num_parts[0][1:])
	if err != nil {
		panic(err)
	}
	n2, err := strconv.Atoi(num_parts[1][:len(num_parts[1])-1])
	if err != nil {
		panic(err)
	}

	return pair{n1, n2}
}

func add(a pair, b pair) pair {
	return pair{a[0] + b[0], a[1] + b[1]}
}

func multiply(a pair, b pair) pair {
	return pair{a[0]*b[0] - a[1]*b[1], a[0]*b[1] + b[0]*a[1]}
}

func divide(a pair, b pair) pair {
	return pair{a[0] / b[0], a[1] / b[1]}
}

func calculate_part1(A pair) string {
	result := pair{0, 0}
	for i := 0; i < 3; i++ {
		result = multiply(result, result)
		result = divide(result, pair{10, 10})
		result = add(result, A)
	}
	return fmt.Sprintf("[%d,%d]", result[0], result[1])
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 2, 1

	data := loader.GetOneLine()
	A := parse_data(data)
	part1 := calculate_part1(A)

	fmt.Printf("%s %s %s\n", part1, "", "")
}
