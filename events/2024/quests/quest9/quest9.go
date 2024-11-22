package quest9

import (
	"fmt"
	"loader"
)

var dots_part1 []int
var dots_part2 []int
var dots_part3 []int
var available []int

func init() {
	dots_part1 = []int{1, 3, 5, 10}
	dots_part2 = []int{1, 3, 5, 10, 15, 16, 20, 24, 25, 30}
	dots_part3 = []int{1, 3, 5, 10, 15, 16, 20, 24, 25, 30, 37, 38, 49, 50, 74, 75, 100, 101}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(nums []int) int {
	m := 0
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func create_dp_slice(sparks []int) []int {
	dp := make([]int, max(sparks)+1)

	for _, dot := range available {
		dp[dot] = 1
	}

	for i := 1; i < len(dp); i++ {
		if dp[i] == 0 {
			continue
		}
		for _, dot := range available {
			if i+dot >= len(dp) {
				break
			}
			if dp[i+dot] == 0 {
				dp[i+dot] = dp[i] + 1
			} else {
				dp[i+dot] = min(dp[i+dot], dp[i]+1)
			}
		}
	}
	return dp
}

func calc_beatles(sparks []int) int {
	dp := create_dp_slice(sparks)

	count := 0
	for _, spark := range sparks {
		count += dp[spark]
	}
	return count
}

func calc_beatles_split(sparks []int) int {
	dp := create_dp_slice(sparks)

	count := 0
	for _, spark := range sparks {
		low := spark / 2
		high := low
		if spark%2 == 1 {
			high++
		}

		smallest := dp[low] + dp[high]
		for {
			low--
			high++
			if high-low > 100 {
				break
			}
			beatles := dp[low] + dp[high]
			if beatles < smallest {
				smallest = beatles
			}
		}
		count += smallest
	}

	return count
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 9, 1

	sparks := loader.GetInts()
	available = dots_part1
	part1 := calc_beatles(sparks)

	loader.Part = 2
	sparks = loader.GetInts()
	available = dots_part2
	part2 := calc_beatles(sparks)

	loader.Part = 3
	sparks = loader.GetInts()
	available = dots_part3
	part3 := calc_beatles_split(sparks)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
