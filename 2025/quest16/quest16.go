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

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 16, 1

	data := loader.GetOneLine()
	nums := parse_data(data)
	part1 := build_wall(nums, 90)

	fmt.Printf("%d %d %d\n", part1, 0, 0)
}
