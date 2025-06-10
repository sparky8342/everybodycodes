package quest1

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"A=4 B=4 C=6 X=3 Y=4 Z=5 M=11",
		"A=8 B=4 C=7 X=8 Y=4 Z=6 M=12",
		"A=2 B=8 C=6 X=2 Y=4 Z=5 M=13",
		"A=5 B=9 C=6 X=8 Y=6 Z=8 M=14",
		"A=5 B=9 C=7 X=6 Y=6 Z=8 M=15",
		"A=8 B=8 C=8 X=6 Y=9 Z=6 M=16",
	}

	param_lines := parse_data(data)

	got := highest_line(param_lines, eni)
	want := 11611972920

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"A=4 B=4 C=6 X=3 Y=14 Z=15 M=11",
		"A=8 B=4 C=7 X=8 Y=14 Z=16 M=12",
		"A=2 B=8 C=6 X=2 Y=14 Z=15 M=13",
		"A=5 B=9 C=6 X=8 Y=16 Z=18 M=14",
		"A=5 B=9 C=7 X=6 Y=16 Z=18 M=15",
		"A=8 B=8 C=8 X=6 Y=19 Z=16 M=16",
	}

	param_lines := parse_data(data)

	got := highest_line(param_lines, eni2)
	want := 11051340

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"A=3657 B=3583 C=9716 X=903056852 Y=9283895500 Z=85920867478 M=188",
		"A=6061 B=4425 C=5082 X=731145782 Y=1550090416 Z=87586428967 M=107",
		"A=7818 B=5395 C=9975 X=122388873 Y=4093041057 Z=58606045432 M=102",
		"A=7681 B=9603 C=5681 X=716116871 Y=6421884967 Z=66298999264 M=196",
		"A=7334 B=9016 C=8524 X=297284338 Y=1565962337 Z=86750102612 M=145",
	}

	param_lines := parse_data(data)

	got := highest_line(param_lines, eni2)
	want := 1507702060886

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	data := []string{
		"A=4 B=4 C=6 X=3000 Y=14000 Z=15000 M=110",
		"A=8 B=4 C=7 X=8000 Y=14000 Z=16000 M=120",
		"A=2 B=8 C=6 X=2000 Y=14000 Z=15000 M=130",
		"A=5 B=9 C=6 X=8000 Y=16000 Z=18000 M=140",
		"A=5 B=9 C=7 X=6000 Y=16000 Z=18000 M=150",
		"A=8 B=8 C=8 X=6000 Y=19000 Z=16000 M=160",
	}

	param_lines := parse_data(data)

	got := highest_line(param_lines, eni3)
	want := 3279640

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test5(t *testing.T) {
	data := []string{
		"A=3657 B=3583 C=9716 X=903056852 Y=9283895500 Z=85920867478 M=188",
		"A=6061 B=4425 C=5082 X=731145782 Y=1550090416 Z=87586428967 M=107",
		"A=7818 B=5395 C=9975 X=122388873 Y=4093041057 Z=58606045432 M=102",
		"A=7681 B=9603 C=5681 X=716116871 Y=6421884967 Z=66298999264 M=196",
		"A=7334 B=9016 C=8524 X=297284338 Y=1565962337 Z=86750102612 M=145",
	}

	param_lines := parse_data(data)

	got := highest_line(param_lines, eni3)
	want := 7276515438396

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
