package blind75

import (
	"testing"

	"github.com/Mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestValidateBinarySearchTree(t *testing.T) {
	tests := []struct {
		in       *common.TreeNode
		expected bool
	}{
		{common.NewTreeNode(2, 1, 3), true},
		{common.NewTreeNode(5, 1, 4, -1, -1, 3, 6), false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, isValidBST(tt.in))
	}
}
