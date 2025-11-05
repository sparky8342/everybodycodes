package quest3

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []byte("10,5,1,10,3,8,5,2,2")

	crates := parse_data(data)
	got := largest_set(crates)
	want := 29

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []byte("4,51,13,64,57,51,82,57,16,88,89,48,32,49,49,2,84,65,49,43,9,13,2,3,75,72,63,48,61,14,40,77")

	crates := parse_data(data)
	got := smallest_20(crates)
	want := 781

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []byte("4,51,13,64,57,51,82,57,16,88,89,48,32,49,49,2,84,65,49,43,9,13,2,3,75,72,63,48,61,14,40,77")

	crates := parse_data(data)
	got := smallest_no_sets(crates)
	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
