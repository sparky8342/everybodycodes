package quest13

import (
	"fmt"
	"loader"
)

func craft_lock(nums []int, turns int) int {
	lock := []int{1}

	for i := 0; i < len(nums); i += 2 {
		lock = append(lock, nums[i])
	}
	for i := 1; i < len(nums); i += 2 {
		lock = append([]int{nums[i]}, lock...)
	}

	return lock[((len(nums)/2)+turns)%len(lock)]
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 13, 1

	nums := loader.GetInts()
	part1 := craft_lock(nums, 2025)

	fmt.Printf("%d %d %d\n", part1, 0, 0)
}
