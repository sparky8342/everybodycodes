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

func Test4(t *testing.T) {
	data := []string{
		"U20,L1,B1,L2,B1,R2,L1,F1,U1",
		"U10,F1,B1,R1,L1,B1,L1,F1,R2,U1",
		"U30,L2,F1,R1,B1,R1,F2,U1,F1",
		"U25,R1,L2,B1,U1,R2,F1,L2",
		"U16,L1,B1,L1,B3,L1,B1,F1",
	}

	got := murkiness(data)
	want := 46

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
