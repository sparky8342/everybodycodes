package quest12

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := []string{
		"989611",
		"857782",
		"746543",
		"766789",
	}

	got := shoot_barrels(grid, 1)
	want := 16

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	grid := []string{
		"9589233445",
		"9679121695",
		"8469121876",
		"8352919876",
		"7342914327",
		"7234193437",
		"6789193538",
		"6781219648",
		"5691219769",
		"5443329859",
	}

	got := shoot_barrels(grid, 2)
	want := 58

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	grid := []string{
		"5411",
		"3362",
		"5235",
		"3112",
	}

	got := shoot_barrels(grid, 3)
	want := 14

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	grid := []string{
		"41951111131882511179",
		"32112222211508122215",
		"31223333322105122219",
		"31234444432147511128",
		"91223333322176021892",
		"60112222211166431583",
		"04661111166111111746",
		"01111119042122222177",
		"41222108881233333219",
		"71222127839122222196",
		"56111026279711111507",
	}

	got := shoot_barrels(grid, 3)
	want := 133

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
