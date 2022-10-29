package blind75

import (
	"testing"

	"github.com/Mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		in       *common.ListNode
		expected *common.ListNode
	}{
		{common.NewListNode(1, 2, 3, 4), common.NewListNode(1, 4, 2, 3)},
		{common.NewListNode(1, 2, 3, 4, 5), common.NewListNode(1, 5, 2, 4, 3)},
	}

	for _, test := range tests {
		reorderList(test.in)
		assert.Equal(t, test.expected, test.in)
	}
}
