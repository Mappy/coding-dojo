package main

import "testing"

func Test_should_return_true_if_empty_string(t *testing.T) {
	validate(palindrome(""), "Must return true", t)
}

func validate(b bool, s string, t *testing.T) {
	if (!b) {
		t.Fatalf("Assertion failed: " + s)
	}
}
