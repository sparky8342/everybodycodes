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

func Test3(t *testing.T) {
	data := []string{
		"CCMM",
		"YYKK",
	}

	got := double_combinations(data)
	want := 18

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	data := []string{
		"CCCMMM",
		"CCMMYY",
		"CCCYKK",
		"CMMYYK",
	}

	got := double_combinations(data)
	want := 343

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test5(t *testing.T) {
	data := []string{
		"CCCCYMMM",
		"CCMMMYYY",
		"CCCYYYMM",
		"MMCCCYYY",
		"YYYMMMCC",
		"CCYYMMKK",
		"MMYYCCKK",
		"YYYCCMMM",
	}

	got := double_combinations(data)
	want := 8000

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
