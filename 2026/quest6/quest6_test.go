package quest6

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"BRDS:FLWR,SLTT,YSTS",
		"FLWR:SLTT,YSTS",
		"SLTT:YSTS,MLKK,BTTR",
		"NULL:BRDS,FLWR,SLTT",
		"YSTS:BRDS,FLWR",
		"MLKK:SPCS,NULL",
		"HNYN:BRDS,FLWR,SLTT,YSTS",
	}

	ingredients := parse_data(data)
	got := find_no_requirements(ingredients, "NULL")
	want := 7

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"BRDS:FLWR,SLTT,YSTS",
		"FLWR:SLTT,YSTS",
		"SLTT:YSTS,MLKK,BTTR",
		"NULL:BRDS,FLWR,SLTT",
		"YSTS:BRDS,FLWR",
		"MLKK:SPCS,NULL",
		"HNYN:BRDS,FLWR,SLTT,YSTS",
	}

	ingredients := parse_data(data)
	got := find_no_products(ingredients, "NULL")
	want := 6

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
