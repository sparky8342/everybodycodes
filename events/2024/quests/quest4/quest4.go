package quest4

import (
	"fmt"
	"loader"
)

func hammer(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	strikes := 0
	for _, num := range nums {
		strikes += num - min
	}
	return strikes
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 4, 1

	nums := loader.GetInts()
	part1 := hammer(nums)

	loader.Part = 2
	nums = loader.GetInts()
	part2 := hammer(nums)

	part3 := -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
