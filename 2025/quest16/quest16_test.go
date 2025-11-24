package quest16

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []byte("1,2,3,5,9")

	nums := parse_data(data)
	got := build_wall(nums, 90)
	want := 193

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
