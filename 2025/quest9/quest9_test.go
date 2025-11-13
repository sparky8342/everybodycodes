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
	got := similarity_part1(sequences)
	want := 414

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"1:GCAGGCGAGTATGATACCCGGCTAGCCACCCC",
		"2:TCTCGCGAGGATATTACTGGGCCAGACCCCCC",
		"3:GGTGGAACATTCGAAAGTTGCATAGGGTGGTG",
		"4:GCTCGCGAGTATATTACCGAACCAGCCCCTCA",
		"5:GCAGCTTAGTATGACCGCCAAATCGCGACTCA",
		"6:AGTGGAACCTTGGATAGTCTCATATAGCGGCA",
		"7:GGCGTAATAATCGGATGCTGCAGAGGCTGCTG",
	}

	sequences := parse_data(data)
	got := similarity_sum(sequences)
	want := 1245

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
