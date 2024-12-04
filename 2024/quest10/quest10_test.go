package quest10

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"**PCBS**",
		"**RLNW**",
		"BV....PT",
		"CR....HZ",
		"FL....JW",
		"SG....MN",
		"**FTZV**",
		"**GMJH**",
	}

	got := solve_grid(data)
	want := "PTBVRCZHFLJWGMNS"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test2(t *testing.T) {
	word := "PTBVRCZHFLJWGMNS"

	got := power(word)
	want := 1851

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"**XFZB**DCST**",
		"**LWQK**GQJH**",
		"?G....WL....DQ",
		"BS....H?....CN",
		"P?....KJ....TV",
		"NM....Z?....SG",
		"**NSHM**VKWZ**",
		"**PJGV**XFNL**",
		"WQ....?L....YS",
		"FX....DJ....HV",
		"?Y....WM....?J",
		"TJ....YK....LP",
		"**XRTK**BMSP**",
		"**DWZN**GCJV**",
	}

	got := solve_multi_grid(data)
	want := 3889

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
