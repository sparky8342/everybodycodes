package quest11

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"A:B,C",
		"B:C,A",
		"C:A",
	}

	rules := parse_data(data)
	got := days(rules, 0, 4)
	want := 8

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
