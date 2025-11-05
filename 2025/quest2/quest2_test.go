package quest2

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []byte("A=[25,9]")

	A := parse_data(data)
	got := calculate_part1(A)
	want := "[357,862]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test2(t *testing.T) {
	data := []byte("A=[35300,-64910]")

	A := parse_data(data)
	got := calculate_part2(A)
	want := 4076

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
