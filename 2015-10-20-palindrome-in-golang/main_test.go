package main

import "testing"

func Test_should_return_true_if_empty_string(t *testing.T) {
	assertTrue(palindrome(""), t)
}

func Test_should_return_true_with_one_character(t *testing.T) {
	assertTrue(palindrome("a"), t)
}

func Test_should_return_false_with_2_diff_characters(t *testing.T) {
	assertFalse(palindrome("ab"), t)
}

func Test_should_return_true_with_2_equal_characters(t *testing.T) {
	assertTrue(palindrome("aa"), t)
}

func Test_should_return_true_with_3_equal_characters(t *testing.T) {
	assertTrue(palindrome("aba"), t)
}

func Test_should_return_True_with_1space_1characters(t *testing.T) {
	assertTrue(palindrome(" a"), t)
}

func Test_should_return_true_with_4_different_characters(t *testing.T) {
	assertFalse(palindrome("abca"), t)
}

func Test_should_return_true_with_many_different_characters(t *testing.T) {
	assertTrue(palindrome("engage le jeu que je le gagne"), t)
}

func assertFalse(b bool, t *testing.T) {
	if (b) {
		t.Fatalf("Assertion failed")
	}
}

func assertTrue(b bool, t *testing.T) {
	if (!b) {
		t.Fatalf("Assertion failed: must return true")
	}
}

func validate(b bool, s string, t *testing.T) {
	if (!b) {
		t.Fatalf("Assertion failed: " + s)
	}
}
