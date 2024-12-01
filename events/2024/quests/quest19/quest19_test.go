package quest19

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"LR",
		"",
		">-IN-",
		"-----",
		"W---<",
	}

	got := decode(data)
	want := "WIN"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
