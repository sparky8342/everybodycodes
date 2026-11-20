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
		'A': 6,
		'+': 2,
		'-': 1,
		'=': 2,
	}
}

func equation_segments(line string, base int) int {
	segments := 0

	parts := strings.Split(line, " ")
	nums := [4]int{}
	for i := 0; i <= 6; i += 2 {
		n, err := strconv.ParseInt(parts[i], base, 64)
		if err != nil {
			panic(err)
		}
		nums[i/2] = int(n)
		for j := 0; j < len(parts[i]); j++ {
			segments += no_segments[parts[i][j]]
		}
	}
	fmt.Println(nums)
	segments += no_segments[parts[1][0]]
	segments += no_segments[parts[3][0]]
	segments += no_segments[parts[5][0]]

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

	if left == right {
		fmt.Println(segments)
		return segments
	} else {
		return 0
	}
}

func count_segments(data []string, base int) string {
	count := 0
	for _, line := range data {
		count += equation_segments(line, base)
	}
	return strconv.FormatInt(int64(count), base)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 20, 1

	data := loader.GetStrings()
	part1 := count_segments(data, 10)

	loader.Part = 2
	data = loader.GetStrings()
	part2 := count_segments(data, 11)

	loader.Part = 3
	part3 := -1

	fmt.Printf("%s %s %d\n", part1, part2, part3)
}
