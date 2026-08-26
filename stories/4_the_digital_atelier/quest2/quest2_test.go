package quest2

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{"START=[5,0]",
		"A=[0,0]",
		"B=[10,0]",
		"C=[5,10]",
		"MOVES=ABCCBABCA",
	}

	locations, moves := parse_data(data)
	got := launch(locations, moves)
	want := 8

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
