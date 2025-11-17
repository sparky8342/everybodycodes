package quest11

import (
	"testing"
)

func Test1(t *testing.T) {
	ducks := []int{9, 1, 1, 4, 9, 6}

	ducks = move_ducks(ducks, 10)
	got := checksum(ducks)
	want := 109

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
