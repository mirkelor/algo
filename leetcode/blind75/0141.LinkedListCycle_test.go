package blind75

import (
	"testing"

	"github.com/Mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestLinkedListCycle(t *testing.T) {
	tests := []struct {
		in       *common.ListNode
		expected bool
	}{
		{common.NewListNodeWithCycle(1, 3, 2, 0, -4), true},
		{common.NewListNodeWithCycle(1, 1, 2), true},
		{common.NewListNodeWithCycle(-1, 1), false},
	}

	for _, test := range tests {
		result := hasCycle(test.in)
		assert.Equal(t, test.expected, result)
	}
}
