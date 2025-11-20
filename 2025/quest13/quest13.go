package quest13

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Range struct {
	start  int
	amount int
	asc    bool
	next   *Range
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

func craft_lock(nums []int, turns int) int {
	ranges := make([][]int, len(nums))
	for i, n := range nums {
		ranges[i] = []int{n, n}
	}
	return craft_lock_ranges(ranges, turns)
}

func craft_lock_ranges(ranges [][]int, turns int) int {
	one := &Range{start: 1, amount: 1, asc: true}
	left := one
	right := one

	total := 1
	for i := 0; i < len(ranges); i += 2 {
		new_range := &Range{
			start:  ranges[i][0],
			amount: ranges[i][1] - ranges[i][0] + 1,
			asc:    true,
		}
		total += new_range.amount
		right.next = new_range
		right = right.next
	}

	for i := 1; i < len(ranges); i += 2 {
		new_range := &Range{
			start:  ranges[i][1],
			amount: ranges[i][1] - ranges[i][0] + 1,
			asc:    false,
		}
		total += new_range.amount
		new_range.next = left
		left = new_range
	}

	right.next = left

	turns %= total

	next := one
	for {
		if turns >= next.amount {
			turns -= next.amount
			next = next.next
		} else if next.asc {
			return next.start + turns
		} else {
			return next.start - turns
		}
	}

	return -1
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
