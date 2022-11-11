package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	tests := []struct {
		in       string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, lengthOfLongestSubstring(test.in))
	}
}
