package quest19

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"CCCCC",
		"MMMMM",
		"YYYYY",
		"KKKKK",
		"CCMYY",
		"",
		"CMYKC",
		"CMYKC",
		"CMYKC",
		"CMYKC",
		"CMYKC",
	}

	grid1, grid2 := parse_data(data)
	got := combine_colours(grid1, grid2)
	want := 144

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"CCCCC",
		"MMMMM",
		"YYYYY",
		"KKKKK",
		"CCMYY",
		"",
		"14231",
		"14231",
		"14231",
		"14231",
		"14231",
	}

	grid1, grid2 := parse_data(data)
	got := combinations(grid1, grid2)
	want := 180

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
