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
