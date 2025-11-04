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

	names, moves := parse_data(data)
	got := get_name(names, moves)
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

	names, moves := parse_data(data)
	got := get_name_circular(names, moves)
	want := "Elarzris"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"Vyrdax,Drakzyph,Fyrryn,Elarzris",
		"",
		"R3,L2,R3,L3",
	}

	names, moves := parse_data(data)
	got := get_name_with_swaps(names, moves)
	want := "Drakzyph"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
