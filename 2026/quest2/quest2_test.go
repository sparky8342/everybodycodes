package quest2

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []string{
		"4<3 2<5 3>1 1>2 1<5 4<1 3>2 4<2 3>5 4<5",
		"4>1 2>1 3>2 1<5 4>5 2<5 3>1 4>3 2<4 5<3",
		"5>1 5>2 4>1 2<4 3>2 3<4 1>2 3>5 1<3 4>5",
	}

	got := arrange_groups(data)
	want := 122157

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
