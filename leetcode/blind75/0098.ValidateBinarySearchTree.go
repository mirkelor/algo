package blind75

import "github.com/mirkelor/algo/common"

func isValidBST(root *common.TreeNode) bool {
	return isValidBSTHelper(root, -1<<63, 1<<63-1)
}

func isValidBSTHelper(node *common.TreeNode, left int64, right int64) bool {
	if node == nil {
		return true
	}
	if !(left < int64(node.Val) && int64(node.Val) < right) {
		return false
	}
	return isValidBSTHelper(node.Left, left, int64(node.Val)) && isValidBSTHelper(node.Right, int64(node.Val), right)
}
