package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, missingNumber(tt.in))
	}
}
