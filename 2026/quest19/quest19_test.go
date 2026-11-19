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

	got := combine_colours(data)
	want := 144

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
