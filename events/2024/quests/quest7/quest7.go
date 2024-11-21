package quest7

import (
	"fmt"
	"loader"
	"sort"
	"strings"
)

type Chariot struct {
	name      string
	changes   []int
	change_no int
	power     int
	essence   int
}

func parse_data(data []string) []*Chariot {
	chariots := make([]*Chariot, len(data))

	for i, line := range data {
		parts := strings.Split(line, ":")
		name := parts[0]
		change_strs := strings.Split(parts[1], ",")
		changes := make([]int, len(change_strs))
		for j, change := range change_strs {
			switch change {
			case "+":
				changes[j] = 1
			case "-":
				changes[j] = -1
			case "=":
				changes[j] = 0
			}
		}
		chariots[i] = &Chariot{
			name:    name,
			changes: changes,
			power:   10,
		}
	}

	return chariots
}

func (c *Chariot) move(override byte) {
	if override == '+' {
		c.power++
	} else if override == '-' {
		c.power--
	} else {
		c.power += c.changes[c.change_no]
	}
	if c.power < 0 {
		c.power = 0
	}
	c.essence += c.power

	c.change_no++
	if c.change_no == len(c.changes) {
		c.change_no = 0
	}
}

func race(chariots []*Chariot, turns int) string {
	for i := 0; i < turns; i++ {
		for _, chariot := range chariots {
			chariot.move('=')
		}
	}

	sort.Slice(chariots, func(i, j int) bool {
		return chariots[i].essence > chariots[j].essence
	})

	order := make([]string, len(chariots))
	for i, chariot := range chariots {
		order[i] = chariot.name
	}

	return strings.Join(order, "")
}

func race_track(chariots []*Chariot, track string, loops int) string {
	for _, chariot := range chariots {
		for i := 0; i < loops; i++ {
			for j := 0; j < len(track); j++ {
				chariot.move(track[j])
			}
		}
	}

	sort.Slice(chariots, func(i, j int) bool {
		return chariots[i].essence > chariots[j].essence
	})

	order := make([]string, len(chariots))
	for i, chariot := range chariots {
		order[i] = chariot.name
	}

	return strings.Join(order, "")
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 7, 1

	data := loader.GetStrings()
	chariots := parse_data(data)
	part1 := race(chariots, 10)

	loader.Part = 2
	data = loader.GetStrings()
	chariots = parse_data(data)
	track := "-=++=-==++=++=-=+=-=+=+=--=-=++=-==++=-+=-=+=-=+=+=++=-+==++=++=-=-=---=++==--+++==++=+=--==++==+++=++=+++=--=+=-=+=-+=-+=-+-=+=-=+=-+++=+==++++==---=+=+=-S"
	part2 := race_track(chariots, track, 10)

	part3 := ""
	fmt.Printf("%s %s %s\n", part1, part2, part3)
}
