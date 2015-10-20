package main

import "testing"

func Test_should_return_true_with_null_or_one_parameter(t *testing.T) {
	validate(palindrome(""), "Must pass with empty parameter", t)

	validate(palindrome("a"), "Must pass with one parameter", t)
}

func Test_should_return_true_with_tow_identical_char(t *testing.T) {
	validate(palindrome("aa"), "Must pass with twice identical value", t)
}

func Test_should_return_true_with_good_palindrome(t *testing.T) {
	validate(palindrome("aba"), "Must pass with good palindrome", t)

	validate(palindrome("abba"), "Must pass with good palindrome", t)

	validate(palindrome("abcba"), "Must pass with good palindrome", t)

	validate(palindrome("abcacba"), "Must pass with good palindrome", t)
}

func validate(b bool, s string, t *testing.T) {
	if (!b) {
		t.Fatalf("Assertion failed: " + s)
	}
}
