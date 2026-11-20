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

	got := count_segments(data)
	want := 75

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
