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

func Test3(t *testing.T) {
	data := []string{
		"Xaryt",
		"",
		"X > a,o",
		"a > r,t",
		"r > y,e,a",
		"h > a,e,v",
		"t > h",
		"v > e",
		"y > p,t",
	}

	prefixes, rules := parse_data(data)
	valid_prefixes, _ := get_valid_names(prefixes, rules)
	got := possible_names(valid_prefixes, rules)
	want := 25

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	data := []string{
		"Khara,Xaryt,Noxer,Kharax",
		"",
		"r > v,e,a,g,y",
		"a > e,v,x,r,g",
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

	prefixes, rules := parse_data(data)
	got := possible_names(prefixes, rules)
	want := 1154

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
