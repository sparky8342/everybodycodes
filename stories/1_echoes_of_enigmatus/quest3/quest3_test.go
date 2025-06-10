package quest3

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"x=1 y=2",
		"x=2 y=3",
		"x=3 y=4",
		"x=4 y=4",
	}

	snails := parse_data(data)
	move_snails(snails, 100)

	got := position_sum(snails)
	want := 1310

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"x=12 y=2",
		"x=8 y=4",
		"x=7 y=1",
		"x=1 y=5",
		"x=1 y=3",
	}

	snails := parse_data(data)
	got := all_top(snails)
	want := 14

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"x=3 y=1",
		"x=3 y=9",
		"x=1 y=5",
		"x=4 y=10",
		"x=5 y=3",
	}

	snails := parse_data(data)
	got := all_top(snails)
	want := 13659

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
