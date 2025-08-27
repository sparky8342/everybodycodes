package quest2

import (
	"fmt"
	"loader"
)

func shoot_balloons(balloons []byte) int {
	bolts := []byte{'R', 'G', 'B'}
	bolt_i := 0

	i := 0
	shots := 0

	for {
		for i < len(balloons) && balloons[i] == bolts[bolt_i] {
			i++
		}
		i++
		shots++
		bolt_i = (bolt_i + 1) % len(bolts)

		if i >= len(balloons) {
			break
		}
	}

	return shots
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2", 2, 1

	balloons := loader.GetOneLine()
	part1 := shoot_balloons(balloons)

	fmt.Printf("%d\n", part1)

}
