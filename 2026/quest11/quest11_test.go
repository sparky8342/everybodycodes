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

	portals := parse_data(data)
	got := resonance(portals)
	want := 60

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
