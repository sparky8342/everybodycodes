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

func multi_gap(walls []Wall) int {
	positions := map[Pos]int{}
	positions[Pos{}] = 0

	for _, wall := range walls {
		next_wall_positions := []Pos{}
		for _, gap := range wall.gaps {
			for y := gap[0]; y <= gap[1]; y++ {
				next_wall_positions = append(next_wall_positions, Pos{x: wall.x, y: y})
			}
		}

		next_positions := map[Pos]int{}
		for pos, flaps := range positions {
			for _, wall_pos := range next_wall_positions {
				if pos.x%2 == wall_pos.x%2 && pos.y%2 != wall_pos.y%2 {
					continue
				}
				if pos.x%2 != wall_pos.x%2 && pos.y%2 == wall_pos.y%2 {
					continue
				}

				x_dist := wall_pos.x - pos.x

				if wall_pos.y > pos.y {
					up_needed := wall_pos.y - pos.y
					if up_needed > x_dist {
						continue
					}
					f := flaps + (x_dist-up_needed)/2 + up_needed
					if val, ok := next_positions[wall_pos]; ok {
						if f > val {
							continue
						}
					}
					next_positions[wall_pos] = f
				} else if wall_pos.y < pos.y {
					down_needed := pos.y - wall_pos.y
					if down_needed > x_dist {
						continue
					}
					f := flaps + (x_dist-down_needed)/2
					if val, ok := next_positions[wall_pos]; ok {
						if f > val {
							continue
						}
					}
					next_positions[wall_pos] = f
				} else if wall_pos.y == pos.y {
					f := flaps + x_dist/2
					if val, ok := next_positions[wall_pos]; ok {
						if f > val {
							continue
						}
					}
					next_positions[wall_pos] = f
				}
			}
		}

		positions = next_positions
	}

	min := math.MaxInt32
	for _, flaps := range positions {
		if flaps < min {
			min = flaps
		}
	}

	return min
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2025", 19, 1

	data := loader.GetStrings()
	walls := parse_data(data)
	part1 := flap(walls)

	loader.Part = 2
	data = loader.GetStrings()
	walls = parse_data(data)
	part2 := multi_gap(walls)

	loader.Part = 3
	data = loader.GetStrings()
	walls = parse_data(data)
	part3 := multi_gap(walls)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
