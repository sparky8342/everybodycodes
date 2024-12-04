package quest7

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"A:+,-,=,=",
		"B:+,=,-,+",
		"C:=,-,+,+",
		"D:=,=,=,+",
	}

	chariots := parse_data(data)
	got := race(chariots, 10)

	want := "BDCA"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"A:+,-,=,=",
		"B:+,=,-,+",
		"C:=,-,+,+",
		"D:=,=,=,+",
	}

	grid := []string{
		"S+===",
		"-   +",
		"=+=-+",
	}

	chariots := parse_data(data)
	track := parse_grid(grid)
	got := race_track(chariots, track, 10)

	want := "DCBA"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
