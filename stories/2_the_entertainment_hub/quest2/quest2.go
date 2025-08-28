package quest2

import (
	"fmt"
	"loader"
)

const bolts = "RGB"

type Balloon struct {
	colour byte
	next   *Balloon
}

func print_circle(head *Balloon) {
	balloon := head
	for balloon.next != head {
		fmt.Print(string(balloon.colour) + " ")
		balloon = balloon.next
	}
	fmt.Println(string(balloon.colour))
}

func shoot_balloon_circle(balloons []byte, repeat int) int {
	dummy_head := &Balloon{}
	balloon := dummy_head
	for i := 0; i < repeat; i++ {
		for _, colour := range balloons {
			balloon.next = &Balloon{colour: colour}
			balloon = balloon.next
		}
	}

	head := balloon
	head.next = dummy_head.next

	no_balloons := len(balloons) * repeat
	opposite := head
	for i := 0; i < no_balloons/2; i++ {
		opposite = opposite.next
	}

	bolt_i := 0
	shots := 0

	for no_balloons > 0 {
		current_colour := head.next.colour
		head.next = head.next.next
		no_balloons--

		if no_balloons&1 == 1 && current_colour == bolts[bolt_i] {
			opposite.next = opposite.next.next
			no_balloons--
		} else if no_balloons&1 == 0 {
			opposite = opposite.next
		}

		shots++
		bolt_i = (bolt_i + 1) % len(bolts)
	}

	return shots
}

func shoot_balloons(balloons []byte) int {
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

	loader.Part = 2
	balloons = loader.GetOneLine()
	part2 := shoot_balloon_circle(balloons, 100)

	loader.Part = 3
	balloons = loader.GetOneLine()
	part3 := shoot_balloon_circle(balloons, 100000)

	fmt.Printf("%d\n%d\n%d\n", part1, part2, part3)
}
