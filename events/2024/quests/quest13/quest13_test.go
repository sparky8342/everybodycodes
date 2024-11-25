package quest13

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		"#######",
		"#6769##",
		"S50505E",
		"#97434#",
		"#######",
	}

	got := find_path(grid)
	want := 28

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
