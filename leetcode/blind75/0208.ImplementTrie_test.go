package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImplementTrie(t *testing.T) {
	trie := TrieConstructor()
	trie.Insert("apple")
	assert.True(t, trie.Search("apple"))
	assert.False(t, trie.Search("app"))
	assert.True(t, trie.StartsWith("app"))
	trie.Insert("app")
	assert.True(t, trie.Search("app"))
}
