package quest5

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"2 3 4 5",
		"3 4 5 2",
		"4 5 2 3",
		"5 2 3 4",
	}

	nums := parse_data(data)
	got := steps(nums, 10)

	want := 2323

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"2 3 4 5",
		"6 7 8 9",
	}

	nums := parse_data(data)
	got := find_repeat(nums, 2024)

	want := 50877075

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"2 3 4 5",
		"6 7 8 9",
	}

	nums := parse_data(data)
	got := highest_top_number(nums)

	want := 6584

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
