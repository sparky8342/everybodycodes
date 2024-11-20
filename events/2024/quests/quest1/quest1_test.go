package quest1

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []byte("ABBAC")
	got := calculate_potions(data)
	want := 5

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []byte("AxBCDDCAxD")
	got := calculate_pairs(data)
	want := 28

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []byte("xBxAAABCDxCC")
	got := calculate_triples(data)
	want := 30

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
