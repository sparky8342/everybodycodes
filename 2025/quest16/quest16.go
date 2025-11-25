package quest16

import (
	"fmt"
	"loader"
	"math"
	"strconv"
	"strings"
)

func parse_data(data []byte) []int {
	parts := strings.Split(string(data), ",")
	nums := make([]int, len(parts))
	for i, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}

func build_wall(nums []int, columns int) int {
	blocks := 0
	for _, n := range nums {
		blocks += columns / n
	}
	return blocks
}

func compare_walls(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func find_pattern(wall []int) ([]int, int) {
	l := len(wall)

	pattern := []int{}
	test_wall := make([]int, l)

	n := 0

outer:
	for !compare_walls(test_wall, wall) {
		n++

		cpy := make([]int, l)
		copy(cpy, test_wall)

		for i := n - 1; i < l; i += n {
			cpy[i]++
			if cpy[i] > wall[i] {
				continue outer
			}
		}

		pattern = append(pattern, n)
		test_wall = cpy
	}

	product := 1
	for _, n := range pattern {
		product *= n
	}

	return pattern, product

}

func find_wall_length(wall []int, blocks int) int {
	pattern, _ := find_pattern(wall)

	min := 1
	max := math.MaxInt64

	for min+1 < max {
		mid := min + (max-min)/2
		blocks_used := build_wall(pattern, mid)
		if blocks_used < 0 || blocks_used > blocks { // < 0 is needed because of int64 overflows
			max = mid
		} else if blocks_used < blocks {
			min = mid
		}
	}

	return min
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 16, 1

	data := loader.GetOneLine()
	nums := parse_data(data)
	part1 := build_wall(nums, 90)

	loader.Part = 2
	data = loader.GetOneLine()
	nums = parse_data(data)
	_, part2 := find_pattern(nums)

	loader.Part = 3
	data = loader.GetOneLine()
	nums = parse_data(data)
	part3 := find_wall_length(nums, 202520252025000)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
