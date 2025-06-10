package quest2

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"ADD id=1 left=[10,A] right=[30,H]",
		"ADD id=2 left=[15,D] right=[25,I]",
		"ADD id=3 left=[12,F] right=[31,J]",
		"ADD id=4 left=[5,B] right=[27,L]",
		"ADD id=5 left=[3,C] right=[28,M]",
		"ADD id=6 left=[20,G] right=[32,K]",
		"ADD id=7 left=[4,E] right=[21,N]",
	}

	got := process_data(data)
	want := "CFGNLK"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"ADD id=1 left=[160,E] right=[175,S]",
		"ADD id=2 left=[140,W] right=[224,D]",
		"ADD id=3 left=[122,U] right=[203,F]",
		"ADD id=4 left=[204,N] right=[114,G]",
		"ADD id=5 left=[136,V] right=[256,H]",
		"ADD id=6 left=[147,G] right=[192,O]",
		"ADD id=7 left=[232,I] right=[154,K]",
		"ADD id=8 left=[118,E] right=[125,Y]",
		"ADD id=9 left=[102,A] right=[210,D]",
		"ADD id=10 left=[183,Q] right=[254,E]",
		"ADD id=11 left=[146,E] right=[148,C]",
		"ADD id=12 left=[173,Y] right=[299,S]",
		"ADD id=13 left=[190,B] right=[277,B]",
		"ADD id=14 left=[124,T] right=[142,N]",
		"ADD id=15 left=[153,R] right=[133,M]",
		"ADD id=16 left=[252,D] right=[276,M]",
		"ADD id=17 left=[258,I] right=[245,P]",
		"ADD id=18 left=[117,O] right=[283,!]",
		"ADD id=19 left=[212,O] right=[127,R]",
		"ADD id=20 left=[278,A] right=[169,C]",
	}

	got := process_data(data)
	want := "EVERYBODYCODES"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
