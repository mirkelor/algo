package blind75

import (
	"testing"

	"github.com/Mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestRemoveNthNodeFromEndOfList(t *testing.T) {
	tests := []struct {
		in1      *common.ListNode
		in2      int
		expected *common.ListNode
	}{
		{common.NewListNode(1, 2, 3, 4, 5), 2, common.NewListNode(1, 2, 3, 5)},
		{common.NewListNode(1), 1, common.NewListNode()},
		{common.NewListNode(1, 2), 1, common.NewListNode(1)},
	}

	for _, test := range tests {
		result := removeNthFromEnd(test.in1, test.in2)
		assert.Equal(t, test.expected, result)
	}
}
