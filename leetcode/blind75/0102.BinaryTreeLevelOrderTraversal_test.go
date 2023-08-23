package blind75

import (
	"testing"

	"github.com/mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {
	tests := []struct {
		in       *common.TreeNode
		expected [][]int
	}{
		{common.NewTreeNode(3, 9, 20, -1, -1, 15, 7), [][]int{{3}, {9, 20}, {15, 7}}},
		{common.NewTreeNode(1), [][]int{{1}}},
		{common.NewTreeNode(), nil},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, levelOrder(test.in))
	}
}
