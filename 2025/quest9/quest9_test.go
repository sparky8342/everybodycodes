package quest9

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"1:CAAGCGCTAAGTTCGCTGGATGTGTGCCCGCG",
		"2:CTTGAATTGGGCCGTTTACCTGGTTTAACCAT",
		"3:CTAGCGCTGAGCTGGCTGCCTGGTTGACCGCG",
	}

	sequences := parse_data(data)
	got := similarity(sequences)
	want := 414

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
