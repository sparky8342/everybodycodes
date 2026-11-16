package quest16

import (
	"testing"
)

func Test1(t *testing.T) {
	tablets := "RRGGBBGRRGGBBG"

	got := sort_tablets(tablets)
	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	tablets := "BBRGGRRGBBRGGBRGBBRRBRRRBGGRRRBGBGGBBRGGRRGBBRGGBRGBBRRBRRRBGGRRRBGBGG"

	got := sort_tablets(tablets)
	want := 26

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
