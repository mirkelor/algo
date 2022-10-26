package blind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		in       string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for _, test := range tests {
		result := isPalindrome(test.in)
		assert.Equal(t, test.expected, result)
	}
}
