package quest18

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		"##########",
		"..#......#",
		"#.P.####P#",
		"#.#...P#.#",
		"##########",
	}

	got := fill(grid)
	want := 11

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	grid := []string{
		"#######################",
		"...P..P...#P....#.....#",
		"#.#######.#.#.#.#####.#",
		"#.....#...#P#.#..P....#",
		"#.#####.#####.#########",
		"#...P....P.P.P.....P#.#",
		"#.#######.#####.#.#.#.#",
		"#...#.....#P...P#.#....",
		"#######################",
	}

	got := fill(grid)
	want := 21

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
