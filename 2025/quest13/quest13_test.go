package quest13

import (
	"testing"
)

func Test1(t *testing.T) {
	nums := []int{72, 58, 47, 61, 67}

	got := craft_lock(nums, 2025)
	want := 67

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"10-15",
		"12-13",
		"20-21",
		"19-23",
		"30-37",
	}

	ranges := parse_data(data)
	got := craft_lock_ranges(ranges, 20252025)
	want := 30

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
