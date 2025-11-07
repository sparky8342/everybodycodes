package quest3

import (
	"fmt"
	"loader"
	"sort"
	"strconv"
	"strings"
)

func parse_data(data []byte) []int {
	num_strs := strings.Split(string(data), ",")
	nums := make([]int, len(num_strs))
	for i := 0; i < len(num_strs); i++ {
		n, err := strconv.Atoi(num_strs[i])
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}

func largest_set(crates []int) int {
	seen := map[int]struct{}{}
	total := 0
	for _, crate := range crates {
		if _, ok := seen[crate]; !ok {
			total += crate
			seen[crate] = struct{}{}
		}
	}
	return total
}

func smallest_20(crates []int) int {
	unique := []int{}
	seen := map[int]struct{}{}
	for _, crate := range crates {
		if _, ok := seen[crate]; !ok {
			unique = append(unique, crate)
			seen[crate] = struct{}{}
		}
	}

	sort.Slice(unique, func(i, j int) bool {
		return unique[i] < unique[j]
	})

	total := 0
	for _, crate := range unique[:20] {
		total += crate
	}

	return total
}

func smallest_no_sets(crates []int) int {
	sort.Slice(crates, func(i, j int) bool {
		return crates[i] > crates[j]
	})

	sets := []int{crates[0]}

outer:
	for i := 1; i < len(crates); i++ {
		for j := range sets {
			if crates[i] < sets[j] {
				sets[j] = crates[i]
				continue outer
			}
		}
		sets = append(sets, crates[i])
	}

	return len(sets)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 3, 1

	data := loader.GetOneLine()
	crates := parse_data(data)
	part1 := largest_set(crates)

	loader.Part = 2
	data = loader.GetOneLine()
	crates = parse_data(data)
	part2 := smallest_20(crates)

	loader.Part = 3
	data = loader.GetOneLine()
	crates = parse_data(data)
	part3 := smallest_no_sets(crates)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
