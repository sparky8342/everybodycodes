package quest2

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		".......",
		".......",
		".......",
		".#.@...",
		".......",
		".......",
		".......",
	}

	start, bone := parse_data(data)
	got := steps_to_bone(start, bone)
	want := 12

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
