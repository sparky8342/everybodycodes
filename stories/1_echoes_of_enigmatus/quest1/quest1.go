package quest1

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

func parse_data(data []string) []map[byte]int {
	param_lines := make([]map[byte]int, len(data))
	for i, line := range data {
		params := map[byte]int{}
		for _, str := range strings.Split(line, " ") {
			pair := strings.Split(str, "=")
			n, err := strconv.Atoi(pair[1])
			if err != nil {
				panic(err)
			}
			params[pair[0][0]] = n
		}
		param_lines[i] = params
	}
	return param_lines
}

func digits(n int) int {
	digits := 0
	for n > 0 {
		n /= 10
		digits++
	}
	return digits
}

func eni(n int, exp int, mod int) int {
	num := 1
	result := 0
	for i := 0; i < exp; i++ {
		num = (num * n) % mod
		n := num
		for j := 0; j < digits(result); j++ {
			n *= 10
		}
		result += n
	}
	return result
}

func highest_line(param_lines []map[byte]int) int {
	max := 0
	for _, params := range param_lines {
		A, B, C, X, Y, Z, M := params['A'], params['B'], params['C'], params['X'], params['Y'], params['Z'], params['M']
		result := eni(A, X, M) + eni(B, Y, M) + eni(C, Z, M)
		if result > max {
			max = result
		}
	}
	return max
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "1", 1, 1

	data := loader.GetStrings()
	param_lines := parse_data(data)

	part1 := highest_line(param_lines)

	part2, part3 := -1, -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
