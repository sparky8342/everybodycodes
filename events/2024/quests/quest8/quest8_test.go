package quest8

import (
	"testing"
)

func Test1(t *testing.T) {
	blocks := 13

	got := build(blocks)
	want := 21

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	multiplier := 3
	mod := 5
	blocks := 50

	got := build_part2(blocks, multiplier, mod)
	want := 27

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	multiplier := 2
	mod := 5
	blocks := 160

	got := build_part3(blocks, multiplier, mod)
	want := 2

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
