package testghactions

import "testing"

func TestAdd2(t *testing.T) {
	if got, want :=  add2(2), 4; got!=want {
		t.Fatalf("want %d, got %d", want, got)
	}
}
