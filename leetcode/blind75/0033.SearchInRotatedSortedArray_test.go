package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, search(test.in1, test.in2))
	}
}
