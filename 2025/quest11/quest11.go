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

func checksum(ducks []int) int {
	c := 0
	for i := 0; i < len(ducks); i++ {
		c += (i + 1) * ducks[i]
	}

	return c
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 11, 1

	ducks := loader.GetInts()
	ducks = move_ducks(ducks, 10)
	part1 := checksum(ducks)

	fmt.Printf("%d %s %s\n", part1, "", "")
}
