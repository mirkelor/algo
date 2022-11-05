package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		in       []string
		expected [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, groupAnagrams(test.in))
	}
}
