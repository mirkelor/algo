package blind75

import (
	"testing"

	"github.com/mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestKthSmallestElementInABST(t *testing.T) {
	tests := []struct {
		in1      *common.TreeNode
		in2      int
		expected int
	}{
		{common.NewTreeNode(3, 1, 4, -1, 2), 1, 1},
		{common.NewTreeNode(5, 3, 6, 2, 4, -1, -1, 1), 3, 3},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, kthSmallest(tt.in1, tt.in2))
	}
}
