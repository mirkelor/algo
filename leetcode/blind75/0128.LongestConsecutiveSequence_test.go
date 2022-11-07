package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsecutiveSubsequence(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, longestConsecutive(test.in))
	}
}
