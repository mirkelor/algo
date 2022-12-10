package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountingBits(t *testing.T) {
	tests := []struct {
		in       int
		expected []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, countBits(tt.in))
	}
}
