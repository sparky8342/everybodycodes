package quest20

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Equation struct {
	eq       []byte
	segments int
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
		' ': 0,
	}

	remove = map[byte][]byte{
		'6': []byte{'5'},
		'7': []byte{'1'},
		'8': []byte{'0', '6', '9', 'A'},
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
		'3': []byte{'9'},
		'5': []byte{'6', '9'},
		'6': []byte{'8'},
		'9': []byte{'8'},
		'A': []byte{'8'},
		'-': []byte{'+'},
	}
}

func parse_line(line string) Equation {
	segments := 0
	for i := 0; i < len(line); i++ {
		segments += no_segments[line[i]]
	}
	return Equation{eq: []byte(line), segments: segments}
}

func (e Equation) valid(base int) bool {
	parts := strings.Split(string(e.eq), " ")
	nums := [4]int{}
	for i := 0; i <= 6; i += 2 {
		n, err := strconv.ParseInt(parts[i], base, 64)
		if err != nil {
			panic(err)
		}
		nums[i/2] = int(n)
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

func count_segments(data []string, base int) string {
	count := 0
	for _, line := range data {
		equation := parse_line(line)
		if equation.valid(base) {
			count += equation.segments
		}
	}
	return strconv.FormatInt(int64(count), base)
}

func (e Equation) balance(base int) bool {
	// change
	for i := 0; i < len(e.eq); i++ {
		d := e.eq[i]

		if val, ok := change[d]; ok {
			for _, new_d := range val {
				e.eq[i] = new_d
				if e.valid(base) {
					return true
				}
			}
		}

		e.eq[i] = d
	}

	// remove and add
	for i := 0; i < len(e.eq); i++ {
		for j := 0; j < len(e.eq); j++ {
			if i == j {
				continue
			}

			d1 := e.eq[i]
			d2 := e.eq[j]

			for _, r := range remove[d1] {
				e.eq[i] = r
				for _, a := range add[d2] {
					e.eq[j] = a
					if e.valid(base) {
						return true
					}
				}
			}

			e.eq[i] = d1
			e.eq[j] = d2
		}
	}

	return false
}

func count_segments_balance(data []string, base int) string {
	count := 0
	for _, line := range data {
		equation := parse_line(line)
		if equation.valid(base) || equation.balance(base) {
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
	data = loader.GetStrings()
	part3 := count_segments_balance(data, 11)

	fmt.Printf("%s %s %s\n", part1, part2, part3)
}
