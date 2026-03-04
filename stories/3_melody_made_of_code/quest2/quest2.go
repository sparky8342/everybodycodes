package quest2

import (
	"fmt"
	"loader"
)

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

	visited := map[Pos]struct{}{}
	visited[sound] = struct{}{}

	dir := -1
	for {
		dir = (dir + 1) % 4

		next := Pos{
			x: sound.x + dirs[dir][0],
			y: sound.y + dirs[dir][1],
		}

		if _, ok := visited[next]; ok {
			continue
		}

		sound = next
		visited[sound] = struct{}{}
		steps++

		if sound == bone {
			return steps
		}
	}
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "3", 2, 1

	data := loader.GetStrings()
	start, bone := parse_data(data)
	part1 := steps_to_bone(start, bone)

	fmt.Printf("%d %d %d\n", part1, -1, -1)
}
