package quest5

import (
	"testing"
)

func Test1(t *testing.T) {
	data := []byte("58:5,3,7,8,9,10,4,5,7,8,8")

	nums := parse_data(data)
	got := quality(nums)
	want := "581078"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
