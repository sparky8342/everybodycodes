package quest9

import (
	"testing"
)

func Test1(t *testing.T) {
	sparks := []int{2, 4, 7, 16}

	available = dots_part1
	got := calc_beatles(sparks)
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	sparks := []int{33, 41, 55, 99}

	available = dots_part2
	got := calc_beatles(sparks)
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	sparks := []int{156488, 352486, 546212}

	available = dots_part3
	got := calc_beatles_split(sparks)
	want := 10449

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
