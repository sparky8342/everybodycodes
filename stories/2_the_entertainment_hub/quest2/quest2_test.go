package quest2

import (
	"testing"
)

func Test1(t *testing.T) {
	balloons := []byte("GRBGGGBBBRRRRRRRR")

	got := shoot_balloons(balloons)
	want := 7

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
