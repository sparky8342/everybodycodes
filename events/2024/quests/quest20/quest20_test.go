package quest20

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		"#....S....#",
		"#.........#",
		"#---------#",
		"#.........#",
		"#..+.+.+..#",
		"#.+-.+.++.#",
		"#.........#",
	}

	got := highest_altitude(grid)
	want := 1045

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
