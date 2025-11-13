package quest8

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []byte("1,5,2,6,8,4,1,7,3")

	nums, _ := parse_data(data)
	got := centre_count(8, nums)
	want := 4

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []byte("1,5,2,6,8,4,1,7,3,5,7,8,2")

	_, lines := parse_data(data)
	got := knots(8, lines)
	want := 21

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []byte("1,5,2,6,8,4,1,7,3,6")

	_, lines := parse_data(data)
	got := best_cut(8, lines)
	want := 7

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
