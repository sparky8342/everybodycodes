package quest11

import (
	"testing"
)

func Test1(t *testing.T) {
	ducks := []int{9, 1, 1, 4, 9, 6}

	move_ducks(ducks, 10)
	got := checksum(ducks)
	want := 109

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	ducks := []int{9, 1, 1, 4, 9, 6}

	got := balance_ducks(ducks)
	want := 11

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	ducks := []int{805, 706, 179, 48, 158, 150, 232, 885, 598, 524, 423}

	got := balance_ducks(ducks)
	want := 1579

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
