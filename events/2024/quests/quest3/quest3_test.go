package quest3

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"..........",
		"..###.##..",
		"...####...",
		"..######..",
		"..######..",
		"...####...",
		"..........",
	}

	got := dig(data)
	want := 35

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
