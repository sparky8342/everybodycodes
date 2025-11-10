package quest6

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []byte("ABabACacBCbca")

	got := total_A_pairs(data)
	want := 5

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []byte("ABabACacBCbca")

	got := total_pairs(data)
	want := 11

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
