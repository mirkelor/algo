package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		expected int
	}{
		{3, 7, 28},
		{3, 2, 3},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, uniquePaths(tt.in1, tt.in2))
	}
}
