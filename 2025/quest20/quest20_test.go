package quest20

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		"T#TTT###T##",
		".##TT#TT##.",
		"..T###T#T..",
		"...##TT#...",
		"....T##....",
		".....#.....",
	}

	got := count_pairs(grid)
	want := 7

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
