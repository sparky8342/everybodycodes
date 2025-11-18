package quest11

import (
	"fmt"
	"loader"
)

func move_right_one_round(ducks []int) (bool, []int) {
	moved := false
	for i := 0; i < len(ducks)-1; i++ {
		if ducks[i] > ducks[i+1] {
			ducks[i]--
			ducks[i+1]++
			moved = true
		}
	}
	return moved, ducks
}

func move_left_one_round(ducks []int) (bool, []int) {
	moved := false
	for i := 0; i < len(ducks)-1; i++ {
		if ducks[i] < ducks[i+1] {
			ducks[i]++
			ducks[i+1]--
			moved = true
		}
	}
	return moved, ducks
}

func move_ducks(ducks []int, rounds int) []int {
	var moved bool

	var i int
	for i = 0; i < rounds; i++ {
		moved, ducks = move_right_one_round(ducks)
		if !moved {
			break
		}
	}

	for ; i < rounds; i++ {
		moved, ducks = move_left_one_round(ducks)
		if !moved {
			break
		}
	}

	return ducks
}

func move_left(ducks []int) int {
	rounds := 0

	total := 0
	for _, duck := range ducks {
		total += duck
	}
	avg := total / len(ducks)

	for i := len(ducks) - 1; i > 0; i-- {
		if ducks[i] > avg {
			diff := ducks[i] - avg
			ducks[i] -= diff
			ducks[i-1] += diff
			if diff > rounds {
				rounds = diff
			}
		}
	}

	return rounds
}

func checksum(ducks []int) int {
	c := 0
	for i := 0; i < len(ducks); i++ {
		c += (i + 1) * ducks[i]
	}
	return c
}

func balance_ducks(ducks []int) int {
	rounds := 0

	var moved bool
	for {
		moved, ducks = move_right_one_round(ducks)
		if moved {
			rounds++
		} else {
			break
		}
	}

	rounds += move_left(ducks)
	return rounds
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 11, 1

	ducks := loader.GetInts()
	ducks = move_ducks(ducks, 10)
	part1 := checksum(ducks)

	loader.Part = 2
	ducks = loader.GetInts()
	part2 := balance_ducks(ducks)

	loader.Part = 3
	ducks = loader.GetInts()
	part3 := balance_ducks(ducks)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
