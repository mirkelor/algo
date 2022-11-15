package blind75

import (
	"testing"

	"github.com/Mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestLowestCommonAncestorOfABinarySearchTree(t *testing.T) {
	tests := []struct {
		in1      *common.TreeNode
		in2      *common.TreeNode
		in3      *common.TreeNode
		expected *common.TreeNode
	}{
		{common.NewTreeNode(6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5), common.NewTreeNode(2), common.NewTreeNode(8), common.NewTreeNode(6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5)},
		{common.NewTreeNode(6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5), common.NewTreeNode(2), common.NewTreeNode(4), common.NewTreeNode(2, 0, 4, -1, -1, 3, 5)},
		{common.NewTreeNode(2, 1), common.NewTreeNode(2), common.NewTreeNode(1), common.NewTreeNode(2, 1)},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, lowestCommonAncestor(test.in1, test.in2, test.in3))
	}
}
