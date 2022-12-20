package blind75

type WordDictionary struct {
	children map[byte]*WordDictionary
	isWord   bool
}

func WordDictionaryConstructor() WordDictionary {
	return WordDictionary{children: make(map[byte]*WordDictionary)}
}

func (wd *WordDictionary) AddWord(word string) {
	cur := wd
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := cur.children[c]; !ok {
			cur.children[c] = &WordDictionary{children: make(map[byte]*WordDictionary)}
		}
		cur = cur.children[c]
	}
	cur.isWord = true
}

func (wd *WordDictionary) Search(word string) bool {
	if word == "" {
		return wd.isWord
	}

	c := word[0]
	if c == '.' {
		for _, child := range wd.children {
			if child.Search(word[1:]) {
				return true
			}
		}
		return false
	}

	if _, ok := wd.children[c]; !ok {
		return false
	}

	return wd.children[c].Search(word[1:])
}
