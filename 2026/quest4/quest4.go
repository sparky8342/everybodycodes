package quest4

import (
	"fmt"
	"loader"
	"sort"
	"strconv"
	"strings"
)

type Warrior struct {
	id        int
	parts     int
	time      int
	points    int
	done_part bool
}

type Entry struct {
	id    int
	quest int
	part  int
	time  int
}

func parse_data(data []string) []Entry {
	entries := make([]Entry, len(data))
	for i := 0; i < len(data); i++ {
		parts := strings.Split(data[i], " ")
		nums := make([]int, 4)
		for j := 0; j < 4; j++ {
			num, err := strconv.Atoi(strings.Split(parts[j], "=")[1])
			if err != nil {
				panic(err)
			}
			nums[j] = num
		}
		entries[i] = Entry{id: nums[0], quest: nums[1], part: nums[2], time: nums[3]}
	}
	return entries
}

func order_warriors(entries []Entry) int {
	warriors := map[int]*Warrior{}

	for _, entry := range entries {
		if _, ok := warriors[entry.id]; !ok {
			warriors[entry.id] = &Warrior{id: entry.id}
		}

		warriors[entry.id].time += entry.time
		warriors[entry.id].parts++
	}

	warriors_sl := []*Warrior{}
	for _, warrior := range warriors {
		warriors_sl = append(warriors_sl, warrior)
	}

	sort.Slice(warriors_sl, func(i, j int) bool {
		if warriors_sl[i].parts == warriors_sl[j].parts {
			return warriors_sl[i].time < warriors_sl[j].time
		} else {
			return warriors_sl[i].parts > warriors_sl[j].parts
		}
	})

	checksum := 0
	for i := 0; i < len(warriors_sl); i++ {
		checksum += (i + 1) * warriors_sl[i].time
	}
	return checksum
}

func order_warrior_points(entries []Entry) int {
	warriors := map[int]*Warrior{}
	warriors_sl := []*Warrior{}
	for _, entry := range entries {
		if _, ok := warriors[entry.id]; !ok {
			warrior := &Warrior{id: entry.id}
			warriors[entry.id] = warrior
			warriors_sl = append(warriors_sl, warrior)
		}
	}

	max_quest := 1
	for _, entry := range entries {
		if entry.quest > max_quest {
			max_quest = entry.quest
		}
	}

	for quest := 1; quest <= max_quest; quest++ {
		for i := 0; i < len(warriors_sl); i++ {
			warriors_sl[i].time = 0
		}

		for part := 1; part <= 3; part++ {
			for i := 0; i < len(warriors_sl); i++ {
				warriors_sl[i].done_part = false
			}

			for _, entry := range entries {
				if entry.quest == quest && entry.part == part {
					warriors[entry.id].time += entry.time
					warriors[entry.id].done_part = true
				}
			}
			sort.Slice(warriors_sl, func(i, j int) bool {
				return warriors_sl[i].time < warriors_sl[j].time
			})

			var points int
			switch part {
			case 1:
				points = 50
			case 2:
				points = 100
			case 3:
				points = 150
			}

			last_time := -1
			last_points := -1
			for i := 0; i < len(warriors_sl); i++ {
				if warriors_sl[i].done_part {
					if warriors_sl[i].time == last_time {
						warriors_sl[i].points += last_points
					} else {
						if points < 1 {
							break
						}
						warriors_sl[i].points += points
						last_time = warriors_sl[i].time
						last_points = points
					}
					points--
				}
			}
		}
	}

	sort.Slice(warriors_sl, func(i, j int) bool {
		return warriors_sl[i].points > warriors_sl[j].points
	})

	checksum := 0
	place := 1
	last_place := -1
	last_points := -1
	for i := 0; i < len(warriors_sl); i++ {
		if warriors_sl[i].points == last_points {
			checksum += last_place * warriors_sl[i].points
		} else {
			checksum += place * warriors_sl[i].points
			last_place = place
			last_points = warriors_sl[i].points
		}
		place++
	}
	return checksum
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 4, 1

	data := loader.GetStrings()
	entries := parse_data(data)
	part1 := order_warriors(entries)

	loader.Part = 2
	data = loader.GetStrings()
	entries = parse_data(data)
	part2 := order_warriors(entries)

	loader.Part = 3
	data = loader.GetStrings()
	entries = parse_data(data)
	part3 := order_warrior_points(entries)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
