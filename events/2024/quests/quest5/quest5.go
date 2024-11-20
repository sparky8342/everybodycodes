package quest5

import (
	"fmt"
	"loader"
	"os"
	"strconv"
	"strings"
)

func parse_data(data []string) [][]int {
	cols := len(strings.Replace(data[0], " ", "", -1))
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
	nums[col_pos] = append(nums[col_pos][:clapper_pos+1], nums[col_pos][clapper_pos:]...)
	nums[col_pos][clapper_pos] = clapper
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

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 5, 1

	data := loader.GetStrings()
	nums := parse_data(data)
	fmt.Println(nums)
	steps(nums, 10)

	part1 := -1

	part2, part3 := -1, -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
