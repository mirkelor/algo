package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		in       [][]int
		expected [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, merge(tt.in))
	}
}
