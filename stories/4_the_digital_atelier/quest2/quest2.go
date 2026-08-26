package quest2

import (
	"fmt"
	"loader"
	"math/rand/v2"
	"regexp"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

var dirs [][]int

func init() {
	dirs = [][]int{
		[]int{0, 1},
		[]int{0, -1},
		[]int{1, 0},
		[]int{-1, 0},
	}
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

	moves := ""
	if len(data) == 5 {
		moves = strings.Split(data[4], "=")[1]
	}

	return locations, moves
}

func launch(locations map[byte]Pos, moves string) map[Pos]struct{} {
	visited := map[Pos]struct{}{}

	location := locations['S']
	visited[location] = struct{}{}

	for i := 0; i < len(moves); i++ {
		move_to := locations[moves[i]]
		location.x = (location.x + move_to.x) / 2
		location.y = (location.y + move_to.y) / 2
		visited[location] = struct{}{}
	}

	return visited
}

func launch_random_fill(locations map[byte]Pos) map[Pos]struct{} {
	visited := map[Pos]struct{}{}

	location := locations['S']
	visited[location] = struct{}{}

	no_add := 0
	for no_add < 10000000000 {
		var move_to Pos
		r := rand.IntN(3)
		switch r {
		case 0:
			move_to = locations['A']
		case 1:
			move_to = locations['B']
		case 2:
			move_to = locations['C']
		}

		location.x = (location.x + move_to.x) / 2
		location.y = (location.y + move_to.y) / 2
		if _, ok := visited[location]; !ok {
			visited[location] = struct{}{}
			no_add = 0
		} else {
			no_add++
		}
	}

	return visited
}

func add_fireflies(visited map[Pos]struct{}) int {
	fireflies := map[Pos]struct{}{}
	for pos := range visited {
		for _, dir := range dirs {
			firefly := Pos{x: pos.x + dir[0], y: pos.y + dir[1]}
			if _, ok := visited[firefly]; !ok {
				fireflies[firefly] = struct{}{}
			}
		}
	}
	return len(fireflies)
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "4", 2, 1

	data := loader.GetStrings()
	locations, moves := parse_data(data)
	part1 := len(launch(locations, moves))

	loader.Part = 2
	data = loader.GetStrings()
	locations, moves = parse_data(data)
	visited := launch(locations, moves)
	part2 := add_fireflies(visited)

	loader.Part = 3
	data = loader.GetStrings()
	locations, _ = parse_data(data)
	visited = launch_random_fill(locations)
	part3 := add_fireflies(visited)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
