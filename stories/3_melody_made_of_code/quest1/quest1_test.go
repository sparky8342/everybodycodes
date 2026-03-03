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

	got := green_dominant(data)
	want := 9166

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
