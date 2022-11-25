package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, longestPalindrome(tt.in))
	}
}
