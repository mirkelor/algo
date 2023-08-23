package blind75

import "github.com/mirkelor/algo/common"

func isSubtree(root *common.TreeNode, subRoot *common.TreeNode) bool {
	if root == nil {
		return false
	}
	if isSubtreeHelper(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSubtreeHelper(p *common.TreeNode, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSubtreeHelper(p.Left, q.Left) && isSubtreeHelper(p.Right, q.Right)
}
