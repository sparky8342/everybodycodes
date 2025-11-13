package quest8

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type pair [2]int

func parse_data(data []byte) []pair {
	num_strs := strings.Split(string(data), ",")
	nums := make([]int, len(num_strs))
	for i, str := range num_strs {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}

	lines := []pair{}
	for i := 0; i < len(nums)-1; i++ {
		a, b := nums[i], nums[i+1]
		if a > b {
			a, b = b, a
		}
		lines = append(lines, pair{a, b})
	}

	return lines
}

func centre_count(nails int, lines []pair) int {
	half := nails / 2
	count := 0
	for _, line := range lines {
		if line[1]-line[0] == half {
			count++
		}
	}
	return count
}

func crossing_lines(line pair, lines []pair) int {
	count := 0
	a, b := line[0], line[1]

	for i := 0; i < len(lines); i++ {
		c, d := lines[i][0], lines[i][1]
		if (c < a || c > b) && (d > a && d < b) {
			count++
		}
		if (d < a || d > b) && (c > a && c < b) {
			count++
		}
	}

	return count
}

func knots(nails int, lines []pair) int {
	count := 0
	for i := 0; i < len(lines); i++ {
		count += crossing_lines(lines[i], lines)

	}
	return count / 2
}

func best_cut(nails int, lines []pair) int {
	max_threads := 0

	for i := 1; i <= nails; i++ {
		for j := i + 2; j <= nails; j++ {
			cut := pair{i, j}
			threads := crossing_lines(cut, lines)
			for _, line := range lines {
				if line[0] == i && line[1] == j {
					threads++
					break
				}
			}
			if threads > max_threads {
				max_threads = threads
			}
		}
	}

	return max_threads
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 8, 1

	data := loader.GetOneLine()
	lines := parse_data(data)
	part1 := centre_count(32, lines)

	loader.Part = 2
	data = loader.GetOneLine()
	lines = parse_data(data)
	part2 := knots(32, lines)

	loader.Part = 3
	data = loader.GetOneLine()
	lines = parse_data(data)
	part3 := best_cut(256, lines)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
