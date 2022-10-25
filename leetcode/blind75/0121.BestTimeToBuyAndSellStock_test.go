package blind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, test := range tests {
		result := maxProfit(test.in)
		assert.Equal(t, test.expected, result)
	}
}
