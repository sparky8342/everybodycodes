package quest2

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{"START=[5,0]",
		"A=[0,0]",
		"B=[10,0]",
		"C=[5,10]",
		"MOVES=ABCCBABCA",
	}

	locations, moves := parse_data(data)
	got := len(launch(locations, moves))
	want := 8

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"START=[5,0]",
		"A=[0,0]",
		"B=[10,0]",
		"C=[5,10]",
		"MOVES=ABCCBABCA",
	}

	locations, moves := parse_data(data)
	visited := launch(locations, moves)
	got := add_fireflies(visited)
	want := 25

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"START=[5,0]",
		"A=[0,0]",
		"B=[10,0]",
		"C=[5,10]",
		"MOVES=BABCAABBCABCCCBBABCCCAAACABABCBCBBCAABBABBCACCBAABCBCBBBCBBBBBCCCAACAACB",
	}

	locations, moves := parse_data(data)
	visited := launch(locations, moves)
	got := add_fireflies(visited)
	want := 46

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	data := []string{
		"START=[5,0]",
		"A=[0,0]",
		"B=[10,0]",
		"C=[5,10]",
	}

	locations, _ := parse_data(data)
	visited := launch_random_fill(locations)
	got := add_fireflies(visited)
	want := 42

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test5(t *testing.T) {
	data := []string{
		"START=[0,0]",
		"A=[0,0]",
		"B=[80,15]",
		"C=[5,30]",
	}

	locations, _ := parse_data(data)
	visited := launch_random_fill(locations)
	got := add_fireflies(visited)
	want := 432

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
