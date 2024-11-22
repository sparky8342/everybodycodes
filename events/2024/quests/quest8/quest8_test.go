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
