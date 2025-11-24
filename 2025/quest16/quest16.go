package quest16

import (
	"fmt"
	"loader"
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

func find_pattern(wall []int) int {
	l := len(wall)

	pattern := []int{}
	test_wall := make([]int, l)

	n := 1

outer:
	for {

		ok := true
		for i := n - 1; i < l; i += n {
			test_wall[i]++
			if test_wall[i] > wall[i] {
				ok = false
			}
		}

		if ok {
			pattern = append(pattern, n)
		} else {
			for i := n - 1; i < l; i += n {
				test_wall[i]--
			}
		}

		n++

		for i := 0; i < l; i++ {
			if test_wall[i] != wall[i] {
				continue outer
			}
		}

		break
	}

	product := 1
	for _, n := range pattern {
		product *= n
	}

	return product

}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 16, 1

	data := loader.GetOneLine()
	nums := parse_data(data)
	part1 := build_wall(nums, 90)

	loader.Part = 2
	data = loader.GetOneLine()
	nums = parse_data(data)
	part2 := find_pattern(nums)

	fmt.Printf("%d %d %d\n", part1, part2, 0)
}
