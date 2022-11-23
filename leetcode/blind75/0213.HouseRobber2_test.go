package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHouseRobber2(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{2, 3, 2}, 3},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{1, 2, 3}, 3},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, rob2(tt.in))
	}
}
