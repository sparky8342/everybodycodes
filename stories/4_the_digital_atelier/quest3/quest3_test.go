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

	grid := parse_data(data)
	draw_pattern(grid)
	got := count_isolated(grid)
	want := 27

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"width=30",
		"height=10",
		"horizontal-offsets=10011",
		"vertical-offsets=11011",
	}

	grid := parse_data(data)
	draw_pattern(grid)
	colour_in(grid)
	got := count_isolated(grid)

	want := 15

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"width=40",
		"height=12",
		"horizontal-offsets=11100",
		"vertical-offsets=001101",
	}

	grid := parse_data(data)
	draw_pattern(grid)
	colour_in(grid)
	got := count_isolated(grid)

	want := 7

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	data := []string{
		"width=100",
		"height=70",
		"horizontal-offsets=111101111101101111000100100110",
		"vertical-offsets=110100001110111011101000001111",
	}

	grid := parse_data(data)
	draw_pattern(grid)
	colour_in(grid)
	got := count_isolated(grid)

	want := 269

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
