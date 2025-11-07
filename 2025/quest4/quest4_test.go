package quest4

import (
	"testing"
)

func Test1(t *testing.T) {
	gears := []int{128, 64, 32, 16, 8}

	got := turns(gears)
	want := 32400

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	gears := []int{102, 75, 50, 35, 13}

	got := turns(gears)
	want := 15888

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	gears := []int{128, 64, 32, 16, 8}

	got := turns_needed(gears)
	want := 625000000000

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	gears := []int{102, 75, 50, 35, 13}

	got := turns_needed(gears)
	want := 1274509803922

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test5(t *testing.T) {
	data := []string{
		"5",
		"5|10",
		"10|20",
		"5",
	}

	gears := parse_data(data)
	got := turns_linked(gears)
	want := 400

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test6(t *testing.T) {
	data := []string{
		"5",
		"7|21",
		"18|36",
		"27|27",
		"10|50",
		"10|50",
		"11",
	}

	gears := parse_data(data)
	got := turns_linked(gears)
	want := 6818

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
