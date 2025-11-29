package quest20

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		"T#TTT###T##",
		".##TT#TT##.",
		"..T###T#T..",
		"...##TT#...",
		"....T##....",
		".....#.....",
	}

	got, _ := count_pairs(grid)
	want := 7

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	grid := []string{
		"TTTTTTTTTTTTTTTTT",
		".TTTT#T#T#TTTTTT.",
		"..TT#TTTETT#TTT..",
		"...TT#T#TTT#TT...",
		"....TTT#T#TTT....",
		".....TTTTTT#.....",
		"......TT#TT......",
		".......#TT.......",
		"........S........",
	}

	_, paths := count_pairs(grid)
	got := bfs(grid, paths)
	want := 32

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	grid := []string{
		"T####T#TTT##T##T#T#",
		".T#####TTTT##TTT##.",
		"..TTTT#T###TTTT#T..",
		"...T#TTT#ETTTT##...",
		"....#TT##T#T##T....",
		".....#TT####T#.....",
		"......T#TT#T#......",
		".......T#TTT.......",
		"........TT#........",
		".........S.........",
	}

	got := bfs_rotate(grid)
	want := 23

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
