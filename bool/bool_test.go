package main

import "testing"

func TestCheck(t *testing.T) {
	b := check()

	if b == false {
		t.Errorf("Expected b to be true but got false")
	}
}
