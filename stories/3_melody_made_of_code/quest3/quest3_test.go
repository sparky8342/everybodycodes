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

	root := parse_data(data, 1)
	got := read_tree(root)
	want := 43

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"id=1, plug=RED TRIANGLE, leftSocket=RED TRIANGLE, rightSocket=RED TRIANGLE, data=?",
		"id=2, plug=GREEN TRIANGLE, leftSocket=BLUE CIRCLE, rightSocket=GREEN CIRCLE, data=?",
		"id=3, plug=BLUE PENTAGON, leftSocket=BLUE CIRCLE, rightSocket=GREEN CIRCLE, data=?",
		"id=4, plug=RED TRIANGLE, leftSocket=BLUE PENTAGON, rightSocket=GREEN PENTAGON, data=?",
		"id=5, plug=RED PENTAGON, leftSocket=GREEN CIRCLE, rightSocket=GREEN CIRCLE, data=?",
	}

	root := parse_data(data, 2)
	got := read_tree(root)
	want := 50

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"id=1, plug=RED TRIANGLE, leftSocket=RED TRIANGLE, rightSocket=RED TRIANGLE, data=?",
		"id=2, plug=GREEN TRIANGLE, leftSocket=BLUE CIRCLE, rightSocket=GREEN CIRCLE, data=?",
		"id=3, plug=BLUE PENTAGON, leftSocket=BLUE CIRCLE, rightSocket=GREEN CIRCLE, data=?",
		"id=4, plug=RED TRIANGLE, leftSocket=BLUE PENTAGON, rightSocket=GREEN PENTAGON, data=?",
		"id=5, plug=RED PENTAGON, leftSocket=GREEN CIRCLE, rightSocket=GREEN CIRCLE, data=?",
	}

	root := parse_data(data, 3)
	got := read_tree(root)
	want := 38

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	data := []string{
		"id=1, plug=RED TRIANGLE, leftSocket=BLUE TRIANGLE, rightSocket=GREEN TRIANGLE, data=?",
		"id=2, plug=GREEN TRIANGLE, leftSocket=BLUE CIRCLE, rightSocket=GREEN CIRCLE, data=?",
		"id=3, plug=BLUE PENTAGON, leftSocket=BLUE CIRCLE, rightSocket=GREEN CIRCLE, data=?",
		"id=4, plug=RED TRIANGLE, leftSocket=BLUE PENTAGON, rightSocket=GREEN PENTAGON, data=?",
		"id=5, plug=BLUE TRIANGLE, leftSocket=GREEN CIRCLE, rightSocket=RED CIRCLE, data=?",
		"id=6, plug=BLUE TRIANGLE, leftSocket=GREEN CIRCLE, rightSocket=RED CIRCLE, data=?",
	}

	root := parse_data(data, 3)
	got := read_tree(root)
	want := 60

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
