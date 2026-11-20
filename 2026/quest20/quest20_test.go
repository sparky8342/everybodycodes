package quest20

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"5 - 4 = 9 - 8",
		"2 + 1 = 3 + 0",
		"6 + 7 = 4 + 9",
		"1 + 2 = 3 + 4",
	}

	got := count_segments(data, 10)
	want := "75"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"10 + 2A = 13 + 27",
		"98 - 54 = 76 - 32",
		"10 + 12 = 30 - 9",
		"20 - 10 = 5 + 5",
	}

	got := count_segments(data, 11)
	want := "101"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
