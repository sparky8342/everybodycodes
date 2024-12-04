package quest13

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		"#######",
		"#6769##",
		"S50505E",
		"#97434#",
		"#######",
	}

	got := find_path(grid, 'S', 'E')
	want := 28

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	grid := []string{
		"SSSSSSSSSSS",
		"S674345621S",
		"S###6#4#18S",
		"S53#6#4532S",
		"S5450E0485S",
		"S##7154532S",
		"S2##314#18S",
		"S971595#34S",
		"SSSSSSSSSSS",
	}

	got := find_path(grid, 'E', 'S')
	want := 14

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
