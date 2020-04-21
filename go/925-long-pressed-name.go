package main

func isLongPressedName(name string, typed string) bool {
	nameRunes, typedRunes := []rune(name), []rune(typed)
	i, j := 0, 0
	for j < len(typedRunes) {
		if nameRunes[i] == typedRunes[j] {
			i++
			j++
			continue
		}
		if j > 0 && typedRunes[j] == typedRunes[j-1] {
			j++
			continue
		}
		return false
	}
	return i == len(nameRunes) && j == len(typedRunes)
}
