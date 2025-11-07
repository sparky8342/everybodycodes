package quest4

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

func turns(gears []int) int {
	return gears[0] * 2025 / gears[len(gears)-1]
}

func turns_needed(gears []int) int {
	return int((10000000000000.0 * float64(gears[len(gears)-1]) / float64(gears[0])) + 0.9)
}

func parse_data(data []string) [][]int {
	gears := [][]int{}

	n, err := strconv.Atoi(data[0])
	if err != nil {
		panic(err)
	}
	gears = append(gears, []int{0, n})

	for i := 1; i < len(data)-1; i++ {
		parts := strings.Split(data[i], "|")
		n1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		gears = append(gears, []int{n1, n2})
	}

	n, err = strconv.Atoi(data[len(data)-1])
	if err != nil {
		panic(err)
	}
	gears = append(gears, []int{n})

	return gears
}

func turns_linked(gears [][]int) int {
	turns := 100.0

	for i := 0; i < len(gears)-1; i++ {
		turns = turns * float64(gears[i][1]) / float64(gears[i+1][0])
	}

	return int(turns)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 4, 1

	gears := loader.GetInts()
	part1 := turns(gears)

	loader.Part = 2
	gears = loader.GetInts()
	part2 := turns_needed(gears)

	loader.Part = 3
	data := loader.GetStrings()
	gears_linked := parse_data(data)
	part3 := turns_linked(gears_linked)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
