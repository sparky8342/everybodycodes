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
	valid_names, _ := get_valid_names(names, rules)
	got := valid_names[0]
	want := "Oroneth"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"Xanverax,Khargyth,Nexzeth,Helther,Braerex,Tirgryph,Kharverax",
		"",
		"r > v,e,a,g,y",
		"a > e,v,x,r",
		"e > r,x,v,t",
		"h > a,e,v",
		"g > r,y",
		"y > p,t",
		"i > v,r",
		"K > h",
		"v > e",
		"B > r",
		"t > h",
		"N > e",
		"p > h",
		"H > e",
		"l > t",
		"z > e",
		"X > a",
		"n > v",
		"x > z",
		"T > i",
	}

	names, rules := parse_data(data)
	_, got := get_valid_names(names, rules)
	want := 23

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
