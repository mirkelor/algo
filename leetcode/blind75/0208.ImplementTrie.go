package blind75

type Trie struct {
	children map[byte]*Trie
	isWord   bool
}

func TrieConstructor() Trie {
	return Trie{children: make(map[byte]*Trie)}
}

func (t *Trie) Insert(word string) {
	cur := t
	for i := 0; i < len(word); i++ {
		if _, ok := cur.children[word[i]]; !ok {
			cur.children[word[i]] = &Trie{children: make(map[byte]*Trie)}
		}
		cur = cur.children[word[i]]
	}
	cur.isWord = true
}

func (t *Trie) Search(word string) bool {
	cur := t
	for i := 0; i < len(word); i++ {
		if _, ok := cur.children[word[i]]; !ok {
			return false
		}
		cur = cur.children[word[i]]
	}
	return cur.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t
	for i := 0; i < len(prefix); i++ {
		if _, ok := cur.children[prefix[i]]; !ok {
			return false
		}
		cur = cur.children[prefix[i]]
	}
	return true
}
