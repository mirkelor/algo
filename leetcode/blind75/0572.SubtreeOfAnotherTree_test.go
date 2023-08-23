package blind75

import (
	"testing"

	"github.com/mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestSubtreeOfAnotherTree(t *testing.T) {
	tests := []struct {
		in1      *common.TreeNode
		in2      *common.TreeNode
		expected bool
	}{
		{common.NewTreeNode(3, 4, 5, 1, 2), common.NewTreeNode(4, 1, 2), true},
		{common.NewTreeNode(3, 4, 5, 1, 2, -1, -1, -1, -1, 0), common.NewTreeNode(4, 1, 2), false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, isSubtree(tt.in1, tt.in2))
	}
}
