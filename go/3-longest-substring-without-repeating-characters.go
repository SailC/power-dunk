package main

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	var lsLen int
	runeMap := map[byte]int{} // map from rune to index of the string
	i := 0
	for j := 0; j < n; j++ {
		rj := s[j]
		if index, ok := runeMap[rj]; ok {
			i = index + 1
			continue
		}
		spanLen := j - i + 1
		if spanLen > lsLen {
			lsLen = spanLen
		}
		runeMap[rj] = j
	}
	return lsLen
}
