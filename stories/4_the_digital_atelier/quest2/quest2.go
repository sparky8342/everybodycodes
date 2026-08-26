package quest2

import (
	"fmt"
	"loader"
	"regexp"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

func parse_data(data []string) (map[byte]Pos, string) {
	r, err := regexp.Compile("(\\w).*?(\\d+),(\\d+)")
	if err != nil {
		panic(err)
	}

	locations := map[byte]Pos{}
	for i := 0; i < 4; i++ {
		m := r.FindStringSubmatch(data[i])
		x, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(m[3])
		if err != nil {
			panic(err)
		}

		locations[m[1][0]] = Pos{x: x, y: y}
	}

	moves := strings.Split(data[4], "=")[1]

	return locations, moves
}

func launch(locations map[byte]Pos, moves string) int {
	visited := map[Pos]struct{}{}

	location := locations['S']
	visited[location] = struct{}{}

	for i := 0; i < len(moves); i++ {
		move_to := locations[moves[i]]
		location.x = (location.x + move_to.x) / 2
		location.y = (location.y + move_to.y) / 2
		visited[location] = struct{}{}
	}

	return len(visited)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "4", 2, 1

	data := loader.GetStrings()
	locations, moves := parse_data(data)
	part1 := launch(locations, moves)

	part2, part3 := -1, -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
