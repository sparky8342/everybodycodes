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
