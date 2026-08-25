package quest1

import (
	"testing"
)

func Test1(t *testing.T) {
	sequences := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
	}

	got := process_sequences(sequences)
	want := 66

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	sequences := [][]int{
		[]int{1, 1, 1, 1, 1},
		[]int{5, 1, 2, 3, 4, 5, 1, 2, 3, 4},
		[]int{2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1},
		[]int{5, 1, 2, 1, 2, 7, 1, 2, 1, 2, 7, 1, 2, 1, 2},
	}

	got := process_sequences(sequences)
	want := 34

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
