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
