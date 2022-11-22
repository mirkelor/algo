package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHouseRobber(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, rob(tt.in))
	}
}
