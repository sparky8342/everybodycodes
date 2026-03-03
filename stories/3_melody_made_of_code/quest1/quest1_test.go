package quest1

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"2456:rrrrrr ggGgGG bbbbBB",
		"7689:rrRrrr ggGggg bbbBBB",
		"3145:rrRrRr gggGgg bbbbBB",
		"6710:rrrRRr ggGGGg bbBBbB",
	}

	scales := parse_data(data)
	got := green_dominant(scales)
	want := 9166

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"2456:rrrrrr ggGgGG bbbbBB sSsSsS",
		"7689:rrRrrr ggGggg bbbBBB ssSSss",
		"3145:rrRrRr gggGgg bbbbBB sSsSsS",
		"6710:rrrRRr ggGGGg bbBBbB ssSSss",
	}

	scales := parse_data(data)
	got := darkest_shiny(scales)
	want := 2456

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
