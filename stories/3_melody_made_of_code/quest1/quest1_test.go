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

func Test3(t *testing.T) {
	data := []string{
		"15437:rRrrRR gGGGGG BBBBBB sSSSSS",
		"94682:RrRrrR gGGggG bBBBBB ssSSSs",
		"56513:RRRrrr ggGGgG bbbBbb ssSsSS",
		"76346:rRRrrR GGgggg bbbBBB ssssSs",
		"87569:rrRRrR gGGGGg BbbbbB SssSss",
		"44191:rrrrrr gGgGGG bBBbbB sSssSS",
		"49176:rRRrRr GggggG BbBbbb sSSssS",
		"85071:RRrrrr GgGGgg BBbbbb SSsSss",
		"44303:rRRrrR gGggGg bBbBBB SsSSSs",
		"94978:rrRrRR ggGggG BBbBBb SSSSSS",
		"26325:rrRRrr gGGGgg BBbBbb SssssS",
		"43463:rrrrRR gGgGgg bBBbBB sSssSs",
		"15059:RRrrrR GGgggG bbBBbb sSSsSS",
		"85004:RRRrrR GgGgGG bbbBBB sSssss",
		"56121:RRrRrr gGgGgg BbbbBB sSsSSs",
		"80219:rRRrRR GGGggg BBbbbb SssSSs",
	}

	scales := parse_data(data)
	got := largest_group(scales)
	want := 292320

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
