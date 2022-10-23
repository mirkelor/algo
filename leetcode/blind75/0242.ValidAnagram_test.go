package blind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		in1 string
		in2 string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for _, test := range tests {
		result := isAnagram(test.in1, test.in2)
		assert.Equal(t, test.expected, result)
	}
}
