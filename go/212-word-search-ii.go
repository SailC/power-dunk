type Trie struct {
    isEnd bool
    childNodes map[rune]*Trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    if this == nil {
        return
    }

    for _, c := range word {
        if this.childNodes == nil {
            this.childNodes = make(map[rune]*Trie)
        }
        childNodes := this.childNodes
        if childNodes[c] == nil {
            childNodes[c] = &Trie{}
        }
        this = childNodes[c]
    }
    this.isEnd = true

}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    if this == nil {
        return false
    }

    for _, c := range word {
        if this == nil || this.childNodes == nil {
            return false
        }
        this = this.childNodes[c]
    }
    return this != nil && this.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    if this == nil {
        return false
    }

    for _, c := range prefix {
        if this == nil || this.childNodes == nil {
            return false
        }
        this = this.childNodes[c]
    }
    return this != nil
}


func findWords(board [][]byte, words []string) []string {
    trie := Trie{}
	wordMap := map[string]bool{}
	for _, word := range words {
		trie.Insert(word)
		wordMap[word] = true
	}

	m := len(board)
	n := len(board[0])
	const visited = '.'
	var results []string

	var dfs func(i int, j int, word string)
	dfs = func(i int, j int, word string) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if board[i][j] == visited || !trie.StartsWith(word) {
			return
		}
		word = word + string(board[i][j])
		if v, ok := wordMap[word]; ok && v {
			results = append(results, word)
            wordMap[word] = false
		}
		tmp := board[i][j]
		board[i][j] = visited
		dfs(i-1, j, word)
		dfs(i+1, j, word)
		dfs(i, j-1, word)
		dfs(i, j+1, word)
		board[i][j] = tmp
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, "")
		}
	}
	return results
}
