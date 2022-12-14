package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumWindowSubstring(t *testing.T) {
	tests := []struct {
		in1      string
		in2      string
		expected string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, minWindow(tt.in1, tt.in2))
	}
}
