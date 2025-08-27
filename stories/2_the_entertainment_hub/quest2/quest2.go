package quest2

import (
	"fmt"
	"loader"
)

type Balloon struct {
	colour byte
	prev   *Balloon
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
	bolts := []byte{'R', 'G', 'B'}
	bolt_i := 0

	dummy_head := &Balloon{}
	balloon := dummy_head
	for i := 0; i < repeat; i++ {
		for _, colour := range balloons {
			new_balloon := &Balloon{colour: colour, prev: balloon}
			balloon.next = new_balloon
			balloon = new_balloon
		}
	}
	head := dummy_head.next
	balloon.next = head
	head.prev = balloon

	no_balloons := len(balloons) * repeat
	opposite := head
	for i := 0; i < no_balloons/2; i++ {
		opposite = opposite.next
	}

	shots := 0

	for no_balloons > 0 {
		current_colour := head.colour

		prev := head.prev
		next := head.next
		prev.next = next
		next.prev = prev
		head.next = nil
		head.prev = nil
		head = next

		no_balloons--

		if no_balloons&1 == 1 && current_colour == bolts[bolt_i] {
			prev := opposite.prev
			next := opposite.next
			prev.next = next
			next.prev = prev
			opposite.next = nil
			opposite.prev = nil
			opposite = next
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

	loader.Part = 2
	balloons = loader.GetOneLine()
	part2 := shoot_balloon_circle(balloons, 100)

	fmt.Printf("%d\n%d\n", part1, part2)

}
