package blind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOf1Bits(t *testing.T) {
	tests := []struct {
		in       uint32
		expected int
	}{
		{11, 3},
		{128, 1},
		{4294967293, 31},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, hammingWeight(tt.in))
	}
}
