package quest5

import (
	"fmt"
	"loader"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Sword struct {
	id       int
	fishbone *Segment
	quality  int
}

type Segment struct {
	value int
	left  int
	right int
	child *Segment
}

func parse_line(line string) *Sword {
	parts := strings.Split(line, ":")
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	num_strs := strings.Split(parts[1], ",")
	nums := make([]int, len(num_strs))
	for i, str := range num_strs {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}

	sword := &Sword{id: id}
	sword.initialise(nums)
	return sword
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

func (s *Sword) initialise(nums []int) {
	fishbone := &Segment{value: nums[0], left: -1, right: -1}

	for i := 1; i < len(nums); i++ {
		add_value(fishbone, nums[i])
	}

	quality := 0

	segment := fishbone
	for segment != nil {
		for i := 0; i < digits(segment.value); i++ {
			quality *= 10
		}
		quality += segment.value
		segment = segment.child
	}

	s.fishbone = fishbone
	s.quality = quality
}

func compare_quality(data []string) int {
	min := math.MaxInt64
	max := 0
	for _, line := range data {
		sword := parse_line(line)
		if sword.quality < min {
			min = sword.quality
		} else if sword.quality > max {
			max = sword.quality
		}
	}
	return max - min
}

func combine_numbers(a int, b int, c int) int {
	comb := 0
	if a != -1 {
		comb = a
	}
	for i := 0; i < digits(b); i++ {
		comb *= 10
	}
	comb += b
	if c != -1 {
		for i := 0; i < digits(c); i++ {
			comb *= 10
		}
		comb += c
	}
	return comb
}

func compare_swords(a *Sword, b *Sword) bool {
	if a.quality == b.quality {
		a_seg := a.fishbone
		b_seg := b.fishbone

		for a_seg != nil {
			a_value := combine_numbers(a_seg.left, a_seg.value, a_seg.right)
			b_value := combine_numbers(b_seg.left, b_seg.value, b_seg.right)

			if a_value != b_value {
				return a_value > b_value
			}

			a_seg = a_seg.child
			b_seg = b_seg.child
		}

		return a.id > b.id
	} else {
		return a.quality > b.quality
	}
}

func order(data []string) int {
	swords := make([]*Sword, len(data))
	for i, line := range data {
		swords[i] = parse_line(line)
	}

	sort.Slice(swords, func(i, j int) bool {
		return compare_swords(swords[i], swords[j])
	})

	checksum := 0
	for i, s := range swords {
		checksum += (i + 1) * s.id
	}

	return checksum
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 5, 1

	data := loader.GetOneLine()
	sword := parse_line(string(data))
	part1 := sword.quality

	loader.Part = 2
	multi_data := loader.GetStrings()
	part2 := compare_quality(multi_data)

	loader.Part = 3
	multi_data = loader.GetStrings()
	part3 := order(multi_data)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
