package quest8

import (
	"testing"
)

func Test1(t *testing.T) {
	nums := []int{1, 5, 2, 6, 8, 4, 1, 7, 3}

	got := centre_count(8, nums)
	want := 4

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
