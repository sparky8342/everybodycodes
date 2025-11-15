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

func Test2(t *testing.T) {
	board := []string{
		"...SSS##.....",
		".S#.##..S#SS.",
		"..S.##.S#..S.",
		".#..#S##..SS.",
		"..SSSS.#.S.#.",
		".##..SS.#S.#S",
		"SS##.#D.S.#..",
		"S.S..S..S###.",
		".##.S#.#....S",
		".SSS.#SS..##.",
		"..#.##...S##.",
		".#...#.S#...S",
		"SS...#.S.#S..",
	}

	got := find_max_sheep(board, 3)
	want := 27

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	board := []string{
		"SSS",
		"..#",
		"#.#",
		"#D.",
	}

	got := find_sequences(board)
	want := 15

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	board := []string{
		"SSS",
		"..#",
		"..#",
		".##",
		".D#",
	}

	got := find_sequences(board)
	want := 8

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test5(t *testing.T) {
	board := []string{
		"..S..",
		".....",
		"..#..",
		".....",
		"..D..",
	}

	got := find_sequences(board)
	want := 44

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test6(t *testing.T) {
	board := []string{
		".SS.S",
		"#...#",
		"...#.",
		"##..#",
		".####",
		"##D.#",
	}

	got := find_sequences(board)
	want := 4406

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test7(t *testing.T) {
	board := []string{
		"SSS.S",
		".....",
		"#.#.#",
		".#.#.",
		"#.D.#",
	}

	got := find_sequences(board)
	want := 13033988838

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
