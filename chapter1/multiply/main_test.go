package main

import "testing"

func TestMultiply(t *testing.T) {
	a := 5
	b := 6
	expected := 30

	res := multiply(a, b)
	if res != expected {
		t.Errorf("Out multiply function doesn't work, %d*%d isn't %d\n", a, b, res)
	}
}
