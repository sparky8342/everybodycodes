package quest17

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"*...*",
		"..*..",
		".....",
		".....",
		"*.*..",
	}

	stars := parse_data(data)
	got := minimum_spanning_tree(stars)
	want := 16

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
