package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNonOverlappingIntervals(t *testing.T) {
	tests := []struct {
		in       [][]int
		expected int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{[][]int{{1, 2}, {2, 3}}, 0},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, eraseOverlapIntervals(tt.in))
	}
}
