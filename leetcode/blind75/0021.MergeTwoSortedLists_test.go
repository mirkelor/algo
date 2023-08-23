package blind75

import (
	"testing"

	"github.com/mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoSortedLists(t *testing.T) {
	tests := []struct {
		in1      *common.ListNode
		in2      *common.ListNode
		expected *common.ListNode
	}{
		{common.NewListNode(1, 2, 4), common.NewListNode(1, 3, 4), common.NewListNode(1, 1, 2, 3, 4, 4)},
		{nil, nil, nil},
		{nil, common.NewListNode(0), common.NewListNode(0)},
	}

	for _, test := range tests {
		result := mergeTwoLists(test.in1, test.in2)
		assert.Equal(t, test.expected, result)
	}
}
