package quest10

import (
	"fmt"
	"loader"
	"strings"
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

type State struct {
	collisions []int
	rounds     int
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

func move_ghosts_one_round(grid []string, ghosts []*Ghost, collision_limit int) ([]*Ghost, int) {
	collisions := 0
	moved_to := map[Pos][]*Ghost{}

	for _, ghost := range ghosts {
		if grid[ghost.y+ghost.dy][ghost.x+ghost.dx] == '#' {
			ghost.dx, ghost.dy = -ghost.dy, ghost.dx
			collisions++
			ghost.collisions++
		} else {
			ghost.y += ghost.dy
			ghost.x += ghost.dx
		}
		if ghost.collisions < collision_limit {
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
		if ghosts[i].collisions >= collision_limit {
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

func jump(grid []string, ghosts []*Ghost, collision_limit int, rounds int, state State) ([]*Ghost, int) {
	diffs := make([]int, len(state.collisions))
	max_diff := 0
	max_collisions := 0
	for i := 0; i < len(state.collisions); i++ {
		diffs[i] = ghosts[i].collisions - state.collisions[i]
		if diffs[i] > max_diff {
			max_diff = diffs[i]
		}
		if ghosts[i].collisions > max_collisions {
			max_collisions = ghosts[i].collisions
		}
	}
	round_diff := rounds - state.rounds

	collisions_left := collision_limit - max_collisions
	repeats := collisions_left / max_diff
	if repeats > 0 {
		repeats--
	}

	for i, ghost := range ghosts {
		ghost.collisions += diffs[i] * repeats
	}
	rounds += round_diff * repeats

	return ghosts, rounds
}

func move_ghosts_until_gone(grid []string, ghosts []*Ghost, collision_limit int, rounds int) int {
	states := map[string]State{}

	for len(ghosts) > 0 {
		l := len(ghosts)
		ghosts, _ = move_ghosts_one_round(grid, ghosts, collision_limit)
		rounds++
		if len(ghosts) == 0 {
			break
		}
		if len(ghosts) < l {
			l = len(ghosts)
			states = map[string]State{}
		}

		hash := hash_ghosts(ghosts)
		if state, ok := states[hash]; ok {
			l := len(ghosts)
			ghosts, rounds = jump(grid, ghosts, collision_limit, rounds, state)
			if len(ghosts) < l {
				states = map[string]State{}
			}
			hash = hash_ghosts(ghosts)
		}
		state := State{collisions: []int{}}
		for _, ghost := range ghosts {
			state.collisions = append(state.collisions, ghost.collisions)
		}
		state.rounds = rounds
		states[hash] = state
	}

	return rounds
}

func hash_ghosts(ghosts []*Ghost) string {
	strs := []string{}
	for _, ghost := range ghosts {
		str := fmt.Sprintf("%d_%d_%d_%d", ghost.x, ghost.y, ghost.dx, ghost.dy)
		strs = append(strs, str)
	}
	return strings.Join(strs, ":")
}

func Run() {
	loader.Event, loader.Quest, loader.Part = "2026", 10, 1

	grid := loader.GetStrings()
	ghosts := find_ghosts(grid)
	part1 := move_ghosts(grid, ghosts)

	loader.Part = 2
	grid = loader.GetStrings()
	ghosts = find_ghosts(grid)
	part2 := move_ghosts_until_gone(grid, ghosts, 10000, 0)

	loader.Part = 3
	grid = loader.GetStrings()
	ghosts = find_ghosts(grid)
	part3 := move_ghosts_until_gone(grid, ghosts, 1000000000000, 0)

	fmt.Printf("%d %d %d\n", part1, part2, part3)
}
