package quest4

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

const PART1_TURNS = 2025
const PART2_TURNS = 10000000000000.0

func turns(gears []int) int {
	return gears[0] * PART1_TURNS / gears[len(gears)-1]
}

func turns_needed(gears []int) int {
	return int((PART2_TURNS * float64(gears[len(gears)-1]) / float64(gears[0])) + 0.9)
}

func parse_data(data []string) [][]float64 {
	gears := [][]float64{}

	n, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		panic(err)
	}
	gears = append(gears, []float64{0, n})

	for i := 1; i < len(data)-1; i++ {
		parts := strings.Split(data[i], "|")
		n1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			panic(err)
		}
		n2, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			panic(err)
		}
		gears = append(gears, []float64{n1, n2})
	}

	n, err = strconv.ParseFloat(data[len(data)-1], 64)
	if err != nil {
		panic(err)
	}
	gears = append(gears, []float64{n})

	return gears
}

func turns_linked(gears [][]float64) int {
	turns := 100.0

	for i := 0; i < len(gears)-1; i++ {
		turns = turns * gears[i][1] / gears[i+1][0]
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
