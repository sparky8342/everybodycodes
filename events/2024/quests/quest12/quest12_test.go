package quest12

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		".............",
		".C...........",
		".B......T....",
		".A......T.T..",
		"=============",
	}

	cannons, targets := find_things(grid)
	got := fire_cannons(grid, cannons, targets)
	want := 13

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	grid := []string{
		".............",
		".C...........",
		".B......H....",
		".A......T.H..",
		"=============",
	}

	cannons, targets := find_things(grid)
	got := fire_cannons(grid, cannons, targets)
	want := 22

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
