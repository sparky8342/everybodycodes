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
	got := similarity(sequences[2], sequences[0], sequences[1])
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
	got, _ := similarity_sum(sequences)
	want := 1245

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
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
	_, got := similarity_sum(sequences)
	want := 12

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	data := []string{
		"1:GCAGGCGAGTATGATACCCGGCTAGCCACCCC",
		"2:TCTCGCGAGGATATTACTGGGCCAGACCCCCC",
		"3:GGTGGAACATTCGAAAGTTGCATAGGGTGGTG",
		"4:GCTCGCGAGTATATTACCGAACCAGCCCCTCA",
		"5:GCAGCTTAGTATGACCGCCAAATCGCGACTCA",
		"6:AGTGGAACCTTGGATAGTCTCATATAGCGGCA",
		"7:GGCGTAATAATCGGATGCTGCAGAGGCTGCTG",
		"8:GGCGTAAAGTATGGATGCTGGCTAGGCACCCG",
	}

	sequences := parse_data(data)
	_, got := similarity_sum(sequences)
	want := 36

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
