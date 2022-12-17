package blind75

import (
	"testing"

	"github.com/Mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeMaximumPathSum(t *testing.T) {
	tests := []struct {
		in       *common.TreeNode
		expected int
	}{
		{common.NewTreeNode(1, 2, 3), 6},
		{common.NewTreeNode(-10, 9, 20, -1, -1, 15, 7), 42},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, maxPathSum(tt.in))
	}
}
