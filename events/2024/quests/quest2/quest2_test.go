package quest2

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"WORDS:THE,OWE,MES,ROD,HER",
		"",
		"AWAKEN THE POWER ADORNED WITH THE FLAMES BRIGHT IRE",
	}

	words, phrase := parse_data(data)

	got := contains(words, phrase)
	want := 4

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	phrase = "THE FLAME SHIELDED THE HEART OF THE KINGS"
	got = contains(words, phrase)
	want = 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	phrase = "POWE PO WER P OWE R"
	got = contains(words, phrase)
	want = 2

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	phrase = "THERE IS THE END"
	got = contains(words, phrase)
	want = 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"WORDS:THE,OWE,MES,ROD,HER,QAQ",
		"",
		"AWAKEN THE POWE ADORNED WITH THE FLAMES BRIGHT IRE",
		"THE FLAME SHIELDED THE HEART OF THE KINGS",
		"POWE PO WER P OWE R",
		"THERE IS THE END",
		"QAQAQ",
	}

	words, phrase := parse_data(data)

	got := runic_symbols(words, phrase)
	want := 42

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"WORDS:THE,OWE,MES,ROD,RODEO",
		"",
		"HELWORLT",
		"ENIGWDXL",
		"TRODEOAL",
	}

	words, phrase := parse_data(data)

	got := scales(words, phrase)
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
