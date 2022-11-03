package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJumpGame(t *testing.T) {
	tests := []struct {
		in       []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, canJump(test.in))
	}
}
