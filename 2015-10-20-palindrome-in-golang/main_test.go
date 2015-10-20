package main

import "testing"

func Test_should_return_true(t *testing.T) {
	validate(true, "Must return true", t)
}

func validate(b bool, s string, t *testing.T) {
	if (!b) {
		t.Fatalf("Assertion failed: " + s)
	}
}
