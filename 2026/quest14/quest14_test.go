package quest14

import (
	"testing"
)

func Test1(t *testing.T) {
	window := []string{
		"AAAEEEFFFF",
		"AADDDDDGFF",
		"BACDDIGGFF",
		"BABBIIIGGH",
		"BBBBIIIIHH",
	}

	got := bindings_needed(window)
	want := 67

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
