package quest12

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		"989611",
		"857782",
		"746543",
		"766789",
	}

	got := shoot_barrels(grid)
	want := 16

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
