package quest10

import (
	"fmt"
	"loader"
)

type Pos struct {
	x int
	y int
}

type Ghost struct {
	x          int
	y          int
	dx         int
	dy         int
	collisions int
}

func find_ghosts(grid []string) []*Ghost {
	height := len(grid)
	width := len(grid[0])

	ghosts := []*Ghost{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch grid[y][x] {
			case '^':
				ghosts = append(ghosts, &Ghost{x: x, y: y, dx: 0, dy: -1})
			case '>':
				ghosts = append(ghosts, &Ghost{x: x, y: y, dx: 1, dy: 0})
			case 'v':
				ghosts = append(ghosts, &Ghost{x: x, y: y, dx: 0, dy: 1})
			case '<':
				ghosts = append(ghosts, &Ghost{x: x, y: y, dx: -1, dy: 0})
			}
		}
	}

	return ghosts
}

func move_ghosts_one_round(grid []string, ghosts []*Ghost, collisions_limit int) ([]*Ghost, int) {
	collisions := 0
	moved_to := map[Pos][]*Ghost{}

	for _, ghost := range ghosts {
		if grid[ghost.y+ghost.dy][ghost.x+ghost.dx] == '#' {
			if ghost.dx == 0 && ghost.dy == -1 {
				ghost.dx, ghost.dy = 1, 0
			} else if ghost.dx == 1 && ghost.dy == 0 {
				ghost.dx, ghost.dy = 0, 1
			} else if ghost.dx == 0 && ghost.dy == 1 {
				ghost.dx, ghost.dy = -1, 0
			} else if ghost.dx == -1 && ghost.dy == 0 {
				ghost.dx, ghost.dy = 0, -1
			}
			collisions++
			ghost.collisions++
		} else {
			ghost.y += ghost.dy
			ghost.x += ghost.dx
		}
		if ghost.collisions < collisions_limit {
			pos := Pos{x: ghost.x, y: ghost.y}
			if _, ok := moved_to[pos]; !ok {
				moved_to[pos] = []*Ghost{ghost}
			} else {
				moved_to[pos] = append(moved_to[pos], ghost)
			}
		}
	}

	for _, ghosts := range moved_to {
		if len(ghosts) > 1 {
			collisions += len(ghosts)
			for _, ghost := range ghosts {
				ghost.dx *= -1
				ghost.dy *= -1
				ghost.collisions++
			}
		}
	}

	i := 0
	for i < len(ghosts) {
		if ghosts[i].collisions >= collisions_limit {
			ghosts = append(ghosts[:i], ghosts[i+1:]...)
		} else {
			i++
		}
	}

	return ghosts, collisions
}

func move_ghosts(grid []string, ghosts []*Ghost) int {
	collisions := 0
	for i := 0; i < 10; i++ {
		var round_collisions int
		ghosts, round_collisions = move_ghosts_one_round(grid, ghosts, 100)
		collisions += round_collisions
	}
	return collisions
}

func move_ghosts_until_gone(grid []string, ghosts []*Ghost, collision_limit int) int {
	rounds := 0
	for len(ghosts) > 0 {
		ghosts, _ = move_ghosts_one_round(grid, ghosts, collision_limit)
		rounds++
	}
	return rounds
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 10, 1

	grid := loader.GetStrings()
	ghosts := find_ghosts(grid)
	part1 := move_ghosts(grid, ghosts)

	loader.Part = 2
	grid = loader.GetStrings()
	ghosts = find_ghosts(grid)
	part2 := move_ghosts_until_gone(grid, ghosts, 10000)

	loader.Part = 3
	part3 := -1

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
