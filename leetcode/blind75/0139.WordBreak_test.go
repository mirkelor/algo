package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		in1      string
		in2      []string
		expected bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, wordBreak(tt.in1, tt.in2))
	}
}
