package quest5

import (
	"testing"
)

func Test1(t *testing.T) {
	data := "58:5,3,7,8,9,10,4,5,7,8,8"

	nums := parse_line(data)
	got := quality(nums)
	want := 581078

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"1:2,4,1,1,8,2,7,9,8,6",
		"2:7,9,9,3,8,3,8,8,6,8",
		"3:4,7,6,9,1,8,3,7,2,2",
		"4:6,4,2,1,7,4,5,5,5,8",
		"5:2,9,3,8,3,9,5,2,1,4",
		"6:2,4,9,6,7,4,1,7,6,8",
		"7:2,3,7,6,2,2,4,1,4,2",
		"8:5,1,5,6,8,3,1,8,3,9",
		"9:5,7,7,3,7,2,3,8,6,7",
		"10:4,1,9,3,8,5,4,3,5,5",
	}

	got := compare_quality(data)
	want := 77053

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
