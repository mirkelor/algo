package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		in1      string
		in2      string
		expected int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, longestCommonSubsequence(tt.in1, tt.in2))
	}
}
