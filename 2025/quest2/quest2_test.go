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
