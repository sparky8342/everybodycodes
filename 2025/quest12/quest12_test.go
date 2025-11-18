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

	got := shoot_barrels(grid, 1)
	want := 16

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	grid := []string{
		"9589233445",
		"9679121695",
		"8469121876",
		"8352919876",
		"7342914327",
		"7234193437",
		"6789193538",
		"6781219648",
		"5691219769",
		"5443329859",
	}

	got := shoot_barrels(grid, 2)
	want := 58

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
