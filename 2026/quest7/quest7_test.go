package quest7

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"id=1, energy=5, connections=2,5,7",
		"id=2, energy=9, connections=6,7",
		"id=3, energy=3, connections=4",
		"id=4, energy=8, connections=3",
		"id=5, energy=1, connections=6,7",
		"id=6, energy=7, connections=1,7",
		"id=7, energy=6, connections=5",
	}

	clouds := parse_data(data)
	got := charge(clouds, 5)
	want := 102

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"id=1, energy=5, connections=2,5,7",
		"id=2, energy=9, connections=6,7",
		"id=3, energy=3, connections=4",
		"id=4, energy=8, connections=3",
		"id=5, energy=1, connections=6,7",
		"id=6, energy=7, connections=1,7",
		"id=7, energy=6, connections=5",
	}

	clouds := parse_data(data)
	got := charge(clouds, 10)
	want := 1316

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"id=1, energy=5, connections=2,3,5,7",
		"id=2, energy=9, connections=1,3,4,5,6,7",
		"id=3, energy=3, connections=1,2,4,6,7",
		"id=4, energy=8, connections=1,2,7",
		"id=5, energy=1, connections=2,3,4,6",
		"id=6, energy=7, connections=1,2,3,4",
		"id=7, energy=6, connections=1,2,4,6",
	}

	clouds := parse_data(data)
	got := charge(clouds, 10)
	want := 1214633

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
