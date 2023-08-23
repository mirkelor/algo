package blind75

import (
	"testing"

	"github.com/mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestMergeKSortedLists(t *testing.T) {
	tests := []struct {
		in       []*common.ListNode
		expected *common.ListNode
	}{
		{[]*common.ListNode{common.NewListNode(1, 4, 5), common.NewListNode(1, 3, 4), common.NewListNode(2, 6)}, common.NewListNode(1, 1, 2, 3, 4, 4, 5, 6)},
		{[]*common.ListNode{}, common.NewListNode()},
		{[]*common.ListNode{common.NewListNode()}, common.NewListNode()},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, mergeKLists(tt.in))
	}
}
