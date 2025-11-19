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
