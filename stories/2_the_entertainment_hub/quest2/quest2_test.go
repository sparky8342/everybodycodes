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

func Test2(t *testing.T) {
	balloons := []byte("GGBR")

	got := shoot_balloon_circle(balloons, 5)
	want := 14

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
