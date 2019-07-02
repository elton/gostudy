package main

import "testing"

func TestAdd(t *testing.T) {
	a, b, c := 1, 2, 3
	res := add(a, b)
	if res != c {
		t.Fail()
	}
}
