package quest9

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"id=1, growth=3, limit=10",
		"id=2, growth=5, limit=20",
		"id=3, growth=7, limit=30",
	}

	plants := parse_data(data)
	got := growth_time(plants)
	want := 27

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"id=1, growth=3, limit=10",
		"id=2, growth=5, limit=20",
		"id=3, growth=7, limit=30",
	}

	plants := parse_data(data)
	got := growth_time_leaves_limit(plants)
	want := 42

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
