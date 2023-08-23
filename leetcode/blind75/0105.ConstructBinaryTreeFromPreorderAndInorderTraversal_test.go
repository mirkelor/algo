package blind75

import (
	"testing"

	"github.com/mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestConstructBinaryTreeFromPreorderAndInorderTraversal(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      []int
		expected *common.TreeNode
	}{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, common.NewTreeNode(3, 9, 20, -1, -1, 15, 7)},
		{[]int{1}, []int{1}, common.NewTreeNode(1)},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, buildTree(tt.in1, tt.in2))
	}
}
