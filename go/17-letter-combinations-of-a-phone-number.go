package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	letterMap := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var combos []string
	var combo []rune

	digitRunes := []rune(digits)

	var dfs func(int)
	dfs = func(start int) {
		if start == len(digitRunes) {
			combos = append(combos, string(combo))
			return
		}
		digit := digitRunes[start]
		for _, letter := range letterMap[digit] {
			combo = append(combo, letter)
			dfs(start + 1)
			combo = combo[:len(combo)-1]
		}
	}

	dfs(0)

	return combos
}
