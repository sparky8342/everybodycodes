package quest14

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		".#.##.",
		"##..#.",
		"..##.#",
		".#.##.",
		".###..",
		"###.##",
	}

	got := steps(grid, 10)
	want := 200

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
