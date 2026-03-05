package quest2

import (
	"fmt"
	"loader"
	"utils"
)

type Area struct {
	grid  map[Pos]struct{}
	min_x int
	max_x int
	min_y int
	max_y int
}

type Pos struct {
	x int
	y int
}

var dirs [][]int
var dir_sequence [][]int

func init() {
	dirs = [][]int{
		[]int{0, -1},
		[]int{1, 0},
		[]int{0, 1},
		[]int{-1, 0},
	}

	dir_sequence = [][]int{
		[]int{0, -1},
		[]int{0, -1},
		[]int{0, -1},
		[]int{1, 0},
		[]int{1, 0},
		[]int{1, 0},
		[]int{0, 1},
		[]int{0, 1},
		[]int{0, 1},
		[]int{-1, 0},
		[]int{-1, 0},
		[]int{-1, 0},
	}
}

func parse_data(data []string) (Pos, []Pos) {
	height := len(data)
	width := len(data[0])

	start := Pos{}
	bones := []Pos{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if data[y][x] == '@' {
				start.x = x
				start.y = y
			} else if data[y][x] == '#' {
				bones = append(bones, Pos{x: x, y: y})
			}
		}
	}

	return start, bones
}

func steps_to_bone(sound Pos, bone Pos) int {
	steps := 0

	grid := map[Pos]struct{}{}
	grid[sound] = struct{}{}

	dir := -1
	for {
		dir = (dir + 1) % 4

		next := Pos{
			x: sound.x + dirs[dir][0],
			y: sound.y + dirs[dir][1],
		}

		if _, ok := grid[next]; ok {
			continue
		}

		sound = next
		grid[sound] = struct{}{}
		steps++

		if sound == bone {
			return steps
		}
	}
}

func surrounded(area Area, bone Pos) bool {
	for _, dir := range dirs {
		n := Pos{
			x: bone.x + dir[0],
			y: bone.y + dir[1],
		}
		if _, ok := area.grid[n]; !ok {
			return false
		}
	}
	return true
}

func bones_surrounded(area Area, bones []Pos) bool {
	for _, bone := range bones {
		if !surrounded(area, bone) {
			return false
		}
	}
	return true
}

func flood_fill(area Area, start Pos) map[Pos]struct{} {
	filled := map[Pos]struct{}{}
	filled[start] = struct{}{}

	queue := []Pos{start}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if pos.x < area.min_x || pos.x > area.max_x || pos.y < area.min_y || pos.y > area.max_y {
			return filled
		}

		for _, dir := range dirs {
			next := Pos{
				x: pos.x + dir[0],
				y: pos.y + dir[1],
			}

			_, in_grid := area.grid[next]
			_, in_filled := filled[next]
			if in_grid || in_filled {
				continue
			}

			queue = append(queue, next)
			filled[next] = struct{}{}
		}
	}

	for space := range filled {
		area.grid[space] = struct{}{}
	}

	return filled
}

func fill(area Area) {
	seen := map[Pos]struct{}{}
	for space := range area.grid {
		for _, dir := range dirs {
			neighbour := Pos{
				x: space.x + dir[0],
				y: space.y + dir[1],
			}
			if _, ok := seen[neighbour]; ok {
				continue
			}
			if _, ok := area.grid[neighbour]; !ok {
				filled := flood_fill(area, neighbour)
				for space := range filled {
					seen[space] = struct{}{}
				}
			}
		}
	}
}

func print_grid(grid map[Pos]struct{}, sound Pos, bone Pos) {
	for y := -20; y < 20; y++ {
		for x := -20; x < 20; x++ {
			pos := Pos{x: x, y: y}
			if pos == sound {
				fmt.Print("@")
			} else if pos == bone {
				fmt.Print("#")
			} else if _, ok := grid[pos]; ok {
				fmt.Print("+")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func steps_to_surround_bones(sound Pos, bones []Pos, full_sequence bool) int {
	steps := 0

	area := Area{}
	area.grid = map[Pos]struct{}{}
	area.grid[sound] = struct{}{}
	for _, bone := range bones {
		area.grid[bone] = struct{}{}

		area.min_x = utils.Min([]int{area.min_x, bone.x})
		area.max_x = utils.Max([]int{area.max_x, bone.x})
		area.min_y = utils.Min([]int{area.min_y, bone.y})
		area.max_y = utils.Max([]int{area.max_y, bone.y})
	}

	move_dirs := dirs
	if full_sequence {
		move_dirs = dir_sequence
	}

	dir := -1
	for {
		dir = (dir + 1) % len(move_dirs)

		next := Pos{
			x: sound.x + move_dirs[dir][0],
			y: sound.y + move_dirs[dir][1],
		}

		if _, ok := area.grid[next]; ok {
			continue
		}

		sound = next
		area.grid[sound] = struct{}{}
		steps++

		area.min_x = utils.Min([]int{area.min_x, next.x})
		area.max_x = utils.Max([]int{area.max_x, next.x})
		area.min_y = utils.Min([]int{area.min_y, next.y})
		area.max_y = utils.Max([]int{area.max_y, next.y})

		fill(area)

		if bones_surrounded(area, bones) {
			return steps
		}
	}
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "3", 2, 1

	data := loader.GetStrings()
	start, bones := parse_data(data)
	part1 := steps_to_bone(start, bones[0])

	loader.Part = 2
	data = loader.GetStrings()
	start, bones = parse_data(data)
	part2 := steps_to_surround_bones(start, bones, false)

	loader.Part = 3
	data = loader.GetStrings()
	start, bones = parse_data(data)
	part3 := steps_to_surround_bones(start, bones, true)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
