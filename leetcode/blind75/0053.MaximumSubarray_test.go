package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, maxSubArray(test.in))
	}
}
