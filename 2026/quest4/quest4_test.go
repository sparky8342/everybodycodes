package quest4

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"id=1 quest=1 part=1 time=187",
		"id=3 quest=1 part=1 time=413",
		"id=5 quest=1 part=3 time=744",
		"id=2 quest=1 part=3 time=745",
		"id=4 quest=1 part=2 time=558",
		"id=1 quest=1 part=2 time=299",
		"id=1 quest=1 part=3 time=362",
		"id=4 quest=1 part=3 time=1268",
		"id=4 quest=1 part=1 time=457",
		"id=5 quest=1 part=1 time=257",
		"id=2 quest=1 part=1 time=192",
		"id=3 quest=1 part=3 time=1168",
		"id=5 quest=1 part=2 time=346",
		"id=3 quest=1 part=2 time=455",
		"id=2 quest=1 part=2 time=562",
	}

	warriors := parse_data(data)

	got := order_warriors(warriors)
	want := 27598

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"id=2 quest=1 part=1 time=100",
		"id=4 quest=2 part=1 time=150",
		"id=1 quest=2 part=1 time=140",
		"id=1 quest=1 part=2 time=180",
		"id=2 quest=2 part=1 time=130",
		"id=4 quest=1 part=1 time=110",
		"id=1 quest=2 part=2 time=210",
		"id=2 quest=2 part=2 time=240",
		"id=1 quest=2 part=3 time=330",
		"id=2 quest=1 part=2 time=170",
		"id=3 quest=1 part=1 time=90",
		"id=4 quest=1 part=2 time=175",
		"id=3 quest=2 part=1 time=120",
		"id=1 quest=1 part=1 time=120",
		"id=3 quest=2 part=2 time=220",
		"id=2 quest=2 part=3 time=350",
		"id=3 quest=1 part=2 time=160",
	}

	warriors := parse_data(data)

	got := order_warriors(warriors)
	want := 6470

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
