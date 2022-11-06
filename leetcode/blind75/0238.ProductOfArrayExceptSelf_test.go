package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, productExceptSelf(test.in))
	}
}
