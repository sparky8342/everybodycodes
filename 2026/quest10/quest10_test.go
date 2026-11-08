package quest10

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		"###########",
		"#####.#####",
		"###.....###",
		"##.......##",
		"##.>>.<<.##",
		"#v...>...v#",
		"##...v...##",
		"##.......##",
		"###.....###",
		"#####.#####",
		"###########",
	}

	ghosts := find_ghosts(grid)
	got := move_ghosts(grid, ghosts)
	want := 29

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
