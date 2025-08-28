package quest3

import (
	"fmt"
	"loader"
	"regexp"
	"strconv"
	"strings"
)

type Die struct {
	faces   []int
	seed    int
	pulse   int
	roll_no int
	current int
}

func parse_data(data []string) []Die {
	r := regexp.MustCompile(`faces=\[(.*?)\] seed=(\d+)`)

	dice := make([]Die, len(data))

	for i, line := range data {
		die := Die{}

		matches := r.FindStringSubmatch(line)

		for _, str := range strings.Split(matches[1], ",") {
			n, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			die.faces = append(die.faces, n)
		}

		seed, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}

		die.seed = seed
		die.pulse = seed
		die.roll_no = 1
		dice[i] = die
	}

	return dice
}

func (d *Die) roll() {
	spin := d.roll_no * d.pulse
	d.pulse += spin
	d.pulse %= d.seed
	d.pulse += 1 + d.roll_no + d.seed
	d.current = (d.current + spin) % len(d.faces)
	d.roll_no++
}

func roll_dice(dice []Die, target int) int {
	score := 0
	rolls := 0
	for score < target {
		for i := range dice {
			dice[i].roll()
			score += dice[i].faces[dice[i].current]
		}
		rolls++
	}
	return rolls
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2", 3, 1

	data := loader.GetStrings()
	dice := parse_data(data)
	part1 := roll_dice(dice, 10000)

	fmt.Printf("%d\n", part1)
}
