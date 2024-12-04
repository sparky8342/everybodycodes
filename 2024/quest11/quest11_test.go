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
	got := days(rules, "A", 4)
	var want uint64 = 8

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"A:B,C",
		"B:C,A,A",
		"C:A",
	}

	rules := parse_data(data)
	got := all_starts(rules)
	var want uint64 = 268815

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
