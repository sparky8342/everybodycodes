package quest4

import (
	"testing"
)

func Test1(t *testing.T) {
	gears := []int{128, 64, 32, 16, 8}

	got := turns(gears)
	want := 32400

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	gears := []int{102, 75, 50, 35, 13}

	got := turns(gears)
	want := 15888

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
