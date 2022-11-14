package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMinimumInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, findMin(test.in))
	}
}
