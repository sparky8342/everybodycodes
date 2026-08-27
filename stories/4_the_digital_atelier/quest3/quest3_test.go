package quest3

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"width=30",
		"height=10",
		"horizontal-offsets=10011",
		"vertical-offsets=11011",
	}

	width, height, h_offsets, v_offsets := parse_data(data)
	got := draw_pattern(width, height, h_offsets, v_offsets)
	want := 27

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
