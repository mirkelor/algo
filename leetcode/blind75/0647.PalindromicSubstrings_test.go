package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindromicSubstrings(t *testing.T) {
	tests := []struct {
		in       string
		expected int
	}{
		{"abc", 3},
		{"aaa", 6},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, countSubstrings(tt.in))
	}
}
