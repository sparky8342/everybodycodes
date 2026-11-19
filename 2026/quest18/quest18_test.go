package quest18

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"x=0 y=5",
		"x=5 y=7",
		"x=6 y=8",
		"x=7 y=1",
		"x=8 y=3",
		"x=10 y=6",
		"x=15 y=3",
		"x=18 y=4",
	}

	got := material_needed(data)
	want := 96

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"x=0 y=5 width=1",
		"x=1 y=4 width=3",
		"x=5 y=3 width=2",
		"x=7 y=9 width=1",
		"x=8 y=8 width=1",
		"x=9 y=9 width=3",
		"x=12 y=3 width=1",
		"x=16 y=7 width=1",
		"x=19 y=3 width=5",
		"x=25 y=5 width=1",
		"x=29 y=1 width=2",
		"x=31 y=6 width=1",
	}

	got := material_needed(data)
	want := 205

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
