package quest3

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"id=1, plug=BLUE HEXAGON, leftSocket=GREEN CIRCLE, rightSocket=BLUE PENTAGON, data=?",
		"id=2, plug=GREEN CIRCLE, leftSocket=BLUE HEXAGON, rightSocket=BLUE CIRCLE, data=?",
		"id=3, plug=BLUE PENTAGON, leftSocket=BLUE CIRCLE, rightSocket=BLUE CIRCLE, data=?",
		"id=4, plug=BLUE CIRCLE, leftSocket=RED HEXAGON, rightSocket=BLUE HEXAGON, data=?",
		"id=5, plug=RED HEXAGON, leftSocket=GREEN CIRCLE, rightSocket=RED HEXAGON, data=?",
	}

	root := parse_data(data)
	got := read_tree(root)
	want := 43

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
