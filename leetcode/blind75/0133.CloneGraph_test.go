package blind75

import (
	"testing"

	"github.com/Mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		in       *common.Node
		expected *common.Node
	}{
		{common.NewNode([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}), common.NewNode([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})},
		{common.NewNode([][]int{{}}), common.NewNode([][]int{{}})},
		{nil, nil},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, cloneGraph(test.in))
	}
}
