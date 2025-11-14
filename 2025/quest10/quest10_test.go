package quest10

import (
	"testing"
)

func Test1(t *testing.T) {
	board := []string{
		"...SSS.......",
		".S......S.SS.",
		"..S....S...S.",
		"..........SS.",
		"..SSSS...S...",
		".....SS..S..S",
		"SS....D.S....",
		"S.S..S..S....",
		"....S.......S",
		".SSS..SS.....",
		".........S...",
		".......S....S",
		"SS.....S..S..",
	}

	got := in_range(board, 3)
	want := 27

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
