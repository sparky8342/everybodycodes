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

func Test2(t *testing.T) {
	data := []byte("1,2,2,2,2,3,1,2,3,3,1,3,1,2,3,2,1,4,1,3,2,2,1,3,2,2")

	nums := parse_data(data)
	_, got := find_pattern(nums)
	want := 270

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []byte("1,2,2,2,2,3,1,2,3,3,1,3,1,2,3,2,1,4,1,3,2,2,1,3,2,2")

	nums := parse_data(data)
	got := find_wall_length(nums, 202520252025000)
	want := 94439495762954

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
