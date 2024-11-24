package quest12

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

type Target struct {
	x        int
	y        int
	strength int
}

func find_things(grid []string) (map[byte]Pos, []Target) {
	height := len(grid)
	width := len(grid[0])

	cannons := map[byte]Pos{}
	targets := []Target{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			space := grid[y][x]
			if space >= 'A' && space <= 'C' {
				cannons[space] = Pos{x: x, y: y}
			} else if space == 'T' {
				targets = append(targets, Target{x: x, y: y, strength: 1})
			} else if space == 'H' {
				targets = append(targets, Target{x: x, y: y, strength: 2})
			}
		}
	}

	return cannons, targets
}

func fire(cannons map[byte]Pos, target Target) int {
	for i, name := range []byte{'C', 'B', 'A'} {
		cannon := cannons[name]
		for power := 1; ; power++ {
			ball := cannon
			ball.x += power * 2
			ball.y -= power
			for ball.y < target.y {
				ball.y++
				ball.x++
			}
			if ball.x == target.x {
				return (3 - i) * power * target.strength
			} else if ball.x > target.x {
				break
			}
		}
	}
	return -1
}

func fire_cannons(grid []string, cannons map[byte]Pos, targets []Target) int {
	power := 0
	for _, target := range targets {
		power += fire(cannons, target)
	}
	return power
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2024", 12, 1

	grid := loader.GetStrings()
	cannons, targets := find_things(grid)
	part1 := fire_cannons(grid, cannons, targets)

	loader.Part = 2
	grid = loader.GetStrings()
	cannons, targets = find_things(grid)
	part2 := fire_cannons(grid, cannons, targets)

	part3 := -1
	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
