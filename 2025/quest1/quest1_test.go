package quest1

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"Vyrdax,Drakzyph,Fyrryn,Elarzris",
		"",
		"R3,L2,R3,L1",
	}

	got := get_name(data, false)
	want := "Fyrryn"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"Vyrdax,Drakzyph,Fyrryn,Elarzris",
		"",
		"R3,L2,R3,L1",
	}

	got := get_name(data, true)
	want := "Elarzris"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
