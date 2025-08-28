package quest3

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"1: faces=[1,2,3,4,5,6] seed=7",
		"2: faces=[-1,1,-1,1,-1] seed=13",
		"3: faces=[9,8,7,8,9] seed=17",
	}

	dice := parse_data(data)
	got := roll_dice(dice, 10000)
	want := 844

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
