package quest4

import (
	"fmt"
	"loader"
	"sort"
	"utils"
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

func hammer_up_and_down(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	var median int
	if l%2 == 0 {
		median = (nums[l/2-1] + nums[l/2]) / 2
	} else {
		median = nums[l/2]
	}
	strikes := 0
	for _, num := range nums {
		strikes += utils.Abs(num - median)
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

	loader.Part = 3
	nums = loader.GetInts()
	part3 := hammer_up_and_down(nums)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
