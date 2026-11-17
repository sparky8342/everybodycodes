package quest16

import (
	"testing"
)

func Test1(t *testing.T) {
	tablets := "RRGGBBGRRGGBBG"

	got, _ := sort_tablets(tablets, 1)
	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	tablets := "BBRGGRRGBBRGGBRGBBRRBRRRBGGRRRBGBGGBBRGGRRGBBRGGBRGBBRRBRRRBGGRRRBGBGG"

	got, _ := sort_tablets(tablets, 1)
	want := 26

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	tablets := "RRYGYBYG"

	_, got := sort_tablets(tablets, 10)
	want := 1053

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	tablets := "RRYGGYBBYG"

	_, got := sort_tablets(tablets, 100)
	want := 135154

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
