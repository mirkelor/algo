package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraphValidTree(t *testing.T) {
	tests := []struct {
		in1      int
		in2      [][]int
		expected bool
	}{
		{5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}, true},
		{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}, false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, validTree(tt.in1, tt.in2))
	}
}
