package quest19

import (
	"fmt"
	"loader"
	"math"
	"strconv"
	"strings"
)

type pair [2]int

type Wall struct {
	x    int
	gaps []pair
}

type Pos struct {
	x int
	y int
}

type Entry struct {
	pos   Pos
	flaps int
}

func parse_data(data []string) []Wall {
	walls := []Wall{}

	for _, line := range data {
		parts := strings.Split(line, ",")
		nums := [3]int{}
		for i := 0; i < 3; i++ {
			n, err := strconv.Atoi(parts[i])
			if err != nil {
				panic(err)
			}
			nums[i] = n
		}

		x := nums[0]
		gap := pair{nums[1], nums[1] + nums[2] - 1}

		if len(walls) > 0 {
			last := len(walls) - 1
			if walls[last].x == x {
				walls[last].gaps = append(walls[last].gaps, gap)
				continue
			}
		}

		walls = append(walls, Wall{
			x:    nums[0],
			gaps: []pair{gap},
		})
	}

	return walls
}

func flap(walls []Wall) int {
	x, y := 0, 0

	flaps := 0
	for _, wall := range walls {
		for ; x <= wall.x; x++ {
			if y < wall.gaps[0][0] {
				flaps++
				y++
			} else {
				y--
			}
		}
	}

	return flaps
}

func multi_flap(walls []Wall) int {
	x_walls := map[int][]pair{}

	// TODO - do this format in parse_data instead (and change part 1)
	for _, wall := range walls {
		if _, ok := x_walls[wall.x]; !ok {
			x_walls[wall.x] = []pair{}
		}
		x_walls[wall.x] = append(x_walls[wall.x], wall.gaps...)
	}

	end_x := walls[len(walls)-1].x + 1

	min_flaps := math.MaxInt32
	search(Pos{}, end_x, 0, x_walls, &min_flaps)

	return min_flaps
}

func search(pos Pos, end_x int, flaps int, x_walls map[int][]pair, min_flaps *int) {
	if flaps > *min_flaps {
		return
	}

	x := pos.x + 1

	if x == end_x {
		if flaps < *min_flaps {
			*min_flaps = flaps
		}
		return
	}

	if gaps, ok := x_walls[x]; ok {
		for _, y := range []int{pos.y - 1, pos.y + 1} {
			for _, gap := range gaps {
				if y >= gap[0] && y <= gap[1] {
					new_flaps := flaps
					if y == pos.y+1 {
						new_flaps++
					}
					search(Pos{x: x, y: y}, end_x, new_flaps, x_walls, min_flaps)
					break
				}
			}
		}
	} else {
		search(Pos{x: x, y: pos.y - 1}, end_x, flaps, x_walls, min_flaps)
		search(Pos{x: x, y: pos.y + 1}, end_x, flaps+1, x_walls, min_flaps)
	}
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 19, 1

	data := loader.GetStrings()
	walls := parse_data(data)
	part1 := flap(walls)

	loader.Part = 2
	data = loader.GetStrings()
	walls = parse_data(data)
	part2 := multi_flap(walls)

	fmt.Printf("%d %d %d\n", part1, part2, 0)
}
