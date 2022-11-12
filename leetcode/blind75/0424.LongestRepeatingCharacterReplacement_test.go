package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestRepeatingCharacterReplacement(t *testing.T) {
	tests := []struct {
		in1      string
		in2      int
		expected int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, characterReplacement(test.in1, test.in2))
	}
}
