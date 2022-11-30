package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, lengthOfLIS(tt.in))
	}
}
