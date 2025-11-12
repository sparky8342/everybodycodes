package quest8

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

func parse_data(data []byte) []int {
	num_strs := strings.Split(string(data), ",")
	nums := make([]int, len(num_strs))
	for i, str := range num_strs {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}

func centre_count(nails int, nums []int) int {
	half := nails / 2
	count := 0
	for i := 1; i < len(nums); i++ {
		a, b := nums[i-1], nums[i]
		if a > b {
			a, b = b, a
		}
		if a+half == b {
			count++
		}
	}
	return count
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 8, 1

	data := loader.GetOneLine()
	nums := parse_data(data)
	part1 := centre_count(32, nums)

	fmt.Printf("%d %s %s\n", part1, "", "")
}
