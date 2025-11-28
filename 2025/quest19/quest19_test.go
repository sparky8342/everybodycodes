package quest19

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"7,7,2",
		"12,0,4",
		"15,5,3",
		"24,1,6",
		"28,5,5",
		"40,8,2",
	}

	walls := parse_data(data)
	got := flap(walls)
	want := 24

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"7,7,2",
		"7,1,3",
		"12,0,4",
		"15,5,3",
		"24,1,6",
		"28,5,5",
		"40,3,3",
		"40,8,2",
	}

	walls := parse_data(data)
	got := multi_flap(walls)
	want := 22

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
