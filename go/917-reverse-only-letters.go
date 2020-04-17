package main

import "unicode"

func reverseOnlyLetters(S string) string {
	runes := []rune(S)
	left, right := 0, len(runes)-1
	for left < right {
		rl, rr := runes[left], runes[right]
		if !unicode.IsLetter(rl) {
			left++
			continue
		}
		if !unicode.IsLetter(rr) {
			right--
			continue
		}
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}
