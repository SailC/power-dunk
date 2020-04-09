![1](http://zxi.mytechroad.com/blog/wp-content/uploads/2017/09/208-ep74-1.png)

https://leetcode.com/articles/implement-trie-prefix-tree/

## Trie v.s. HashTable
- HashTable takes O(1) to search a key
- Trie takes O(l) to search a key (l is the length of the key)
- But Trie is good at search all keys that start with a common prefix

- 每个节点上只保存一个字符
- dummy根节点不带表字符
- every child represents a char
- 插入过程中如果节点不存在，就一直创造节点，直到单词最后一个字符。
- 优点： 访问速度快:
  - O(l) , l 是单词长度（字符个数）
- 缺点: 占用空间多 (total number of prefixes), 用空间换时间
  - 每个单词都不一样， 每个前缀也都不一样
  - O(n * l ^ l) n个单词 * 每个单词l个前缀，每个前缀大小～l
   


