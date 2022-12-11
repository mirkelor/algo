package blind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBits(t *testing.T) {
	tests := []struct {
		in       uint32
		expected uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, reverseBits(tt.in))
	}
}
