package quest1

import (
	"fmt"
	"loader"
)

func process_sequence(sequence []int) int {
	number_line := []bool{true}
	n := 0
	for _, num := range sequence {
		if n - num < 0 || number_line[n - num] == true {
			n += num
		} else {
			n -= num
		}

		for n >= len(number_line) {
			number_line = append(number_line, false)
		}
		number_line[n] = true
	}
	return n
}

func process_sequences(sequences [][]int) int {
	total := 0
	for _, sequence := range sequences {
		total += process_sequence(sequence)
	}
	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "4", 1, 1

	sequences := loader.GetCSVInts()
	part1 := process_sequences(sequences)

	part2, part3 := -1, -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
