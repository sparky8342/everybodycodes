package quest6

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"RR:A,B,C",
		"A:D,E",
		"B:F,@",
		"C:G,H",
		"D:@",
		"E:@",
		"F:@",
		"G:@",
		"H:@",
	}

	root := parse_data(data)
	got := find_path(root)

	want := "RRB@"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
