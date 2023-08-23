package blind75

import "github.com/mirkelor/algo/common"

func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + common.Max(maxDepth(root.Left), maxDepth(root.Right))
}
