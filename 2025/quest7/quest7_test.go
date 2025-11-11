package quest7

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"Oronris,Urakris,Oroneth,Uraketh",
		"",
		"r > a,i,o",
		"i > p,w",
		"n > e,r",
		"o > n,m",
		"k > f,r",
		"a > k",
		"U > r",
		"e > t",
		"O > r",
		"t > h",
	}

	names, rules := parse_data(data)
	got := get_valid_name(names, rules)
	want := "Oroneth"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
