package quest3

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"ISO-3210 7654",
		"ISO-8742 2468",
		"ISO-1111 1111",
	}

	got := score(data)
	want := 1478

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"ISO-3210 7654",
		"ISO-8742 2468",
		"ISO-1 1111",
		"ISO-13 777",
		"ISO-34 140856",
		"ISO-789 2465112",
		"ISO-2341 945",
	}

	got := score(data)
	want := 1161463

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
