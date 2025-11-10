package quest6

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []byte("ABabACacBCbca")

	got := total_A_pairs(data)
	want := 5

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []byte("ABabACacBCbca")

	got := total_pairs(data)
	want := 11

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []byte("AABCBABCABCabcabcABCCBAACBCa")

	got := total_pairs_in_range(data, 1, 10)
	want := 34

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	data := []byte("AABCBABCABCabcabcABCCBAACBCa")

	got := total_pairs_in_range(data, 2, 10)
	want := 72

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test5(t *testing.T) {
	data := []byte("AABCBABCABCabcabcABCCBAACBCa")

	got := total_pairs_in_range(data, 1000, 1000)
	want := 3442321

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
