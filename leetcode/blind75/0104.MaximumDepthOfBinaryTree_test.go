package blind75

import (
	"testing"

	"github.com/Mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestMaximumDepthOfBinaryTree(t *testing.T) {
	tests := []struct {
		in       *common.TreeNode
		expected int
	}{
		{common.NewTreeNode(3, 9, 20, -1, -1, 15, 7), 3},
		{common.NewTreeNode(1, -1, 2), 2},
	}
	for _, tt := range tests {
		result := maxDepth(tt.in)
		assert.Equal(t, tt.expected, result)
	}
}
