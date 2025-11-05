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

func calculate_part2(A pair) int {
	engraved := 0

	top_left := A
	bottom_right := add(A, pair{1000, 1000})

	for x := top_left[0]; x <= bottom_right[0]; x += 10 {
		for y := top_left[1]; y <= bottom_right[1]; y += 10 {
			point := pair{x, y}
			result := pair{0, 0}

			ok := true
			for i := 0; i < 100; i++ {
				result = multiply(result, result)
				result = divide(result, pair{100000, 100000})
				result = add(result, point)
				if result[0] < -1000000 || result[0] > 1000000 || result[1] < -1000000 || result[1] > 1000000 {
					ok = false
					break
				}
			}

			if ok {
				engraved++
			}
		}
	}

	return engraved
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 2, 1

	data := loader.GetOneLine()
	A := parse_data(data)
	part1 := calculate_part1(A)

	loader.Part = 2
	data = loader.GetOneLine()
	A = parse_data(data)
	part2 := calculate_part2(A)

	fmt.Printf("%s %d %s\n", part1, part2, "")
}
