package quest5

import (
	"fmt"
	"loader"
	"math"
	"strconv"
	"strings"
)

type Segment struct {
	value int
	left  int
	right int
	child *Segment
}

func parse_line(line string) []int {
	parts := strings.Split(line, ":")
	num_strs := strings.Split(parts[1], ",")
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

func add_value(segment *Segment, value int) {
	if value < segment.value && segment.left == -1 {
		segment.left = value
	} else if value > segment.value && segment.right == -1 {
		segment.right = value
	} else if segment.child == nil {
		segment.child = &Segment{value: value, left: -1, right: -1}
	} else {
		add_value(segment.child, value)
	}
}

func digits(n int) int {
	d := 0
	for n > 0 {
		n /= 10
		d++
	}
	return d
}

func quality(nums []int) int {
	fishbone := &Segment{value: nums[0], left: -1, right: -1}

	for i := 1; i < len(nums); i++ {
		add_value(fishbone, nums[i])
	}

	qual := 0

	segment := fishbone
	for segment != nil {
		for i := 0; i < digits(segment.value); i++ {
			qual *= 10
		}
		qual += segment.value
		segment = segment.child
	}

	return qual
}

func compare_quality(data []string) int {
	min := math.MaxInt64
	max := 0
	for _, line := range data {
		nums := parse_line(line)
		qual := quality(nums)
		if qual < min {
			min = qual
		} else if qual > max {
			max = qual
		}
	}
	return max - min
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 5, 1

	data := loader.GetOneLine()
	nums := parse_line(string(data))
	part1 := quality(nums)

	loader.Part = 2
	data_part2 := loader.GetStrings()
	part2 := compare_quality(data_part2)

	fmt.Printf("%d %d %d\n", part1, part2, 0)
}
