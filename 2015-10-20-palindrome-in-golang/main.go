package main

import (
	"fmt"
	"strings"
)

func main() {
   fmt.Printf("2015-10-20-palindrome-in-golang\n")
}

func palindrome(s string) bool {
	whitoutSpaces := strings.Replace(s," ", "", -1)

	if len(whitoutSpaces) < 2 {
		return true
	} else {
		r := []rune(whitoutSpaces)
		return r[0] == r[len(r) - 1] && palindrome(string(r[1:len(r)-1]))
	}
}
