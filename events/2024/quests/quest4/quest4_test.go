package quest4

import (
	"testing"
)

func Test1(t *testing.T) {
	nums := []int{3, 4, 7, 8}

	got := hammer(nums)
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
