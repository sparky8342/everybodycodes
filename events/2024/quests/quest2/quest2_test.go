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
