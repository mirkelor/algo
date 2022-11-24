package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumProductSubarray(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, maxProduct(tt.in))
	}
}
