package quest3

import (
	"fmt"
	"loader"
	"sort"
	"strconv"
	"strings"
)

func parse_data(data []byte) []int {
	nums := []int{}
	for _, str := range strings.Split(string(data), ",") {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	return nums
}

func largest_set(crates []int) int {
	sort.Slice(crates, func(i, j int) bool {
		return crates[i] > crates[j]
	})

	last := crates[0]
	total := crates[0]

	for i := 1; i < len(crates); i++ {
		if crates[i] < last {
			last = crates[i]
			total += crates[i]
		}
	}

	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 3, 1

	data := loader.GetOneLine()
	crates := parse_data(data)
	part1 := largest_set(crates)

	fmt.Printf("%d %d %d\n", part1, 0, 0)
}
