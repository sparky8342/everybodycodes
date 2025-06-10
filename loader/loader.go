package loader

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Event string
var Quest int
var Part int

func get_filename() string {
	return fmt.Sprintf("inputs/%s/everybody_codes_e%s_q%02d_p%d.txt", Event, Event, Quest, Part)
}

func GetOneLine() []byte {
	data, err := os.ReadFile(get_filename())
	if err != nil {
		panic(err)
	}
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return data
}

func GetStrings() []string {
	data, err := os.ReadFile(get_filename())
	if err != nil {
		panic(err)
	}
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func GetOneInt() int {
	data, err := os.ReadFile(get_filename())
	if err != nil {
		panic(err)
	}
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	n, err := strconv.Atoi(string(data))
	if err != nil {
		panic(err)
	}
	return n
}

func GetInts() []int {
	data, err := os.ReadFile(get_filename())
	if err != nil {
		panic(err)
	}
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	lines := strings.Split(string(data), "\n")
	nums := make([]int, len(lines))
	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}
