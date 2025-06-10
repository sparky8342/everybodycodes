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
