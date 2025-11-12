package quest8

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type pair [2]int

func parse_data(data []byte) []int {
	num_strs := strings.Split(string(data), ",")
	nums := make([]int, len(num_strs))
	for i, str := range num_strs {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}

func centre_count(nails int, nums []int) int {
	half := nails / 2
	count := 0
	for i := 1; i < len(nums); i++ {
		a, b := nums[i-1], nums[i]
		if a > b {
			a, b = b, a
		}
		if a+half == b {
			count++
		}
	}
	return count
}

func knots(nails int, nums []int) int {
	lines := []pair{}
	for i := 0; i < len(nums)-1; i++ {
		a, b := nums[i], nums[i+1]
		if a > b {
			a, b = b, a
		}
		lines = append(lines, pair{a, b})
	}

	count := 0
	for i := 0; i < len(lines); i++ {
		a, b := lines[i][0], lines[i][1]

		for j := i + 1; j < len(lines); j++ {
			c, d := lines[j][0], lines[j][1]
			if (c < a || c > b) && (d > a && d < b) {
				count++
			}
			if (d < a || d > b) && (c > a && c < b) {
				count++
			}
		}
	}

	return count
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 8, 1

	data := loader.GetOneLine()
	nums := parse_data(data)
	part1 := centre_count(32, nums)

	loader.Part = 2
	data = loader.GetOneLine()
	nums = parse_data(data)
	part2 := knots(32, nums)

	fmt.Printf("%d %d %s\n", part1, part2, "")
}
