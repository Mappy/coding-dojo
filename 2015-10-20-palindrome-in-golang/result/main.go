package main

func palindrome(s string) bool {
	if (len(s) == 0 || len(s) == 1) {
		return true
	}
	r := []rune(s)

	if (r[0] == r[len(s) - 1]) {
		return palindrome(string(r[1:len(s) - 1]))
	}
	return false
}
