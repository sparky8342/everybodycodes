package quest18

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"Plant 1 with thickness 1:",
		"- free branch with thickness 1",
		"",
		"Plant 2 with thickness 1:",
		"- free branch with thickness 1",
		"",
		"Plant 3 with thickness 1:",
		"- free branch with thickness 1",
		"",
		"Plant 4 with thickness 17:",
		"- branch to Plant 1 with thickness 15",
		"- branch to Plant 2 with thickness 3",
		"",
		"Plant 5 with thickness 24:",
		"- branch to Plant 2 with thickness 11",
		"- branch to Plant 3 with thickness 13",
		"",
		"Plant 6 with thickness 15:",
		"- branch to Plant 3 with thickness 14",
		"",
		"Plant 7 with thickness 10:",
		"- branch to Plant 4 with thickness 15",
		"- branch to Plant 5 with thickness 21",
		"- branch to Plant 6 with thickness 34",
	}

	top, _, _ := parse_data(data)
	got := top.energy()
	want := 774

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{

		"Plant 1 with thickness 1:",
		"- free branch with thickness 1",
		"",
		"Plant 2 with thickness 1:",
		"- free branch with thickness 1",
		"",
		"Plant 3 with thickness 1:",
		"- free branch with thickness 1",
		"",
		"Plant 4 with thickness 10:",
		"- branch to Plant 1 with thickness -25",
		"- branch to Plant 2 with thickness 17",
		"- branch to Plant 3 with thickness 12",
		"",
		"Plant 5 with thickness 14:",
		"- branch to Plant 1 with thickness 14",
		"- branch to Plant 2 with thickness -26",
		"- branch to Plant 3 with thickness 15",
		"",
		"Plant 6 with thickness 150:",
		"- branch to Plant 4 with thickness 5",
		"- branch to Plant 5 with thickness 6",
		"",
		"",
		"1 0 1",
		"0 0 1",
		"0 1 1",
	}

	top, free_plants, test_cases := parse_data(data)
	got := run_test_cases(top, free_plants, test_cases)
	want := 324

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
