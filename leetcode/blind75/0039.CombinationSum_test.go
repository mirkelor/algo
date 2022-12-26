package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{2}, 1, [][]int{}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, combinationSum(tt.in1, tt.in2))
	}
}
