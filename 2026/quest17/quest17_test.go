package quest17

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		"####.######",
		"#.#.##..#..",
		"##.##.##..#",
		"#.#.###.###",
		"#.###.###.#",
		"###.##.#.##",
	}

	got := shortest_path(grid)
	want := 21

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	grid := []string{
		"####.######",
		"#.#.##..#..",
		"##.##.##..#",
		"#.#.###.###",
		"#.###.###.#",
		"###.##.#.##",
	}

	got := edge_distance(grid)
	want := 94

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
