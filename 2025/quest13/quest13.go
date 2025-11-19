package quest13

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Number struct {
	value int
	prev  *Number
	next  *Number
}

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

func parse_data(data []string) [][]int {
	ranges := [][]int{}
	for _, row := range data {
		parts := strings.Split(row, "-")
		n1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, []int{n1, n2})
	}
	return ranges
}

func craft_lock_ranges(ranges [][]int, turns int) int {
	one := &Number{value: 1}
	left := one
	right := one

	total := 1
	for i := 0; i < len(ranges); i += 2 {
		for j := ranges[i][0]; j <= ranges[i][1]; j++ {
			new_number := &Number{value: j}
			right.next = new_number
			new_number.prev = right
			right = right.next
		}
		total += ranges[i][1] - ranges[i][0] + 1
	}

	for i := 1; i < len(ranges); i += 2 {
		for j := ranges[i][0]; j <= ranges[i][1]; j++ {
			new_number := &Number{value: j}
			left.prev = new_number
			new_number.next = left
			left = left.prev
		}
		total += ranges[i][1] - ranges[i][0] + 1
	}

	right.next = left
	left.prev = right

	turns %= total

	next := one
	for i := 0; i < turns; i++ {
		next = next.next
	}

	return next.value
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 13, 1

	nums := loader.GetInts()
	part1 := craft_lock(nums, 2025)

	loader.Part = 2
	data := loader.GetStrings()
	ranges := parse_data(data)
	part2 := craft_lock_ranges(ranges, 20252025)

	loader.Part = 3
	data = loader.GetStrings()
	ranges = parse_data(data)
	part3 := craft_lock_ranges(ranges, 202520252025)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
