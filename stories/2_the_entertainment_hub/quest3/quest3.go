package quest3

import (
	"fmt"
	"loader"
	"regexp"
	"strconv"
	"strings"
)

type Die struct {
	faces     []int
	seed      int
	pulse     int
	roll_no   int
	current   int
	track_pos int
}

func parse_data(data []string) ([]Die, []int) {
	r := regexp.MustCompile(`faces=\[(.*?)\] seed=(\d+)`)

	track := []int{}
	if data[len(data)-2] == "" {
		track_str := data[len(data)-1]
		data = data[:len(data)-2]
		for _, ru := range track_str {
			track = append(track, int(ru-'0'))
		}
	}

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

	return dice, track
}

func (d *Die) roll() int {
	spin := d.roll_no * d.pulse
	d.pulse += spin
	d.pulse %= d.seed
	d.pulse += 1 + d.roll_no + d.seed
	d.current = (d.current + spin) % len(d.faces)
	d.roll_no++
	return d.faces[d.current]
}

func roll_dice(dice []Die, target int) int {
	score := 0
	rolls := 0
	for score < target {
		for i := range dice {
			score += dice[i].roll()
		}
		rolls++
	}
	return rolls
}

func race_track(dice []Die, track []int) string {
	order := []string{}
	for len(order) < len(dice) {
		for i := range dice {
			if dice[i].track_pos < len(track) {
				roll := dice[i].roll()
				if roll == track[dice[i].track_pos] {
					dice[i].track_pos++
					if dice[i].track_pos == len(track) {
						order = append(order, strconv.Itoa(i+1))
					}
				}
			}
		}
	}
	return strings.Join(order, ",")
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2", 3, 1

	data := loader.GetStrings()
	dice, _ := parse_data(data)
	part1 := roll_dice(dice, 10000)

	loader.Part = 2
	data = loader.GetStrings()
	dice, track := parse_data(data)
	part2 := race_track(dice, track)

	fmt.Printf("%d\n%s\n", part1, part2)
}
