package blind75

import (
	"testing"

	"github.com/mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestSameTree(t *testing.T) {
	tests := []struct {
		in1      common.TreeNode
		in2      common.TreeNode
		expected bool
	}{
		{*common.NewTreeNode(1, 2, 3), *common.NewTreeNode(1, 2, 3), true},
		{*common.NewTreeNode(1, 2), *common.NewTreeNode(1, -1, 2), false},
		{*common.NewTreeNode(1, 2, 1), *common.NewTreeNode(1, 1, 2), false},
	}
	for _, tt := range tests {
		result := isSameTree(&tt.in1, &tt.in2)
		assert.Equal(t, tt.expected, result)
	}
}
