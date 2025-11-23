package quest15

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []byte("L6,L3,L6,R3,L6,L3,L3,R6,L6,R6,L6,L6,R3,L3,L3,R3,R3,L6,L6,L3")

	moves := parse_data(data)
	got := find_exit(moves)
	want := 16

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []byte("L6,L3,L6,R3,L6,L3,L3,R6,L6,R6,L6,L6,R3,L3,L3,R3,R3,L6,L6,L3")

	moves := parse_data(data)
	got := find_exit_large(moves)
	want := 16

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
