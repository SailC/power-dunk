type Trie struct {
    isEnd bool
    childNodes map[rune]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
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


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
