package main

import "testing"

func TestA(t *testing.T) {
	if add(3, 2) != 5 {
		t.Fatalf("add 3, 2 should be 5, got %d", add(3, 2))
	}
}
