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

type Warriors struct {
	warrior_map   map[int]*Warrior
	warrior_slice []*Warrior
}

func newWarriors() *Warriors {
	w := Warriors{}
	w.warrior_map = map[int]*Warrior{}
	return &w
}

func (w *Warriors) add(id int) {
	if _, ok := w.warrior_map[id]; !ok {
		warrior := &Warrior{id: id}
		w.warrior_map[id] = warrior
		w.warrior_slice = append(w.warrior_slice, warrior)
	}
}

func (w *Warriors) add_part_time(id int, time int) {
	w.warrior_map[id].time += time
	w.warrior_map[id].done_part = true
	w.warrior_map[id].parts++
}

func (w *Warriors) checksum() int {
	sort.Slice(w.warrior_slice, func(i, j int) bool {
		if w.warrior_slice[i].parts == w.warrior_slice[j].parts {
			return w.warrior_slice[i].time < w.warrior_slice[j].time
		} else {
			return w.warrior_slice[i].parts > w.warrior_slice[j].parts
		}
	})
	checksum := 0
	for i := 0; i < len(w.warrior_slice); i++ {
		checksum += (i + 1) * w.warrior_slice[i].time
	}
	return checksum
}

func (w *Warriors) reset_time() {
	for _, warrior := range w.warrior_slice {
		warrior.time = 0
	}
}

func (w *Warriors) reset_done_part() {
	for _, warrior := range w.warrior_slice {
		warrior.done_part = false
	}
}

func (w *Warriors) assign_points(points int) {
	sort.Slice(w.warrior_slice, func(i, j int) bool {
		return w.warrior_slice[i].time < w.warrior_slice[j].time
	})

	last_time := -1
	last_points := -1
	for i := 0; i < len(w.warrior_slice); i++ {
		if w.warrior_slice[i].done_part {
			if w.warrior_slice[i].time == last_time {
				w.warrior_slice[i].points += last_points
			} else {
				if points < 1 {
					break
				}
				w.warrior_slice[i].points += points
				last_time = w.warrior_slice[i].time
				last_points = points
			}
			points--
		}
	}
}

func (w *Warriors) points_checksum() int {
	sort.Slice(w.warrior_slice, func(i, j int) bool {
		return w.warrior_slice[i].points > w.warrior_slice[j].points
	})

	checksum := 0
	place := 1
	last_place := -1
	last_points := -1
	for i := 0; i < len(w.warrior_slice); i++ {
		if w.warrior_slice[i].points == last_points {
			checksum += last_place * w.warrior_slice[i].points
		} else {
			checksum += place * w.warrior_slice[i].points
			last_place = place
			last_points = w.warrior_slice[i].points
		}
		place++
	}
	return checksum
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
	warriors := newWarriors()

	for _, entry := range entries {
		warriors.add(entry.id)
		warriors.add_part_time(entry.id, entry.time)
	}

	return warriors.checksum()
}

func order_warrior_points(entries []Entry) int {
	warriors := newWarriors()

	for _, entry := range entries {
		warriors.add(entry.id)
	}

	max_quest := 1
	for _, entry := range entries {
		if entry.quest > max_quest {
			max_quest = entry.quest
		}
	}

	for quest := 1; quest <= max_quest; quest++ {
		warriors.reset_time()

		for part := 1; part <= 3; part++ {
			warriors.reset_done_part()

			for _, entry := range entries {
				if entry.quest == quest && entry.part == part {
					warriors.add_part_time(entry.id, entry.time)
				}
			}

			warriors.assign_points(part * 50)
		}
	}

	return warriors.points_checksum()
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
