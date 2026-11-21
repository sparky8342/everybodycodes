package quest20

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"5 - 4 = 9 - 8",
		"2 + 1 = 3 + 0",
		"6 + 7 = 4 + 9",
		"1 + 2 = 3 + 4",
	}

	got := count_segments(data, 10)
	want := "75"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"10 + 2A = 13 + 27",
		"98 - 54 = 76 - 32",
		"10 + 12 = 30 - 9",
		"20 - 10 = 5 + 5",
	}

	got := count_segments(data, 11)
	want := "101"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"1 + 1 = 5 - 1",
		"1 + 1 = 3 - 3",
		"0 - 1 = 1 + 7",
		"1 + 1 = 1 - 6",
		"1 + 1 = 2 - 1",
		"1 + 1 = 5 + 7",
		"1 + 4 = 9 - 9",
		"1 + 1 = 2 + 9",
		"1 + 1 = 1 + 2",
	}

	got := count_segments_balance(data, 11)
	want := "125"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test4(t *testing.T) {
	data := []string{
		"4469 - 5934 = 1631 - 2AA7",
		"8471 - 2276 = 4510 + 2736",
		"A767 - 6104 = 1825 + 2538",
		"9961 + 1140 = 5963 + 2974",
		"9188 - 406A = 3598 + 1533",
		"96A6 - 1606 = 4676 - 3524",
		"9114 - 15A2 = 5258 + 2371",
		"7023 - 2209 = 1433 + 3436",
		"921A - 2696 = 2A91 + 3688",
		"8492 - 2055 = 4186 + 3260",
		"1234 + 4567 = 8901 + 2345",
	}

	got := count_segments_balance(data, 11)
	want := "693"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
