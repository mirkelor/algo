package blind75

import (
	"testing"

	"github.com/Mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		in       *common.ListNode
		expected *common.ListNode
	}{
		{common.NewListNode(1, 2, 3, 4, 5), common.NewListNode(5, 4, 3, 2, 1)},
		{common.NewListNode(1, 2), common.NewListNode(2, 1)},
		{common.NewListNode(), common.NewListNode()},
	}

	for _, test := range tests {
		result := reverseList(test.in)
		assert.Equal(t, test.expected, result)
	}
}
