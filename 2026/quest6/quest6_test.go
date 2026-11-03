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

func Test3(t *testing.T) {
	data := []string{
		"MEAL:BRDS,CAKE,DOUG",
		"BRDS:DOUG,FLWR",
		"CAKE:DOUG,FLWR,NULLv2,NULLv3",
		"DOUG:FLWR,YSTS,SLTT,NULLv1",
		"FLWR:BRDS,DOUG",
		"YSTS:SPCS",
		"SOUP:BRDS,CAKE,DOUG,FLWR",
	}

	ingredients := parse_data(data)
	got := find_products_with_multiple(ingredients, []string{"NULLv1", "NULLv2", "NULLv3"})
	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
