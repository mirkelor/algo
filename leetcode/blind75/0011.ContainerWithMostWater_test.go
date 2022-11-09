package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainerWithMostWater(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, maxArea(test.in))
	}
}
