package quest14

import (
	"testing"
)

func Test1(t *testing.T) {
	window := []string{
		"AAAEEEFFFF",
		"AADDDDDGFF",
		"BACDDIGGFF",
		"BABBIIIGGH",
		"BBBBIIIIHH",
	}

	got := bindings_needed(window)
	want := 67

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	window := []string{
		"AAAEEEFFFF",
		"AADDDDDGFF",
		"BACDDIGGFF",
		"BABBIIIGGH",
		"BBBBIIIIHH",
	}

	got := minimum_bindings_needed(window)
	want := 12

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	window := []string{
		"IIIIIILOHHHHHHHHGGGGGQSRRRRKKK",
		"IIIIIILOOHHHHHHGGXGGGQSRRRRKKK",
		"PPPPILLLOOOODDHBGGGGGQSSRRRKKK",
		"PPZPLLLLFFOODDBBBBGGQQQSSRKKKK",
		"PPPPLLLLFFFONDDDBBMGGQQSSKKJJJ",
		"PPPCCCCCCFFONNDDMMMGQQQSSSKJJJ",
		"PCCCCCCCCFFNNDDDDMAAQQQSSSKKJJ",
		"PPCCECCCFFNNNNDADMAATQTTSSSJJJ",
		"PPEEEEECCFFNNAAAMMATTTTTSTSSSJ",
		"PPEEEEEEEFFNNNAAAAAATTTTTTSSJJ",
	}

	got := minimum_bindings_needed(window)
	want := 63

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
