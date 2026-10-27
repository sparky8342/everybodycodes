package quest1

import (
	"testing"
)

func Test1(t *testing.T) {
	pines := []int{3, 11, 12, 16, 24, 25}

	got := pines_needed(pines, next_square)
	want := 11

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	pines := []int{3, 11, 12, 16, 24, 25}

	got := pines_needed(pines, next_frame)
	want := 9

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	pines := []int{3, 11, 12, 16, 24, 25}

	got := pines_needed(pines, next_square_empty_centre)
	want := 61

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
