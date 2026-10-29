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

	got := total_score(data)
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

	got := total_score(data)
	want := 1161463

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"ISO2-3 23",
		"ISO2-61 149",
		"ISO2-978 99",
		"ISO2-5887 5923",
		"ISO2-76988 123749",
		"ISO2-435657 929977",
		"ISO2-3949 870737",
	}

	got := total_score(data)
	want := 50987

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
