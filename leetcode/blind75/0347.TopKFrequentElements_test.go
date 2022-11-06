package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopKFrequentElements(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, topKFrequent(test.in1, test.in2))
	}
}
