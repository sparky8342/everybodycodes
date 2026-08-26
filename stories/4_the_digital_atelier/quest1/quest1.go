package quest1

import (
	"fmt"
	"loader"
	"math/big"
)

func process_sequence(sequence []int, unique_points bool) int {
	number_line := big.NewInt(1)
	n := 0
	for _, num := range sequence {
		if n-num < 0 || number_line.Bit(n-num) == 1 {
			n += num
		} else {
			n -= num
		}

		if unique_points {
			for number_line.Bit(n) == 1 {
				n++
			}
		}

		number_line.SetBit(number_line, n, 1)
	}
	return n
}

func process_sequences(sequences [][]int, unique_points bool) int {
	total := 0
	for _, sequence := range sequences {
		total += process_sequence(sequence, unique_points)
	}
	return total
}

func crosses_range(ranges [][]int, from int, to int) bool {
	for _, r := range ranges {
		if from < r[0] && to > r[0] && to < r[1] {
			return true
		}
		if from > r[0] && from < r[1] && to > r[1] {
			return true
		}
	}
	return false
}

func process_sequence_no_crossover(sequence []int) int {
	number_line := big.NewInt(1)
	n := 0

	highest := 0

	ranges := make([][][]int, 2)

	range_id := 0

outer:
	for _, num := range sequence {
		// go back
		if n-num > 0 && number_line.Bit(n-num) == 0 && !crosses_range(ranges[range_id], n-num, n) {
			number_line.SetBit(number_line, n-num, 1)
			ranges[range_id] = append(ranges[range_id], []int{n - num, n})
			range_id = (range_id + 1) % 2
			n -= num
			continue
		}

		// go forward
		to := n + num
		for number_line.Bit(to) == 1 || crosses_range(ranges[range_id], n, to) {
			to++
			if to > highest+1 {
				continue outer
			}
		}

		number_line.SetBit(number_line, to, 1)
		ranges[range_id] = append(ranges[range_id], []int{n, to})
		range_id = (range_id + 1) % 2

		n = to
		if to > highest {
			highest = to
		}

	}

	return n
}

func process_sequences_no_crossover(sequences [][]int) int {
	total := 0
	for _, sequence := range sequences {
		total += process_sequence_no_crossover(sequence)
	}
	return total
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "4", 1, 1

	sequences := loader.GetCSVInts()
	part1 := process_sequences(sequences, false)

	loader.Part = 2
	sequences = loader.GetCSVInts()
	part2 := process_sequences(sequences, true)

	loader.Part = 3
	sequences = loader.GetCSVInts()
	part3 := process_sequences_no_crossover(sequences)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
