package main

import "testing"

func TestLenForMap(t *testing.T) {
	v := map[string]int{"A": 1, "B": 2}
	actual, expected := len(v), 2
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}

func TestLenForString(t *testing.T) {
	v := "one"
	actual, expected := len(v), 3
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}

func TestLenForSlice(t *testing.T) {
	v := []int{5, 0, 4, 1}
	actual, expected := len(v), 4
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}
