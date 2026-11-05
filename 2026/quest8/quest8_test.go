package quest8

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		"....1.9.6.1.7....",
		"...9.2.3.3.8.1...",
		"..7.8.3.5.1.4.3..",
		".4.6.7.4.1.1.8.9.",
		"4.5.3.6.5.8.9.4.9",
		".9.5.1.5.6.2.8.9.",
		"..7.7.6.4.6.2.4..",
		"...9.3.8.3.8.4...",
		"....1.6.6.2.8....",
	}

	got := energy_needed_nw(grid)
	want := 32

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	grid := []string{
		"....1.9.6.1.7....",
		"...9.2.3.3.8.1...",
		"..7.8.3.5.1.4.3..",
		".4.6.7.4.1.1.8.9.",
		"4.5.3.6.5.8.9.4.9",
		".9.5.1.5.6.2.8.9.",
		"..7.7.6.4.6.2.4..",
		"...9.3.8.3.8.4...",
		"....1.6.6.2.8....",
	}

	got := all_energy_needed(grid)
	want := 187

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
