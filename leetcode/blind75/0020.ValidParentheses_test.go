package blind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		in       string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
	}

	for _, test := range tests {
		result := isValid(test.in)
		assert.Equal(t, test.expected, result)
	}
}
