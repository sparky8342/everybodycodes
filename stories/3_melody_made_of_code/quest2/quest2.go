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

func init() {
	dirs = [][]int{
		[]int{0, -1},
		[]int{1, 0},
		[]int{0, 1},
		[]int{-1, 0},
	}
}

func parse_data(data []string) (Pos, Pos) {
	height := len(data)
	width := len(data[0])

	start, bone := Pos{}, Pos{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if data[y][x] == '@' {
				start.x = x
				start.y = y
			} else if data[y][x] == '#' {
				bone.x = x
				bone.y = y
			}
		}
	}

	return start, bone
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

func flood_fill(area Area, start Pos) {
	filled := map[Pos]struct{}{}
	filled[start] = struct{}{}

	queue := []Pos{start}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if pos.x < area.min_x || pos.x > area.max_x || pos.y < area.min_y || pos.y > area.max_y {
			return
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
}

func fill(area Area) {
	for space := range area.grid {
		for _, dir := range dirs {
			neighbour := Pos{
				x: space.x + dir[0],
				y: space.y + dir[1],
			}
			if _, ok := area.grid[neighbour]; !ok {
				flood_fill(area, neighbour)
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

func steps_to_surround_bone(sound Pos, bone Pos) int {
	steps := 0

	area := Area{}
	area.grid = map[Pos]struct{}{}
	area.grid[sound] = struct{}{}
	area.grid[bone] = struct{}{}

	area.min_x = utils.Min([]int{sound.x, bone.x})
	area.max_x = utils.Max([]int{sound.x, bone.x})
	area.min_y = utils.Min([]int{sound.y, bone.y})
	area.max_y = utils.Max([]int{sound.y, bone.y})

	dir := -1
	for {
		dir = (dir + 1) % 4

		next := Pos{
			x: sound.x + dirs[dir][0],
			y: sound.y + dirs[dir][1],
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

		if surrounded(area, bone) {
			return steps
		}
	}
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "3", 2, 1

	data := loader.GetStrings()
	start, bone := parse_data(data)
	part1 := steps_to_bone(start, bone)

	loader.Part = 2
	data = loader.GetStrings()
	start, bone = parse_data(data)
	part2 := steps_to_surround_bone(start, bone)

	fmt.Printf("%d %d %d\n", part1, part2, -1)
}
