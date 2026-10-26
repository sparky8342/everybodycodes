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
