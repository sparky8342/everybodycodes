package quest1

import (
	"testing"
)

func Test1(t *testing.T) {
	pines := []int{3, 11, 12, 16, 24, 25}

	got := pines_needed(pines)
	want := 11

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	pines := []int{3, 11, 12, 16, 24, 25}

	got := pines_needed_frame(pines)
	want := 9

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	pines := []int{3, 11, 12, 16, 24, 25}

	got := pines_needed_square_empty_centre(pines)
	want := 61

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
