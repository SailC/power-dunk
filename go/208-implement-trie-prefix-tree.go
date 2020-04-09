// https://leetcode-cn.com/problems/implement-trie-prefix-tree/

type Trie struct {
    isEnd bool
    childNodes map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
        isEnd: false,
        childNodes: make(map[rune]*Trie),
    }
}

/** Inserts a word into the trie. */
func (root *Trie) Insert(word string)  {
    cur := root
    for _, c := range word {
        childNodes := cur.childNodes
        if childNodes[c] == nil {
            childNodes[c] = &Trie{
                isEnd:false,
                childNodes: make(map[rune]*Trie),
            }
        }
        cur = childNodes[c]
    }
    cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (root *Trie) Search(word string) bool {
   lastNode := root.getLastNode(word)
   return lastNode != nil && lastNode.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (root *Trie) StartsWith(prefix string) bool {
    lastNode := root.getLastNode(prefix)
    return lastNode != nil
}

func (root *Trie) getLastNode(prefix string) *Trie {
	cur := root
	for _, c := range prefix {
        if _, ok := cur.childNodes[c]; !ok {
	        return nil
        }
        cur = cur.childNodes[c]
    }
    return cur
}
