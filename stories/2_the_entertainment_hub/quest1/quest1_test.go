package quest1

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"*.*.*.*.*.*.*.*.*",
		".*.*.*.*.*.*.*.*.",
		"*.*.*...*.*...*..",
		".*.*.*.*.*...*.*.",
		"*.*.....*...*.*.*",
		".*.*.*.*.*.*.*.*.",
		"*...*...*.*.*.*.*",
		".*.*.*.*.*.*.*.*.",
		"*.*.*...*.*.*.*.*",
		".*...*...*.*.*.*.",
		"*.*.*.*.*.*.*.*.*",
		".*.*.*.*.*.*.*.*.",
		"",
		"RRRLRLRRRRRL",
		"LLLLRLRRRRRR",
		"RLLLLLRLRLRL",
		"LRLLLRRRLRLR",
		"LLRLLRLLLRRL",
		"LRLRLLLRRRRL",
		"LRLLLLLLRLLL",
		"RRLLLRLLRLRR",
		"RLLLLLRLLLRL",
	}

	grid, tokens := parse_data(data)
	got := play_tokens(grid, tokens)
	want := 26

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
