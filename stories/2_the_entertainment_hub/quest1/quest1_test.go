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

func Test2(t *testing.T) {
	data := []string{
		"*.*.*.*.*.*.*.*.*.*.*.*.*",
		".*.*.*.*.*.*.*.*.*.*.*.*.",
		"..*.*.*.*...*.*...*.*.*..",
		".*...*.*.*.*.*.*.....*.*.",
		"*.*...*.*.*.*.*.*...*.*.*",
		".*.*.*.*.*.*.*.*.......*.",
		"*.*.*.*.*.*.*.*.*.*...*..",
		".*.*.*.*.*.*.*.*.....*.*.",
		"*.*...*.*.*.*.*.*.*.*....",
		".*.*.*.*.*.*.*.*.*.*.*.*.",
		"*.*.*.*.*.*.*.*.*.*.*.*.*",
		".*.*.*.*.*.*.*.*.*...*.*.",
		"*.*.*.*.*.*.*.*.*...*.*.*",
		".*.*.*.*.*.*.*.*.....*.*.",
		"*.*.*.*.*.*.*.*...*...*.*",
		".*.*.*.*.*.*.*.*.*.*.*.*.",
		"*.*.*...*.*.*.*.*.*.*.*.*",
		".*...*.*.*.*...*.*.*...*.",
		"*.*.*.*.*.*.*.*.*.*.*.*.*",
		".*.*.*.*.*.*.*.*.*.*.*.*.",
		"",
		"RRRLLRRRLLRLRRLLLRLR",
		"RRRRRRRRRRLRRRRRLLRR",
		"LLLLLLLLRLRRLLRRLRLL",
		"RRRLLRRRLLRLLRLLLRRL",
		"RLRLLLRRLRRRLRRLRRRL",
		"LLLLLLLLRLLRRLLRLLLL",
		"LRLLRRLRLLLLLLLRLRRL",
		"LRLLRRLLLRRRRRLRRLRR",
		"LRLLRRLRLLRLRRLLLRLL",
		"RLLRRRRLRLRLRLRLLRRL",
	}

	grid, tokens := parse_data(data)
	got := maximise_tokens(grid, tokens)
	want := 115

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
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
	}

	grid, tokens := parse_data(data)
	got := unique_slots(grid, tokens)
	want := "13 43"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
