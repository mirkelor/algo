package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	tests := []struct {
		in1      [][]int
		in2      []int
		expected [][]int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, insert(tt.in1, tt.in2))
	}
}
