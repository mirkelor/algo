package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDesignAddAndSearchWordsDataStructure(t *testing.T) {
	wd := WordDictionaryConstructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	assert.False(t, wd.Search("pad"))
	assert.True(t, wd.Search("bad"))
	assert.True(t, wd.Search(".ad"))
	assert.True(t, wd.Search("b.."))
}
