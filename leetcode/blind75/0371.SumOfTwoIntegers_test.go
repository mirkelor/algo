package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfTwoIntegers(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		expected int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{-2, 3, 1},
		{-2, -3, -5},
		{2, -2, 0},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, getSum(tt.in1, tt.in2))
	}
}
