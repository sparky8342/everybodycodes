package quest5

import (
	"fmt"
	"loader"
	"os"
	"strconv"
	"strings"
)

func parse_data(data []string) [][]int {
	cols := len(strings.Split(data[0], " "))
	nums := make([][]int, cols)

	for _, line := range data {
		parts := strings.Split(line, " ")
		for i, part := range parts {
			n, err := strconv.Atoi(part)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error %v\n", err)
				os.Exit(1)
			}
			nums[i] = append(nums[i], n)
		}
	}

	return nums
}

func step(nums [][]int, col_pos int) {
	clapper := nums[col_pos][0]
	nums[col_pos] = nums[col_pos][1:]

	clapper_pos := -1
	moving_down := true
	col_pos++
	if col_pos == len(nums) {
		col_pos = 0
	}

	for i := 0; i < clapper; i++ {
		if moving_down {
			clapper_pos++
			if clapper_pos == len(nums[col_pos]) {
				clapper_pos--
				moving_down = false
			}
		} else {
			clapper_pos--
			if clapper_pos == -1 {
				clapper_pos = 0
				moving_down = true
			}
		}
	}

	if !moving_down {
		clapper_pos++
	}
	if clapper_pos == len(nums[col_pos]) {
		nums[col_pos] = append(nums[col_pos], clapper)
	} else {
		nums[col_pos] = append(nums[col_pos][:clapper_pos+1], nums[col_pos][clapper_pos:]...)
		nums[col_pos][clapper_pos] = clapper
	}
}

func steps(nums [][]int, amount int) int {
	col_pos := 0
	for i := 0; i < amount; i++ {
		step(nums, col_pos)
		col_pos++
		if col_pos == len(nums) {
			col_pos = 0
		}
	}
	n := 0
	for i := 0; i < len(nums); i++ {
		n = n*10 + nums[i][0]
	}
	return n
}

func no_digits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func find_repeat(nums [][]int, repeat int) int {
	digits := no_digits(nums[0][0])

	seen := map[int]int{}

	col_pos := 0
	round := 0
	for {
		step(nums, col_pos)
		col_pos++
		if col_pos == len(nums) {
			col_pos = 0
		}
		round++

		n := 0
		for i := 0; i < len(nums); i++ {
			for j := 0; j < digits; j++ {
				n *= 10
			}
			n += nums[i][0]
		}

		seen[n]++
		if seen[n] == repeat {
			return round * n
		}
	}
}

func highest_top_number(nums [][]int) int {
	digits := no_digits(nums[0][0])

	seen := map[[2]int]struct{}{}

	col_pos := 0
	max := 0

	for {
		step(nums, col_pos)
		col_pos++
		if col_pos == len(nums) {
			col_pos = 0
		}

		n := 0
		for i := 0; i < len(nums); i++ {
			for j := 0; j < digits; j++ {
				n *= 10
			}
			n += nums[i][0]
		}

		if n > max {
			max = n
		}

		if _, ok := seen[[2]int{n, col_pos}]; ok {
			return max
		}
		seen[[2]int{n, col_pos}] = struct{}{}
	}
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 5, 1

	data := loader.GetStrings()
	nums := parse_data(data)
	part1 := steps(nums, 10)

	loader.Part = 2
	data = loader.GetStrings()
	nums = parse_data(data)
	part2 := find_repeat(nums, 2024)

	loader.Part = 3
	data = loader.GetStrings()
	nums = parse_data(data)
	part3 := highest_top_number(nums)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
