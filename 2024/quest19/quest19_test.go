package quest19

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"LR",
		"",
		">-IN-",
		"-----",
		"W---<",
	}

	got := decode(data, 1)
	want := "WIN"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"RRLL",
		"",
		"A.VI..>...T",
		".CC...<...O",
		".....EIB.R.",
		".DHB...YF..",
		".....F..G..",
		"D.H........",
	}

	got := decode(data, 100)
	want := "VICTORY"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
