package quest16

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"1,2,3",
		"",
		"^_^ -.- ^,-",
		">.- ^_^ >.<",
		"-_- -.- >.<",
		"    -.^ ^_^",
		"    >.>",
	}

	machine := parse_data(data)
	machine.spin(100)

	got := machine.display()
	want := ">.- -.- ^,-"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"1,2,3",
		"",
		"^_^ -.- ^,-",
		">.- ^_^ >.<",
		"-_- -.- >.<",
		"    -.^ ^_^",
		"    >.>",
	}

	machine := parse_data(data)

	got := machine.spin_with_score(10)
	want := 15
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	machine.reset()
	got = machine.spin_with_score(100)
	want = 138
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	machine.reset()
	got = machine.spin_with_score(10000000000)
	want = 13833333333
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	machine.reset()
	got = machine.spin_with_score(202420242024)
	want = 280014668134
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

}

func Test3(t *testing.T) {
	data := []string{
		"1,2,3",
		"",
		"^_^ -.- ^,-",
		">.- ^_^ >.<",
		"-_- -.- ^.^",
		"    -.^ >.<",
		"    >.>",
	}

	machine := parse_data(data)

	got_min, got_max := machine.minmax(100)
	want_min := 50
	want_max := 246

	if got_min != want_min {
		t.Errorf("got %d, wanted %d", got_min, want_min)
	}

	if got_max != want_max {
		t.Errorf("got %d, wanted %d", got_max, want_max)
	}
}
