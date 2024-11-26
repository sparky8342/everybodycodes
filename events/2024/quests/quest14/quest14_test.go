package quest14

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []byte("U5,R3,D2,L5,U4,R5,D2")

	got := max_height(data)
	want := 7

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"U5,R3,D2,L5,U4,R5,D2",
		"U6,L1,D2,R3,U2,L1",
	}

	got := unique_segments(data)
	want := 32

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"U5,R3,D2,L5,U4,R5,D2",
		"U6,L1,D2,R3,U2,L1",
	}

	got := murkiness(data)
	want := 5

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
