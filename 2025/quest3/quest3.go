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

func smallest_20(crates []int) int {
	sort.Slice(crates, func(i, j int) bool {
		return crates[i] < crates[j]
	})

	last := crates[0]
	total := crates[0]
	added := 1

	crate := 1
	for added < 20 {
		if crates[crate] > last {
			last = crates[crate]
			total += crates[crate]
			added++
		}
		crate++
	}

	return total
}

func smallest_no_sets(crates []int) int {
	sort.Slice(crates, func(i, j int) bool {
		return crates[i] > crates[j]
	})

	sets := [][]int{[]int{crates[0]}}

outer:
	for i := 1; i < len(crates); i++ {
		for j := range sets {
			if crates[i] < sets[j][len(sets[j])-1] {
				sets[j] = append(sets[j], crates[i])
				continue outer
			}
		}
		sets = append(sets, []int{crates[i]})
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
