package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfConnectedComponentsInAnUndirectedGraph(t *testing.T) {
	tests := []struct {
		in1      int
		in2      [][]int
		expected int
	}{
		{5, [][]int{{0, 1}, {1, 2}, {3, 4}}, 2},
		{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, 1},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, countComponents(tt.in1, tt.in2))
	}
}
