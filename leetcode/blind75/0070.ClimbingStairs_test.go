package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbingStairs(t *testing.T) {
	tests := []struct {
		in       int
		expected int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, climbStairs(tt.in))
	}
}
