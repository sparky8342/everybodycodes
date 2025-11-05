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
