package quest11

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"portals: 10",
		"",
		"5 10",
		"2 7",
	}

	got := resonance(data, 1)
	want := 60

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"portals: 10",
		"",
		"5 10",
		"2 7",
	}

	got := resonance_extra_connection(data, 3)
	want := 6

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"portals: 32",
		"",
		"1 10",
		"20 30",
		"15 29",
		"13 10",
	}

	got := resonance_with_crystals(data, 3)
	want := 854

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	data := []string{
		"portals: 64",
		"",
		"1 40",
		"3 35",
		"5 42",
		"7 50",
	}

	got := resonance_with_crystals(data, 3)
	want := 9634

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
