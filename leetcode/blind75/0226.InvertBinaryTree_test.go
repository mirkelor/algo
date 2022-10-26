package blind75

import (
	"testing"

	"github.com/Mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestInvertBinaryTree(t *testing.T) {
	tests := []struct {
		in       *common.TreeNode
		expected *common.TreeNode
	}{
		{common.NewTreeNode(4, 2, 7, 1, 3, 6, 9), common.NewTreeNode(4, 7, 2, 9, 6, 3, 1)},
		{common.NewTreeNode(2, 1, 3), common.NewTreeNode(2, 3, 1)},
		{common.NewTreeNode(), common.NewTreeNode()},
	}

	for _, test := range tests {
		result := invertTree(test.in)
		assert.Equal(t, test.expected, result)
	}
}
