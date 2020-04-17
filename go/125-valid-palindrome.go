package main

import "unicode"

func isPalindrome(s string) bool {
	runes := []rune(s)
	left, right := 0, len(s)-1
	for left < right {
		if !valid(runes[left]) {
			left++
			continue
		}
		if !valid(rune(runes[right])) {
			right--
			continue
		}
		if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func valid(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}
