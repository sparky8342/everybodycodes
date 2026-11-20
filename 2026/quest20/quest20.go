package quest20

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

var no_segments map[byte]int

func init() {
	no_segments = map[byte]int{
		'0': 6,
		'1': 2,
		'2': 5,
		'3': 5,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 3,
		'8': 7,
		'9': 6,
		'+': 2,
		'-': 1,
		'=': 2,
	}
}

func equation_valid(line string) bool {
	parts := strings.Split(line, " ")
	nums := [4]int{}
	for i := 0; i <= 6; i += 2 {
		n, err := strconv.Atoi(parts[i])
		if err != nil {
			panic(err)
		}
		nums[i/2] = n
	}

	var left, right int
	if parts[1] == "+" {
		left = nums[0] + nums[1]
	} else if parts[1] == "-" {
		left = nums[0] - nums[1]
	}
	if parts[5] == "+" {
		right = nums[2] + nums[3]
	} else if parts[5] == "-" {
		right = nums[2] - nums[3]
	}

	return left == right
}

func count_segments(data []string) int {
	count := 0
	for _, line := range data {
		if equation_valid(line) {
			for i := 0; i < len(line); i++ {
				if line[i] != ' ' {
					count += no_segments[line[i]]
				}
			}
		}
	}
	return count
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 20, 1

	data := loader.GetStrings()
	part1 := count_segments(data)

	loader.Part = 2
	part2 := -1

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
