package blind75

import (
	"testing"

	"github.com/mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestSerializeAndDeserializeBinaryTree(t *testing.T) {
	tests := []struct {
		in       *common.TreeNode
		expected *common.TreeNode
	}{
		{common.NewTreeNode(1, 2, 3, -1, -1, 4, 5), common.NewTreeNode(1, 2, 3, -1, -1, 4, 5)},
		{common.NewTreeNode(1, 2, 3, -1, -1, 4, 5, -1, -1, -1, -1, 6, 7), common.NewTreeNode(1, 2, 3, -1, -1, 4, 5, -1, -1, -1, -1, 6, 7)},
		{nil, nil},
	}

	for _, tt := range tests {
		c := Codec{}
		assert.Equal(t, tt.expected, c.deserialize(c.serialize(tt.in)))
	}
}
