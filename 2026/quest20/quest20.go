package quest20

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Equation struct {
	n1       int
	op1      byte
	n2       int
	n3       int
	op2      byte
	n4       int
	segments int
	valid    bool
}

var no_segments map[byte]int
var remove map[byte][]byte
var change map[byte][]byte
var add map[byte][]byte

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

	remove = map[byte][]byte{
		'6': []byte{'5'},
		'7': []byte{'1'},
		'8': []byte{'9'},
		'9': []byte{'3', '5'},
		'+': []byte{'-'},
	}

	change = map[byte][]byte{
		'0': []byte{'6', '9', 'A'},
		'2': []byte{'3'},
		'3': []byte{'2', '5'},
		'5': []byte{'3'},
		'6': []byte{'0', '9', 'A'},
		'9': []byte{'0', '6', 'A'},
		'A': []byte{'0', '6', '9'},
	}

	add = map[byte][]byte{
		'0': []byte{'8'},
		'1': []byte{'7'},
		'5': []byte{'6', '9'},
		'6': []byte{'8'},
		'9': []byte{'8'},
	}
}

func parse_line(line string, base int) Equation {
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

	valid := false
	if left == right {
		valid = true
	}

	return Equation{
		n1:       nums[0],
		op1:      parts[1][0],
		n2:       nums[1],
		n3:       nums[2],
		op2:      parts[5][0],
		n4:       nums[3],
		segments: segments,
		valid:    valid,
	}
}

func count_segments(data []string, base int) string {
	count := 0
	for _, line := range data {
		equation := parse_line(line, base)
		if equation.valid {
			count += equation.segments
		}
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
