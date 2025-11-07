package quest5

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Segment struct {
	value int
	left  int
	right int
	child *Segment
}

func parse_data(data []byte) []int {
	parts := strings.Split(string(data), ":")
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

func quality(nums []int) string {
	fishbone := &Segment{value: nums[0], left: -1, right: -1}

	for i := 1; i < len(nums); i++ {
		add_value(fishbone, nums[i])
	}

	qual := []string{}

	segment := fishbone
	for segment != nil {
		qual = append(qual, strconv.Itoa(segment.value))
		segment = segment.child
	}

	return strings.Join(qual, "")
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 5, 1

	data := loader.GetOneLine()
	nums := parse_data(data)
	part1 := quality(nums)

	fmt.Printf("%s %d %d\n", part1, 0, 0)
}
