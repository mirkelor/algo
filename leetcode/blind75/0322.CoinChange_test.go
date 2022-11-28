package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, coinChange(tt.in1, tt.in2))
	}
}
